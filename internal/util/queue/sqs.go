package queue

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"god-help-service/internal/util"
	"god-help-service/internal/util/logger"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	sqsSession *session.Session
	sqsClient  *sqs.SQS
	sqsOnce    sync.Once
)

type SQSConfig struct {
	Region      string
	Endpoint    string
	AccessKey   string
	SecretKey   string
	QueuePrefix string
	QueueURL    string
}

func getSQSConfig() SQSConfig {
	ctx := context.Background()
	config := SQSConfig{
		Region:      util.GetConfigString(ctx, "aws.region", ""),
		Endpoint:    util.GetConfigString(ctx, "aws.endpoint", ""),
		AccessKey:   util.GetConfigString(ctx, "aws.accessKey", ""),
		SecretKey:   util.GetConfigString(ctx, "aws.secretKey", ""),
		QueuePrefix: util.GetConfigString(ctx, "aws.queuePrefix", ""),
		QueueURL:    util.GetConfigString(ctx, "aws.queueURL", ""),
	}
	if config.Region == "" {
		logger.Warnf("aws.region 未配置，请检查配置文件")
	}
	if config.Endpoint == "" {
		logger.Warnf("aws.endpoint 未配置，请检查配置文件")
	}
	return config
}

func getProxyConfig() (string, string, string) {
	proxyURL := util.GetConfigString(context.Background(), "proxy.url", "")
	proxyUser := util.GetConfigString(context.Background(), "proxy.user", "")
	proxyPass := util.GetConfigString(context.Background(), "proxy.password", "")
	return proxyURL, proxyUser, proxyPass
}

func initSQS() error {
	var initErr error
	sqsOnce.Do(func() {
		config := getSQSConfig()
		proxyURL, proxyUser, proxyPass := getProxyConfig()

		logger.Infof("初始化SQS客户端，Region: %s, Endpoint: %s", config.Region, config.Endpoint)

		var err error
		transport := &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 5,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		}

		// 配置代理
		if proxyURL != "" {
			proxy, err := url.Parse("http://" + proxyURL)
			if err != nil {
				logger.Warnf("解析代理URL失败: %v", err)
			} else {
				// 设置代理认证
				if proxyUser != "" && proxyPass != "" {
					proxy.User = url.UserPassword(proxyUser, proxyPass)
				}
				transport.Proxy = http.ProxyURL(proxy)
				logger.Infof("使用代理: %s", proxyURL)
			}
		}

		// 根据 endpoint 自动判断是否禁用 SSL（LocalStack 使用 http）
		disableSSL := strings.HasPrefix(config.Endpoint, "http://")

		sqsConfig := &aws.Config{
			Region:           aws.String(config.Region),
			Endpoint:         aws.String(config.Endpoint),
			DisableSSL:       aws.Bool(disableSSL),
			S3ForcePathStyle: aws.Bool(disableSSL), // LocalStack 需要路径风格
			LogLevel:         aws.LogLevel(aws.LogOff),
			HTTPClient: &http.Client{
				Timeout:   30 * time.Second,
				Transport: transport,
			},
		}

		accessKey := config.AccessKey
		secretKey := config.SecretKey
		if accessKey != "" && secretKey != "" {
			logger.Info("使用静态凭证连接SQS")
			sqsConfig.Credentials = credentials.NewStaticCredentials(accessKey, secretKey, "")
		} else {
			logger.Info("使用IAM角色/默认凭证提供者链连接SQS")
		}

		sqsSession, err = session.NewSession(sqsConfig)
		if err != nil {
			initErr = err
			logger.Errorf("创建SQS会话失败: %v", err)
			return
		}

		sqsClient = sqs.New(sqsSession)
		logger.Infof("SQS客户端初始化完成，Endpoint: %s, Region: %s", config.Endpoint, config.Region)
	})
	return initErr
}

func GetQueueName(subject string) string {
	config := getSQSConfig()
	return config.QueuePrefix + subject
}

// ensureQueue 确保队列存在，不存在则自动创建，返回队列 URL
func ensureQueue(queueName string) (string, error) {
	if err := initSQS(); err != nil {
		return "", err
	}

	// 先尝试获取
	result, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err == nil {
		return aws.StringValue(result.QueueUrl), nil
	}

	// 队列不存在，自动创建
	logger.Infof("队列 %s 不存在，自动创建中...", queueName)
	createResult, err := sqsClient.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
		Attributes: map[string]*string{
			"VisibilityTimeout":             aws.String("30"),
			"MessageRetentionPeriod":        aws.String("86400"),
			"ReceiveMessageWaitTimeSeconds": aws.String("20"),
		},
	})
	if err != nil {
		logger.Errorf("创建SQS队列失败: %v", err)
		return "", err
	}

	queueURL := aws.StringValue(createResult.QueueUrl)
	logger.Infof("SQS队列创建成功: %s, URL: %s", queueName, queueURL)
	return queueURL, nil
}

func PublishToQueue(queueName string, data []byte) error {
	queueURL, err := ensureQueue(queueName)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err = sqsClient.SendMessageWithContext(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(string(data)),
	})
	if err != nil {
		logger.Errorf("发送消息到SQS失败: %v", err)
		return err
	}

	logger.Infof("发送消息到SQS成功, queue: %s", queueName)
	return nil
}

type MessageHandler func(data []byte) error

func ConsumeQueue(queueName string, handler MessageHandler) error {
	queueURL, err := ensureQueue(queueName)
	if err != nil {
		return err
	}

	logger.Infof("开始消费SQS队列: %s", queueName)

	for {
		result, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL),
			MaxNumberOfMessages: aws.Int64(1),
			WaitTimeSeconds:     aws.Int64(20),
			VisibilityTimeout:   aws.Int64(30),
		})
		if err != nil {
			logger.Errorf("接收SQS消息失败: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, msg := range result.Messages {
			err = handler([]byte(aws.StringValue(msg.Body)))
			if err != nil {
				logger.Errorf("处理SQS消息失败: %v", err)
				continue
			}

			_, err = sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueURL),
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				logger.Errorf("删除SQS消息失败: %v", err)
			}
		}
	}
}

type QueueStats struct {
	MessagesAvailable int64
	MessagesInFlight  int64
}
