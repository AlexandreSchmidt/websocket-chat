package hub

import "time"

type messageResponse struct {
	ClientAlias string    `json:"clientAlias"`
	Message     string    `json:"message"`
	TimeStamps  time.Time `json:"timestamps"`
}
