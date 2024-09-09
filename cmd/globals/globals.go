package globals

import "github.com/gorilla/websocket"

var AttackChan = make(chan string)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
