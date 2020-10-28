package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/jj/repo/calculator-app/server/models"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	List []models.CalculatedResult
}

// Read fn is used only for detecting when client closes the connection or error occurs in ws connection
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(messageType, p)
	}
}
