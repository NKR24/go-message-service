package model

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}
