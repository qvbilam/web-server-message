package business

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"message/enum"
)

type SocketClient struct {
	UserId           int64
	Conn             *websocket.Conn
	UserMessageQueue chan []byte
}

var UserClient map[int64]*SocketClient = make(map[int64]*SocketClient, 0)

// SendUser 用户消息推送到队列
func SendUser(userId int64, message []byte) {
	client, exists := UserClient[userId]

	if exists {
		client.UserMessageQueue <- message
		fmt.Printf("[info]尝试向本机用户: %d, 发送消息\n", userId)
		return
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
		//Dispatch(client.UserId, data, true)
		SendUser(client.UserId, data)
	}
}

func Dispatch(data []byte) {
	fmt.Println("分发原始数据", string(data))
	type Content struct {
		SenderId int64  `json:"sender_id"`
		TargetId int64  `json:"target_id"`
		Type     string `json:"type"`
	}
	c := Content{}
	err := json.Unmarshal(data, &c)
	if err != nil {
		zap.S().Errorf("错误的消息类型")
		return
	}
	switch c.Type {
	case enum.ObjTypePrivate: // 私聊消息
		SendUser(c.TargetId, data)
	default: // 默认私聊消息
		SendUser(c.TargetId, data)
	}
}

//func Dispatch(senderId int64, data []byte) {
//	//fmt.Println("接受到数据", string(data))
//
//	// 解析数据
//	type objType struct {
//		Type string `json:"type"`
//	}
//
//	t := objType{}
//	err := json.Unmarshal(data, &t)
//	if err != nil { //类型错误: 缺少type
//		t.Type = "lack-type"
//	}
//
//	switch t.Type {
//	case enum.ObjTypePrivate:
//		o := resource.PrivateObject{}
//		d := o.Decode(data)
//
//		fmt.Printf("向用户发送私聊消息: %s\n", d.TargetId)
//		SendUser(d.TargetId, data, isLocal)
//
//	default:
//		fmt.Printf("用户发送的消息类型不支持: %s\n", senderId)
//		SendUser(senderId, []byte(fmt.Sprintf("暂不支持的消息类型, 内容: %s,", data)), isLocal)
//	}
//}
