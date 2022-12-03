package global

import (
	ut "github.com/go-playground/universal-translator"
	userProto "video/api/qvbilam/user/v1"
	proto "video/api/qvbilam/video/v1"
	"video/config"
)

var (
	Trans        ut.Translator // 表单验证
	ServerConfig *config.ServerConfig

	CategoryServerClient proto.CategoryClient
	DramaServerClient    proto.DramaClient
	RegionServerClient   proto.RegionClient
	VideoServerClient    proto.VideoClient
	UserServerClient     userProto.UserClient
)
