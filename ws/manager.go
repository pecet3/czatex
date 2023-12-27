package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type manager struct {
	rooms map[*room]bool
	mutex sync.Mutex
}

var (
	upgrader = &websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func checkOrigin(r *http.Request) bool {
	return true
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("New connection")
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &client{
		conn:    conn,
		receive: make(chan []byte),
		room:    r,
	}

	r.join <- client

	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
