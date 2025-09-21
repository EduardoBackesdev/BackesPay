package handlers

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type wsRequest struct {
	Id int
}

type ResponseError struct {
	Message string
}

var map_clients = make(map[int]*websocket.Conn)
var mutex = &sync.Mutex{}

func Ws(r *gin.Context) {

	var a wsRequest

	if err := r.ShouldBindJSON(&a); err != nil {
		r.JSON(400, ResponseError{Message: fmt.Sprintf("Error request Json: %w", err)})
		return
	}

	id := a.Id

	conn, err := upgrader.Upgrade(r.Writer, r.Request, nil)
	if err != nil {
		return
	}

	mutex.Lock()
	map_clients[id] = conn
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(map_clients, id)
		mutex.Unlock()
		conn.Close()
	}()

	select {}
}
