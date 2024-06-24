package hub

type Hub struct {
	clients      []Client
	lastMessages [30]Message
}
