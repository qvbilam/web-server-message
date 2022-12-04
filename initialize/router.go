package initialize

import (
	"encoding/json"
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

	router.GET("/ws", func(ctx *gin.Context) {
		paramUserId := ctx.Query("u")
		fmt.Println(paramUserId)
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
		userClient[userId] = &business.SocketClient{
			UserId:           userId,
			Conn:             conn,
			UserMessageQueue: make(chan []byte, 50),
		}

		// 发送数据
		go write(userClient[userId])
		// 接受数据
		go accept(userClient[userId])

		// 发送用户信息
		sendUser(userId, []byte("你好"))
		sendUser(userId, []byte("我是你爹"))

		<-ch
	})

	return router
}

// 用户消息推送到队列
func sendUser(userId int64, message []byte) {
	client, exists := userClient[userId]
	if exists {
		client.UserMessageQueue <- message
	}
}

// 推送消息
func write(client *business.SocketClient) {
	for {
		select {
		case msg := <-client.UserMessageQueue:
			if err := client.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Printf("发送消息失败，内容: %s\n", msg)
				return
			}
		}
	}
}

// 接受消息
func accept(client *business.SocketClient) {
	for {
		_, data, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				fmt.Printf("[info]退出链接, 用户:%d: \n", client.UserId)
				break
			}
			fmt.Printf("[error]接受数据失败: %s\n", err)
			break
		}
		// 分发数据
		dispatch(data)
	}
}

func dispatch(data []byte) {
	type Message struct {
		UserId  int64  `json:"user_id"`
		Message string `json:"message"`
	}
	fmt.Println("接受到数据", string(data))
	// 解析数据
	message := Message{}
	if err := json.Unmarshal(data, &message); err != nil {
		fmt.Println("解析数据失败", err.Error())
		return
	}
	fmt.Println("解析数据为:", message)
	fmt.Println(message)
	sendUser(message.UserId, data)
}
