package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/streadway/amqp"
	proto "message/api/qvbilam/message/v1"
	userProto "message/api/qvbilam/user/v1"
	"message/config"
)

var (
	Trans               ut.Translator // 表单验证
	ServerConfig        *config.ServerConfig
	MessageQueueClient  *amqp.Connection
	MessageServerClient proto.MessageClient
	UserServerClient    userProto.UserClient
)
