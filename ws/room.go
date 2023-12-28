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

func (r *room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.receive)

			if len(r.clients) == 0 {
				close(r.join)
				close(r.forward)
				close(r.leave)
				return
			}
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}
		}
	}
}
