package hub

import (
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	Connection *websocket.Conn
	Alias      string
	Created    time.Time
}

func newClient(connection *websocket.Conn) client {

	c := client{
		Connection: connection,
		Alias:      randomAlias(),
		Created:    time.Now(),
	}

	return c
}

func randomAlias() string {
	alias := "User_"

	return alias

}
