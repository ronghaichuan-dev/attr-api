package queue

import (
	"context"
	"god-help-service/internal/consts"
	"god-help-service/internal/util"
	"sync"
	"time"

	"god-help-service/internal/util/logger"

	"github.com/nats-io/nats.go"
)

var (
	// ConnPool Nats连接池实例
	ConnPool *NatsConnPool
	// natsConnOnce 确保Nats连接池只初始化一次
	natsConnOnce sync.Once
)

// NatsConnPool Nats连接池结构
type NatsConnPool struct {
	connections chan *nats.Conn
	jetStreams  chan nats.JetStreamContext
	size        int
	mu          sync.Mutex
	closed      bool
	lastUsed    map[*nats.Conn]time.Time // 记录连接最后使用时间
	idleTimeout time.Duration            // 空闲连接超时时间
}

// InitNatsConn 初始化Nats连接池
func InitNatsConn() error {
	var err error
	natsConnOnce.Do(func() {
		err = initNatsPool()
	})
	return err
}

// initNatsPool 初始化Nats连接池
func initNatsPool() error {
	// 从配置中读取连接池大小，默认5
	poolSize := util.GetConfigInt(context.Background(), "nats.poolSize", 5)
	if poolSize <= 0 {
		poolSize = 5
	}

	// 创建连接池
	ConnPool = &NatsConnPool{
		connections: make(chan *nats.Conn, poolSize),
		jetStreams:  make(chan nats.JetStreamContext, poolSize),
		size:        poolSize,
		closed:      false,
		lastUsed:    make(map[*nats.Conn]time.Time),
		idleTimeout: 30 * time.Minute, // 30分钟空闲超时
	}

	// 启动空闲连接回收协程
	go ConnPool.reapIdleConnections()

	// 填充连接池
	for i := 0; i < poolSize; i++ {
		conn, js, err := createNatsConn()
		if err != nil {
			logger.Errorf("创建Nats连接失败: %v", err)
			continue
		}
		ConnPool.connections <- conn
		ConnPool.jetStreams <- js
		ConnPool.lastUsed[conn] = time.Now()
	}

	// 检查是否有至少一个连接
	if len(ConnPool.connections) == 0 {
		return nats.ErrConnectionClosed
	}

	logger.Infof("Nats连接池初始化完成，大小: %d, 活跃连接: %d", poolSize, len(ConnPool.connections))
	return nil
}

// reapIdleConnections 回收空闲连接
func (p *NatsConnPool) reapIdleConnections() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C
		p.mu.Lock()
		if p.closed {
			p.mu.Unlock()
			return
		}

		now := time.Now()
		idleConns := []*nats.Conn{}

		// 检查所有连接的最后使用时间
		for conn, lastTime := range p.lastUsed {
			if now.Sub(lastTime) > p.idleTimeout {
				idleConns = append(idleConns, conn)
			}
		}

		// 移除并关闭空闲连接
		for _, conn := range idleConns {
			delete(p.lastUsed, conn)
			conn.Close()
		}

		p.mu.Unlock()
	}
}

// createNatsConn 创建单个Nats连接
func createNatsConn() (*nats.Conn, nats.JetStreamContext, error) {
	// 连接到Nats服务器
	client, err := nats.Connect(
		nats.DefaultURL,
		nats.Name("god-help-service"),
		nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {
			logger.Errorf("Nats连接错误: %v", err)
		}),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			logger.Warnf("Nats服务已断开连接: %v", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logger.Info("Nats重新连接成功")
		}),
		nats.MaxReconnects(-1),            // 无限次重新连接
		nats.ReconnectWait(5*time.Second), // 每次重连等待5秒
		nats.Timeout(10*time.Second),      // 添加连接超时
	)
	if err != nil {
		logger.Errorf("连接Nats失败: %v", err)
		return nil, nil, err
	}

	// 初始化JetStream
	js, err := initJetStream(client)
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, js, nil
}

// initJetStream 初始化JetStream
func initJetStream(nc *nats.Conn) (nats.JetStreamContext, error) {
	// 初始化JetStream
	js, err := nc.JetStream()
	if err != nil {
		logger.Errorf("初始化JetStream失败: %v", err)
		return nil, err
	}

	// 创建Nats流
	config := &nats.StreamConfig{
		Name: "god-help-service-stream",
		Subjects: []string{
			consts.EventQueueSubject,
			consts.NotificationQueueSubject,
			consts.AttributionQueueSubject,
		},
		Storage:   nats.FileStorage,
		Retention: nats.WorkQueuePolicy,
		MaxAge:    24 * 3600 * time.Second, // 24小时过期
	}

	// 尝试删除现有的流，以避免消费者冲突
	_ = js.DeleteStream(config.Name)

	// 创建新的流
	_, err = js.AddStream(config)
	if err != nil {
		logger.Errorf("创建Nats流失败: %v", err)
		return nil, err
	}

	return js, nil
}

// GetConnection 从连接池获取Nats连接
func (p *NatsConnPool) GetConnection() (*nats.Conn, nats.JetStreamContext, error) {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return nil, nil, nats.ErrConnectionClosed
	}
	p.mu.Unlock()

	select {
	case conn := <-p.connections:
		js := <-p.jetStreams
		// 检查连接是否健康
		if !conn.IsConnected() {
			// 连接已断开，创建新连接
			newConn, newJS, err := createNatsConn()
			if err != nil {
				// 创建失败，尝试使用其他连接
				p.connections <- conn
				p.jetStreams <- js
				return nil, nil, err
			}
			// 更新最后使用时间
			p.mu.Lock()
			p.lastUsed[newConn] = time.Now()
			p.mu.Unlock()
			return newConn, newJS, nil
		}
		// 更新最后使用时间
		p.mu.Lock()
		p.lastUsed[conn] = time.Now()
		p.mu.Unlock()
		return conn, js, nil
	default:
		// 连接池为空，创建新连接
		newConn, newJS, err := createNatsConn()
		if err != nil {
			return nil, nil, err
		}
		// 更新最后使用时间
		p.mu.Lock()
		p.lastUsed[newConn] = time.Now()
		p.mu.Unlock()
		return newConn, newJS, nil
	}
}

// ReturnConnection 归还Nats连接到连接池
func (p *NatsConnPool) ReturnConnection(conn *nats.Conn, js nats.JetStreamContext) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		delete(p.lastUsed, conn)
		conn.Close()
		return
	}

	// 更新最后使用时间
	p.lastUsed[conn] = time.Now()

	select {
	case p.connections <- conn:
		select {
		case p.jetStreams <- js:
		default:
			delete(p.lastUsed, conn)
			conn.Close()
		}
	default:
		// 连接池已满，关闭多余的连接
		delete(p.lastUsed, conn)
		conn.Close()
	}
}

// Close 关闭连接池
func (p *NatsConnPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.connections)
	close(p.jetStreams)

	// 关闭所有连接
	for conn := range p.connections {
		if conn != nil {
			delete(p.lastUsed, conn)
			conn.Close()
		}
	}

	logger.Info("Nats连接池已关闭")
}

// PublishToStream 发布消息到Nats流
func PublishToStream(subject string, data []byte) error {
	// 确保连接池已初始化
	if ConnPool == nil {
		if err := InitNatsConn(); err != nil {
			return err
		}
	}

	// 从连接池获取连接
	conn, js, err := ConnPool.GetConnection()
	if err != nil {
		return err
	}
	defer ConnPool.ReturnConnection(conn, js)

	_, err = js.Publish(subject, data)
	if err != nil {
		logger.Errorf("发布消息到Nats流失败: %v", err)
		return err
	}

	return nil
}

// SubscribeToStream 订阅Nats流消息
func SubscribeToStream(subject string, queueGroup string, handler nats.MsgHandler) (*nats.Subscription, error) {
	if ConnPool == nil {
		if err := InitNatsConn(); err != nil {
			return nil, err
		}
	}

	conn, js, err := ConnPool.GetConnection()
	if err != nil {
		return nil, err
	}

	sub, err := js.QueueSubscribe(subject, queueGroup, handler,
		nats.AckWait(30*time.Second),
		nats.MaxDeliver(5),
	)
	if err != nil {
		logger.Errorf("订阅Nats流失败: %v", err)
		ConnPool.ReturnConnection(conn, js)
		return nil, err
	}

	ConnPool.ReturnConnection(conn, js)

	return sub, nil
}
