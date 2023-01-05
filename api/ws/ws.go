package ws

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	proto "message/api/qvbilam/message/v1"
	"message/business"
	"message/enum"
	"message/global"
	"net/http"
	"strconv"
)

var u = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handel(ctx *gin.Context) {
	paramUserId := ctx.Query("u")
	userIdInt, _ := strconv.Atoi(paramUserId)
	userId := int64(userIdInt)

	conn, err := u.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("链接错误")
		return
	}
	// 防止退出
	ch := make(chan int)

	// 关闭连接
	defer func(conn *websocket.Conn) {
		_ = conn.Close()
	}(conn)

	fmt.Printf("用户: %s 链接成功\n", paramUserId)

	// 创建频道
	business.UserClient[userId] = &business.SocketClient{
		UserId:           userId,
		Conn:             conn,
		UserMessageQueue: make(chan []byte, 50),
	}
	useClient := business.UserClient[userId]

	// 发送数据
	go business.Write(useClient)
	// 接受数据
	go business.Accept(business.UserClient[userId])

	// 发送用户信息
	fmt.Printf("准备发送消息")
	_, err = global.MessageServerClient.CreatePrivateMessage(context.Background(), &proto.CreatePrivateRequest{
		UserId:       0,
		TargetUserId: userId,
		Message: &proto.MessageRequest{
			Type:    enum.TipMsgType,
			Content: "连接成功",
		},
	})

	if err != nil {
		zap.S().Errorf("发送消息失败: %s", err.Error())
	}
	fmt.Printf("用户: %d, 向用户: %d 发送消息\n", 1, userId)

	<-ch
}
