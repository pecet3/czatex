package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/pecet3/czatex/ws"
)

// func main() {
// 	manager := ws.NewManager()
// 	index := http.FileServer(http.Dir("view"))
// 	http.Handle("/", index)
// 	http.Handle("/ws", manager)

// 	cert, err := tls.LoadX509KeyPair("fullchain.pem",
// 		"privkey.pem")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	config := &tls.Config{
// 		Certificates: []tls.Certificate{cert},
// 	}

// 	address := "0.0.0.0:443"

// 	server := &http.Server{
// 		Addr:      address,
// 		TLSConfig: config,
// 	}

// 	log.Println("Server is running: ", address)
// 	log.Fatal(server.ListenAndServeTLS("", ""))
// }

func main() {
	manager := ws.NewManager()
	index := http.FileServer(http.Dir("view"))
	http.Handle("/", index)
	http.Handle("/ws", manager)

	cert, err := tls.LoadX509KeyPair("/etc/letsencrypt/live/czatex.pecet.it-0001/fullchain.pem",
		"/etc/letsencrypt/live/czatex.pecet.it-0001/privkey.pem")

	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	address := "0.0.0.0:8443"

	server := &http.Server{
		Addr:      address,
		TLSConfig: config,
	}

	log.Println("Server is running: ", address)
	log.Fatal(server.ListenAndServeTLS("", ""))
}
