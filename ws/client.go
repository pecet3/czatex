package ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pecet3/czatex/utils"
)

type client struct {
	name    string
	conn    *websocket.Conn
	receive chan []byte
	room    *room
}

func (c *client) read() {
	defer c.conn.Close()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			continue
		}

		result,_:= utils.DecodeMessage(msg)

		

		if string(result.Message[0]) == "/"{
			
			continue
		}
		

		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.conn.Close()

	for msg := range c.receive {

		result, err := utils.DecodeMessage(msg)

		var wg sync.WaitGroup
		namesChan := make(chan []string)

		wg.Add(1)

		go createNamesArr(c.room.clients, &wg, namesChan)

		namesArr := <-namesChan
		close(namesChan)

		message, err := utils.MarshalJsonMessage(result.Name, result.Message, namesArr)

		log.Println("new message in room: ", c.room.name, string(message))
		err = c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}

