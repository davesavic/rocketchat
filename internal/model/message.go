package model

type Message struct{}

type MessageList struct {
	Messages []Message `json:"messages"`
	Count    int       `json:"count"`
	Offset   int       `json:"offset"`
	Total    int       `json:"total"`
}
