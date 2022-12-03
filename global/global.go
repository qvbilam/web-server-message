package global

import (
	ut "github.com/go-playground/universal-translator"
	"message/config"
)

var (
	Trans        ut.Translator // 表单验证
	ServerConfig *config.ServerConfig
)
