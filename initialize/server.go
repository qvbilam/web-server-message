package initialize

import (
	"fmt"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
	userProto "video/api/qvbilam/user/v1"
	proto "video/api/qvbilam/video/v1"
	"video/global"
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
	s := serverClientConfig{
		userDialConfig: &dialConfig{
			host: global.ServerConfig.UserServerConfig.Host,
			port: global.ServerConfig.UserServerConfig.Port,
		},
		videoDialConfig: &dialConfig{
			host: global.ServerConfig.VideoServerConfig.Host,
			port: global.ServerConfig.VideoServerConfig.Port,
		},
	}

	s.initVideoServer()
	s.initUserServer()
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

func (s *serverClientConfig) initVideoServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.videoDialConfig.host, s.videoDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", global.ServerConfig.VideoServerConfig.Name, err)
	}

	// 分类
	categoryClient := proto.NewCategoryClient(conn)
	global.CategoryServerClient = categoryClient

	// 剧集
	dramaClient := proto.NewDramaClient(conn)
	global.DramaServerClient = dramaClient

	// 区域
	regionClient := proto.NewRegionClient(conn)
	global.RegionServerClient = regionClient

	// 视频
	videoClient := proto.NewVideoClient(conn)
	global.VideoServerClient = videoClient
}

func (s *serverClientConfig) initUserServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.userDialConfig.host, s.userDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", global.ServerConfig.UserServerConfig.Name, err)
	}

	userClient := userProto.NewUserClient(conn)
	global.UserServerClient = userClient
	//u, err := global.UserServerClient.Detail(context.Background(), &userProto.GetUserRequest{Id: 1})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(u)
}
