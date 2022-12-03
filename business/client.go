package business

import "github.com/gorilla/websocket"

type SocketClient struct {
	UserId           int64
	Conn             *websocket.Conn
	UserMessageQueue chan []byte
}
