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

// SendUser 用户消息推送到队列 {"send_user_id":3,"target_id":4,"type":"private","content":{"type":"TextMsg","content":"你好啊， 我是3","user":{"id":1,"code":5721,"nickname":"QvBiLam106907","avatar":"https://blogupy.qvbilam.xin/bg/6666.JPG","gender":"male","extra":""},"extra":""}}
func SendUser(userId int64, message []byte) {
	client, exists := UserClient[userId]

	if exists {
		client.UserMessageQueue <- message
		fmt.Printf("[info]尝试向本机用户: %d, 发送消息\n", userId)
		return
	}
}

func SendRoom(roomId int64, message []byte) {
	// todo 获取房间的用户
	// 循环 sendUser
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
				// todo 处理用户退出连接
				fmt.Printf("[info]退出链接, 用户:%d: \n", client.UserId)
				break
			}
			fmt.Printf("[error]接受数据失败: %s\n", err)
			break
		}
		// 分发数据
		//Dispatch(client.UserId, data, true)
		Dispatch(data)
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
