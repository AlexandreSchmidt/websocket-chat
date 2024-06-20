package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

type Client struct {
	Connection *websocket.Conn
	Alias      string
}
type MessageResponse struct {
	ClientAlias string `json:"clientAlias"`
	Message     string `json:"message"`
}

var hubClients []Client

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
	client := &Client{Connection: conn, Alias: alias}

	hubClients = append(hubClients, *client)

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		messageResponse := MessageResponse{ClientAlias: client.Alias, Message: string(message)}

		for i := range hubClients {
			err := hubClients[i].Connection.WriteJSON(messageResponse)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}
