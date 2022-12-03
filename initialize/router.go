package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"message/business"
	"message/middleware"
	"net/http"
	"strconv"
)

var u = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var userClient map[int64]*business.SocketClient = make(map[int64]*business.SocketClient, 0)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	router.GET("/", func(ctx *gin.Context) {
		paramUserId := ctx.Param("user_id")
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

		// 创建频道
		fmt.Println("创建频道")
		userClient[userId] = &business.SocketClient{
			Conn: conn,
		}

		// 发送用户信息
		fmt.Println("发送消息准备")
		sendUser(userId, []byte("你好"))
		fmt.Println("发送消息完毕")

		write(userClient[userId])
		<-ch
	})

	return router
}

// 用户消息推送到队列
func sendUser(userId int64, message []byte) {
	client, exist := userClient[userId]
	fmt.Println("是否存在:", exist)
	//if exists {
	client.MessageQueue <- message
	fmt.Println("nmd")
	fmt.Println(client.MessageQueue)
	//}
}

// 推送消息
func write(client *business.SocketClient) {
	for {
		select {
		case msg := <-client.MessageQueue:
			if err := client.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Printf("发送消息失败，内容: %s\n", msg)
				return
			}
		default:
			fmt.Println("发个鸡")
		}
	}
}
