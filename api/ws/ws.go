package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"message/business"
	"message/enum"
	"message/resource"
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
	clientUserId := strconv.Itoa(int(userId))
	business.UserClient[clientUserId] = &business.SocketClient{
		UserId:           userId,
		Conn:             conn,
		UserMessageQueue: make(chan []byte, 50),
	}
	useClient := business.UserClient[clientUserId]

	// 发送数据
	go business.Write(useClient)
	// 接受数据
	go business.Accept(business.UserClient[clientUserId])

	// 发送用户信息
	privateObj := resource.PrivateObject{
		SendUserId: "system",
		TargetId:   fmt.Sprintf("%d", userId),
		ObjectName: enum.MsgTypeTxt,
		Content: resource.Text{
			Content: "你好，欢迎链接",
			User:    resource.User{},
			Extra:   "",
		},
	}
	business.SendUser(clientUserId, privateObj.Encode(), true)

	<-ch
}
