package business

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"message/enum"
	"message/resource"
	"strconv"
)

type SocketClient struct {
	UserId           int64
	Conn             *websocket.Conn
	UserMessageQueue chan []byte
}

var UserClient map[string]*SocketClient = make(map[string]*SocketClient, 0)

// SendUser 用户消息推送到队列
func SendUser(userId string, message []byte, isLocal bool) {
	client, exists := UserClient[userId]

	fmt.Printf("[info]webscoket连接用户: %v\n", getClientUserIds())

	if exists {
		client.UserMessageQueue <- message
		fmt.Printf("[info]尝试向本机用户: %s, 发送消息\n", userId)
		return
	}

	// 本地消息推送到交换机中
	if isLocal {
		fmt.Printf("[info]接受用户: %s, 未连到本地，消息推送至交换机", userId)
		MessagePushExchange(message)
	}
}

// Write 推送消息
func Write(client *SocketClient) {
	for {
		select {
		case msg := <-client.UserMessageQueue:
			if err := client.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Printf("[error]发送消息失败，内容: %s\n", msg)
				return
			}
		}
	}
}

// Accept 接受消息
func Accept(client *SocketClient) {
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
		Dispatch(senderId, data, true)
	}
}

func Dispatch(senderId string, data []byte, isLocal bool) {
	//fmt.Println("接受到数据", string(data))

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

		fmt.Printf("向用户发送私聊消息: %s\n", d.TargetId)
		SendUser(d.TargetId, data, isLocal)
	default:
		fmt.Printf("用户发送的消息类型不支持: %s\n", senderId)
		SendUser(senderId, []byte(fmt.Sprintf("暂不支持的消息类型, 内容: %s,", data)), isLocal)
	}
}

func getClientUserIds() []int64 {
	var res []int64
	for k, _ := range UserClient {
		u, _ := strconv.Atoi(k)
		res = append(res, int64(u))
	}
	return res
}
