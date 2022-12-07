package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/streadway/amqp"
	"message/config"
)

var (
	Trans              ut.Translator // 表单验证
	ServerConfig       *config.ServerConfig
	MessageQueueClient *amqp.Connection
)
