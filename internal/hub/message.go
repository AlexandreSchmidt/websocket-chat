package hub

type Message struct {
	ClientAlias string `json:"clientAlias"`
	Message     string `json:"message"`
	TimeStamps  string `json:"timestamps"`
}
