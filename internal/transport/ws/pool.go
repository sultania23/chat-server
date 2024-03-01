package ws

import (
	"fmt"
	. "github.com/google/uuid"
	"github.com/tuxoo/idler/internal/model/entity"
)

type Pool struct {
	id         UUID
	clients    map[*Client]bool
	broadcast  chan entity.Message
	register   chan *Client
	unregister chan *Client
}

func NewPool(id UUID) *Pool {
	return &Pool{
		id:         id,
		broadcast:  make(chan entity.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Pool) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				if client.user == message.Sender {
					continue
				}

				if err := client.conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

func (h *Pool) Send(msg entity.Message) {
	h.broadcast <- msg
}
