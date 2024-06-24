package hub

import (
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Connection *websocket.Conn
	Alias      string
	Created    string
}

func (c *Client) New(connection *websocket.Conn) {

	c.Created = time.Now().Format(time.RFC3339)
}
