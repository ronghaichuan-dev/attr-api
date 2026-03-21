package queue

import (
	"context"
	"net/http"
	"net/url"
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
	return SQSConfig{
		Region:      util.GetConfigString(context.Background(), "aws.region", "us-east-1"),
		Endpoint:    util.GetConfigString(context.Background(), "aws.endpoint", "https://sqs.us-west-2.amazonaws.com"),
		AccessKey:   util.GetConfigString(context.Background(), "aws.accessKey", ""),
		SecretKey:   util.GetConfigString(context.Background(), "aws.secretKey", ""),
		QueuePrefix: util.GetConfigString(context.Background(), "aws.queuePrefix", ""),
		QueueURL:    util.GetConfigString(context.Background(), "aws.queueURL", "https://sqs.us-west-2.amazonaws.com/717279709976/rszai-free.fifo"),
	}
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

		sqsConfig := &aws.Config{
			Region:     aws.String(config.Region),
			Endpoint:   aws.String(config.Endpoint),
			DisableSSL: aws.Bool(false),
			LogLevel:   aws.LogLevel(aws.LogOff),
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

func getDirectQueueURL() string {
	config := getSQSConfig()
	return config.QueueURL
}

func getQueueURL(queueName string) (string, error) {
	if err := initSQS(); err != nil {
		return "", err
	}

	result, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		logger.Errorf("获取队列URL失败: %v", err)
		return "", err
	}

	return aws.StringValue(result.QueueUrl), nil
}

func CreateQueue(queueName string) (string, error) {
	if err := initSQS(); err != nil {
		return "", err
	}

	fullQueueName := GetQueueName(queueName)

	result, err := sqsClient.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(fullQueueName),
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

	logger.Infof("SQS队列创建成功: %s, URL: %s", fullQueueName, aws.StringValue(result.QueueUrl))
	return aws.StringValue(result.QueueUrl), nil
}

func PublishToQueue(queueName string, data []byte) error {
	logger.Info("开始 PublishToQueue")

	if err := initSQS(); err != nil {
		logger.Errorf("initSQS失败: %v", err)
		return err
	}
	logger.Info("initSQS完成")

	// 先尝试获取队列URL，检查队列是否存在
	queueURL, err := getQueueURL(queueName)
	if err != nil {
		logger.Warnf("队列不存在，正在创建: %s", queueName)
		// 队列不存在，创建队列
		queueURL, err = CreateQueue(queueName)
		if err != nil {
			logger.Errorf("CreateQueue失败: %v", err)
			return err
		}
		logger.Infof("队列创建成功: %s", queueURL)
	}

	logger.Infof("开始发送消息到SQS, queueURL: %s", queueURL)
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

	logger.Info("发送消息到SQS成功")

	return nil
}

type MessageHandler func(data []byte) error

func ConsumeQueue(queueName string, handler MessageHandler) error {
	if err := initSQS(); err != nil {
		return err
	}

	queueURL, err := getQueueURL(queueName)
	if err != nil {
		queueURL, err = CreateQueue(queueName)
		if err != nil {
			return err
		}
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
