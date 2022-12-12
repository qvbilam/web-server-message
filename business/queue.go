package business

import (
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"message/global"
)

func CreateExchange(exchangeName string) {
	// 建立 amqp 通道
	ch, err := global.MessageQueueClient.Channel()
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "建立通道失败", err)
	}

	// 创建交换机(不存在创建)
	if err := ch.ExchangeDeclare(
		exchangeName,
		"fanout",
		true,
		false,
		false,
		false,
		nil); err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列交换机", err)
	}
}

func CreateQueue(queueName, exchangeName string) {
	ch, err := global.MessageQueueClient.Channel()
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列通道", err)
	}
	q, _ := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil)

	// 绑定交换机
	if exchangeName != "" {
		if err := ch.QueueBind(q.Name, "", exchangeName, false, nil); err != nil {
			zap.S().Fatalf("%s dial error: %s", "队列绑定交换机失败", err)
		}
	}
}

func DeleteQueue(queueName string) {
	ch, err := global.MessageQueueClient.Channel()
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列通道", err)
	}
	_, err = ch.QueueDelete(queueName, false, true, true)
	if err != nil {
		// 删除队列
		return
	}
	// 删除队列失败
}

// ConsumeQueue 消费消息
func ConsumeQueue(queueName string) {
	ch, err := global.MessageQueueClient.Channel()
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "队列通道", err)
	}
	deliveries, err := ch.Consume(queueName, "go-consumer", true, false, false, false, nil)
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", "消费消息失败", err)
	}

	for msg := range deliveries {
		//fmt.Printf("read message: %s\n", msg.Body)
		// 传递到消息
		Dispatch("system", msg.Body, false)
	}
}

// PushExchange 发送交换机消息
func PushExchange(exchangeName string, body []byte) {
	ch, _ := global.MessageQueueClient.Channel()
	if err := ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			Body: body,
		},
	); err != nil {
		fmt.Printf("send exchange message err: %s", err)
	}
}

// PushQueue 发送队列消息
func PushQueue(queueName string, body []byte) {
	ch, _ := global.MessageQueueClient.Channel()
	err := ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			Body: body,
		})
	if err != nil {
		fmt.Printf("send queue message err: %s", err)
	}
}

// MessagePushExchange 广播
func MessagePushExchange(body []byte) {
	exchangeName := global.ServerConfig.RabbitMQServerConfig.MessageExchangeName
	PushExchange(exchangeName, body)
}
