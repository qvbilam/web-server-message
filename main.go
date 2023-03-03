package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	proto "message/api/qvbilam/message/v1"
	"message/business"
	"message/global"
	"message/initialize"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化日志
	initialize.InitLogger()
	// 初始化配置
	initialize.InitConfig()
	// 初始化路由
	Router := initialize.InitRouters()
	// 初始化表单验证
	if err := initialize.InitValidateTran("zh"); err != nil {
		zap.S().Panic("翻译器初始化失败: ", err.Error())
	}
	// 初始化grpc客户端
	initialize.InitServer()
	// 初始化队列
	initialize.InitQueue()
	go initialize.InitQueueHealth()

	Name := global.ServerConfig.Name
	Host := "0.0.0.0"
	Port := global.ServerConfig.Port

	// 启动服务
	go func() {
		zap.S().Infof("%s start listen: %s:%d", Name, Host, Port)

		if err := Router.Run(fmt.Sprintf(":%d", Port)); err != nil {
			zap.S().Panic("%s 服务启动失败: %s", Name, err.Error())
		}
	}()

	// 监听结束
	// 接受终止信号(优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 服务注销

	// 删除队列
	business.DeleteQueue(global.ServerConfig.RabbitMQServerConfig.MessageQueueName)
	if global.MessageServerClient != nil {
		_, err := global.MessageServerClient.DeleteQueue(context.Background(), &proto.UpdateQueueRequest{
			Name: global.ServerConfig.RabbitMQServerConfig.MessageQueueName,
		})
		if err != nil {
			zap.S().Infof("删除队列: %s,失败: %s", global.ServerConfig.RabbitMQServerConfig.MessageQueueName, err)
		} else {
			zap.S().Infof("删除队列: %s,成功", global.ServerConfig.RabbitMQServerConfig.MessageQueueName)
		}
	}
	zap.S().Info("服务注销成功")
}
