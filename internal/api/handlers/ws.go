package handlers

import (
	"fmt"
	"main/internal/ws"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type ResponseError struct {
	Message string
}

func Ws(r *gin.Context) {

	queryPa := r.Query("id")

	id, err := strconv.Atoi(queryPa)

	if err != nil {
		r.JSON(400, ResponseError{Message: fmt.Sprintf("Error request Json: %w", err)})
		return
	}

	conn, err := upgrader.Upgrade(r.Writer, r.Request, nil)
	if err != nil {
		r.JSON(400, ResponseError{Message: fmt.Sprintf("Error request Json: %w", err)})
		return
	}

	ws.Mutex.Lock()
	ws.Map_clients[id] = conn
	ws.Mutex.Unlock()

	defer func() {
		ws.Mutex.Lock()
		delete(ws.Map_clients, id)
		ws.Mutex.Unlock()
		conn.Close()
	}()

	select {}
}
