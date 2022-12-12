package initialize

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	proto "message/api/qvbilam/message/v1"
	"message/business"
	"message/enum"
	"message/global"
	"strconv"
)

func InitQueue() {
	user := "admin"
	password := "admin"
	host := "127.0.0.1"
	port := 5672

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)
	conn, err := amqp.Dial(url)
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列服务", err)
		panic(any(err))
	}

	global.MessageQueueClient = conn
	suffix := global.ServerConfig.RabbitMQServerConfig.QueueSuffix
	if suffix == "" {
		suffix = strconv.Itoa(int(global.ServerConfig.Port))
	}
	res, err := global.MessageServerClient.CreateQueue(context.Background(), &proto.UpdateQueueRequest{})
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列服务", err)
		panic(any(err))
	}
	exchangeName := res.ExchangeName
	queueName := res.Name

	fmt.Printf("create queue exchange: %s\n", exchangeName)
	fmt.Printf("create queue: %s\n", queueName)

	// 全局变量
	global.ServerConfig.RabbitMQServerConfig.MessageExchangeName = exchangeName
	global.ServerConfig.RabbitMQServerConfig.MessageQueueName = queueName

	// 创建队列
	business.CreateExchange(exchangeName)
	business.CreateQueue(queueName, exchangeName)
	_, err = global.MessageServerClient.UpdateQueue(context.Background(), &proto.UpdateQueueRequest{
		Name:   queueName,
		Status: enum.QueueStatusOpen,
	})
	if err != nil {
		panic(any(err))
	}

	// 接受消息
	go business.ConsumeQueue(queueName)
}
