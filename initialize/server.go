package initialize

import (
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc/codes"
	"time"
)

type dialConfig struct {
	host string
	port int64
}

type serverClientConfig struct {
	userDialConfig  *dialConfig
	videoDialConfig *dialConfig
}

func InitServer() {
	//s := serverClientConfig{
	//	userDialConfig: &dialConfig{
	//		host: global.ServerConfig.UserServerConfig.Host,
	//		port: global.ServerConfig.UserServerConfig.Port,
	//	},
	//	videoDialConfig: &dialConfig{
	//		host: global.ServerConfig.VideoServerConfig.Host,
	//		port: global.ServerConfig.VideoServerConfig.Port,
	//	},
	//}

	//s.initVideoServer()
	//s.initUserServer()
}

func clientOption() []retry.CallOption {
	opts := []retry.CallOption{
		retry.WithBackoff(retry.BackoffLinear(100 * time.Millisecond)), // 重试间隔
		retry.WithMax(3), // 最大重试次数
		retry.WithPerRetryTimeout(1 * time.Second),                                 // 请求超时时间
		retry.WithCodes(codes.NotFound, codes.DeadlineExceeded, codes.Unavailable), // 指定返回码重试
	}
	return opts
}
