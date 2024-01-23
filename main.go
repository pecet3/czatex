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

	address := "0.0.0.0:8080" // Zmieniono port na standardowy dla HTTP

	server := &http.Server{
		Addr: address,
	}

	log.Println("Server is running: ", address)
	log.Fatal(server.ListenAndServe()) // Usunięto argumenty i zmieniono na ListenAndServe
}