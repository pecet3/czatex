package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/pecet3/czatex/ws"
)

func main() {
	manager := ws.NewManager()
	index := http.FileServer(http.Dir("view"))
	log.Println(index)
	http.HandleFunc("/", func(w http.ResponseWrite, r *http.Request){
		log.Println("new conn")
		fmt.Println("hello")
		w.Write(index)
	})

	http.Handle("/ws", manager)
	log.Println("Starting the server")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
