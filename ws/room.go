package ws

type room struct {
	name string

	clients map[*client]bool

	join chan *client

	leave chan *client

	forward chan []byte
}

func NewRoom(name string) *room {

	return &room{
		name:    name,
		clients: make(map[*client]bool),
		join:    make(chan *client),
		leave:   make(chan *client),
		forward: make(chan []byte),
	}
}

func (r *room) Run(m *manager) {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.receive)

			if len(r.clients) == 0 {
				m.RemoveRoom(r.name)
				return
			}
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}
		}
	}
}
