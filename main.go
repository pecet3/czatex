package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pecet3/czatex/ws"
)

func main() {
	manager := ws.NewManager()
	index := http.FileServer(http.Dir("view"))
	http.Handle("/", index)
	http.Handle("/ws", manager)
	log.Println("Starting the server")

	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	if certFile == "" || keyFile == "" {
		log.Fatal("CERT_FILE and KEY_FILE environment variables must be set")
	}

	log.Fatal(http.ListenAndServeTLS("0.0.0.0:8080", certFile, keyFile, nil))
}
