package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

var Map_clients = make(map[int]*websocket.Conn)
var Mutex = &sync.Mutex{}
