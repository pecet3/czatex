package main

import (
	"log"
	"net/http"

	"github.com/pecet3/czatex/ws"
)

func main() {
	manager := ws.NewManager()

	http.Handle("/ws", manager)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
