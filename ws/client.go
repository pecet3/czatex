package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	conn    *websocket.Conn
	receive chan []byte
	room    *room
	name 	string
}

func (c *client) read() {
	defer c.conn.Close()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.conn.Close()

	for msg := range c.receive {
		msgWithName := []byte(c.name + ": " + string(msg))
		log.Println(msgWithName)
		err := c.conn.WriteMessage(websocket.TextMessage, msgWithName)
		if err != nil {
			return
		}
	}
}
