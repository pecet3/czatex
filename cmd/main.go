package main

import (
	"log"
	"net/http"

	"github.com/pecet3/czatex/ws"
)

func main() {
	manager := ws.NewManager()
	index := http.FileServer(http.Dir("view"))
	log.Println("heeeeeeeeeeeeeeeelo")
	http.Handle("/", index)
	http.Handle("/ws", manager)
	log.Println("Starting the server")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
