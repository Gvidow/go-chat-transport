package entity

type Message struct {
	Username string `json:"username" example:"user-2024"`
	Time     uint64 `json:"time" example:"1711902448"`
	Content  string `json:"content" example:"hello!"`
}

type MessageWithErrorFlag struct {
	Message
	Error bool `json:"error"`
}
