package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlexandreSchmidt/websocket-webchat/internal/hub"
	"github.com/gorilla/websocket"
)

var hubClients = &hub.Hub{}

func (s *Server) WebSocket(w http.ResponseWriter, r *http.Request) {

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		// return a error code (have no idea what code is)
		return
	}

	clientRef := hubClients.AddClient(conn)

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		hubClients.WriteBroadcast(string(message), clientRef)
	}

}
