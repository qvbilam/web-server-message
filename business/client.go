package business

import "github.com/gorilla/websocket"

type SocketClient struct {
	Conn         *websocket.Conn
	MessageQueue chan []byte
}
