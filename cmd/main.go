package main

import (
	"log"
	"net/http"

	"github.com/pecet3/czatex/ws"
)

func main() {
	room := ws.NewRoom()
	http.Handle("/ws", room)

	go room.Run()

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
