package hub

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients      []client
	lastMessages [30]messageResponse
}

func (h *Hub) UpdateLastMessage(message messageResponse) {

}

func (h *Hub) AddClient(connection *websocket.Conn) client {
	client := newClient(connection)
	h.clients = append(h.clients, client)

	return client
}

func (h *Hub) WriteBroadcast(message string, client client) {
	messageResponse := messageResponse{ClientAlias: client.Alias, Message: message, TimeStamps: time.Now()}

	for i := range h.clients {
		err := h.clients[i].Connection.WriteJSON(messageResponse)
		if err != nil {
			fmt.Println(err)
		}
	}
}
