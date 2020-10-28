package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var pool *Pool
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

func serveWs(pool *Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("New client trying to open ws")
	conn, err := upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func SetUpWS() {
	r := gin.Default()

	pool = NewPool()
	go pool.Start()

	r.GET("/ws", func(c *gin.Context) {
		serveWs(pool, c.Writer, c.Request)
	})

	go r.Run("localhost:8080")
}

func BroadCastLatestCacluations(msgs []string) {
	pool.Broadcast <- msgs
}
