package ws

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pecet3/czatex/utils"
)

type room struct {
	name string

	clients map[*client]bool

	join chan *client

	leave chan *client

	forward chan []byte
}

func NewRoom(name string) *room {

	return &room{
		name:    name,
		clients: make(map[*client]bool),
		join:    make(chan *client),
		leave:   make(chan *client),
		forward: make(chan []byte),
	}
}

func (r *room) Run(m *manager) {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			var wg sync.WaitGroup
			namesChan := make(chan []string)

			wg.Add(1)
			go createNamesArr(r.clients, &wg, namesChan)

			namesArr := <-namesChan

			close(namesChan)
			serverMsg := client.name + " dołączył do pokoju " + r.name
			jsonMessage, err := utils.MarshalJsonMessage("serwer", serverMsg, namesArr)

			if err != nil {
				return
			}

			for roomClient := range r.clients {
				roomClient.conn.WriteMessage(websocket.TextMessage, jsonMessage)
			}
		case client := <-r.leave:
			var wg sync.WaitGroup
			namesChan := make(chan []string)

			wg.Add(1)
			delete(r.clients, client)
			go createNamesArr(r.clients, &wg, namesChan)

			namesArr := <-namesChan
			close(namesChan)

			serverMsg := client.name + " wyszedł z pokoju " + r.name
			jsonMessage, err := utils.MarshalJsonMessage("serwer", serverMsg, namesArr)

			if err == nil {
				for roomClient := range r.clients {

					roomClient.conn.WriteMessage(websocket.TextMessage, jsonMessage)
				}
			}

			delete(r.clients, client)
			close(client.receive)

			if len(r.clients) == 0 {
				m.RemoveRoom(r.name)
				return
			}
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}
		}
	}
}

func createNamesArr(clients map[*client]bool, wg *sync.WaitGroup, namesChan chan []string) {
	defer wg.Done()
	var names []string

	for client := range clients {
		names = append(names, client.name)
	}

	namesChan <- names

}
