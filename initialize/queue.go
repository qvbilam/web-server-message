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
	"time"
)

func InitQueue() {
	user := "admin"
	password := "admin"
	host := "127.0.0.1"
	port := 5672

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port)
	conn, err := amqp.Dial(url)
	if err != nil {
		//zap.S().Fatalf("%s dial error: %s", "队列服务", err)
		//panic(any(err))
		zap.S().Errorf("%s dial error: %s", "队列服务", err)
		return
	}

	global.MessageQueueClient = conn
	if global.MessageQueueClient.IsClosed() {
		return
	}

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

func InitQueueHealth() {

	// 监听状态
	closeChan := make(chan *amqp.Error, 1)
	notifyClose := global.MessageQueueClient.NotifyClose(closeChan)

	// 健康检查
	timer := time.NewTimer(5 * time.Second)

	for {
		select {
		case e := <-notifyClose:
			if e != nil {
				fmt.Printf("chan通道错误,e:%+v\n", e)
				InitQueue()
			}
			//close(closeChan)
			//InitQueue()
		case <-timer.C:
			//timer.Reset(time.Second * time.Duration(rand.Intn(5)))
			timer.Reset(time.Second * 5)
			if global.MessageQueueClient.IsClosed() == true {
				fmt.Printf("定期检查错误，尝试重启\n")
				InitQueue()
			}
			//fmt.Printf("定时检测rabbitMq: %d\n", time.Now().Second())
		}
	}

}
