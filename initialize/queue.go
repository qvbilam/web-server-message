package initialize

import (
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"message/business"
	"message/global"
)

var ExchangeName = "qvbilam-message-exchange"
var QueueName = "qvbilam-message-queue-1"

func InitQueue() {
	user := "admin"
	password := "admin"
	host := "127.0.0.1"
	port := 5672

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)
	conn, err := amqp.Dial(url)
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列服务", err)
	}

	global.MessageQueueClient = conn
	// 创建队列
	business.CreateExchange(ExchangeName)
	business.CreateQueue(QueueName, ExchangeName)

	// 接受消息
	go business.Consume(QueueName)
}

