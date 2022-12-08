package ws

import (
	"encoding/json"
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

var userClient map[string]*business.SocketClient = make(map[string]*business.SocketClient, 0)

func Handel(ctx *gin.Context) {
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
	clientUserId := strconv.Itoa(int(userId))
	userClient[clientUserId] = &business.SocketClient{
		UserId:           userId,
		Conn:             conn,
		UserMessageQueue: make(chan []byte, 50),
	}

	// 发送数据
	go write(userClient[clientUserId])
	// 接受数据
	go accept(userClient[clientUserId])

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
	sendUser(clientUserId, privateObj.Encode())

	<-ch
}

// 用户消息推送到队列
func sendUser(userId string, message []byte) {
	client, exists := userClient[userId]
	if exists {
		client.UserMessageQueue <- message
	} else {

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
		senderId := strconv.Itoa(int(client.UserId))
		dispatch(senderId, data)
	}
}

func dispatch(senderId string, data []byte) {
	fmt.Println("接受到数据", string(data))

	// 解析数据
	type objType struct {
		Type string `json:"type"`
	}

	t := objType{}
	err := json.Unmarshal(data, &t)
	if err != nil { //类型错误: 缺少type
		t.Type = "lack-type"
	}

	switch t.Type {
	case enum.ObjTypePrivate:
		o := resource.PrivateObject{}
		d := o.Decode(data)

		sendUser(d.TargetId, data)
	default:
		sendUser(senderId, []byte(fmt.Sprintf("暂不支持的消息类型, 内容: %s,", data)))
	}

	//toUser := message.TargetId
	//target, _ := strconv.Atoi(toUser)
	//
	//fmt.Printf("接受消息: %d", target)
	//sendUser(int64(target), data)
}
