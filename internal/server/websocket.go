package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AlexandreSchmidt/websocket-webchat/internal/hub"
	"github.com/gorilla/websocket"
)

var hubClients []hub.Client

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

	alias := "User_" + strconv.Itoa(len(hubClients))
	client := &hub.Client{Connection: conn, Alias: alias}

	hubClients = append(hubClients, *client)

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		timestampsMessage := time.Now().Format(time.RFC3339)
		messageResponse := MessageResponse{ClientAlias: client.Alias, Message: string(message), TimeStamps: timestampsMessage}

		for i := range hubClients {
			err := hubClients[i].Connection.WriteJSON(messageResponse)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}
