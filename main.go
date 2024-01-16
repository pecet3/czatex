package main

import (
	"log"
	"net/http"

	"github.com/pecet3/czatex/ws"
)

func main() {
	manager := ws.NewManager()
	index := http.FileServer(http.Dir("view"))
	http.Handle("/", index)
	http.Handle("/ws", manager)
	log.Println("Starting the server")
	log.Fatal(http.ListenAndServeTLS("0.0.0.0:8080", "server.crt", "server.key", nil))
}
