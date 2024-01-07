package ws

import (
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type manager struct {
	rooms map[string]*room
	mutex sync.Mutex
}

func (m *manager) RemoveRoom(name string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if room, ok := m.rooms[name]; ok {
		close(room.join)
		close(room.forward)
		close(room.leave)
		delete(m.rooms, name)
		log.Println("Closing a room with name:", room.name)
		return
	}
}
func NewManager() *manager {
	return &manager{
		rooms: make(map[string]*room),
	}
}

func (m *manager) GetRoom(name string) *room {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.rooms[name]
}

func (m *manager) CreateRoom(name string) *room {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if existingRoom, ok := m.rooms[name]; ok {
		return existingRoom
	}

	newRoom := NewRoom(name)
	m.rooms[name] = newRoom
	log.Println("Created a room with name:", name)
	return newRoom
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

func (m *manager) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	room := req.URL.Query().Get("room")
	if room == "" {
		log.Println("User connected but has'nt provided room name")
		return
	}

	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	log.Println("New connection, ip:", ip)

	currentRoom := m.GetRoom(room)

	if currentRoom == nil {
		currentRoom = m.CreateRoom(room)
		go currentRoom.Run(m)
	}

	client := &client{
		conn:    conn,
		receive: make(chan []byte),
		room:    currentRoom,
	}

	currentRoom.join <- client

	defer func() { currentRoom.leave <- client }()
	go client.write()
	client.read()
}
