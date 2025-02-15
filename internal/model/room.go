package model

type Room struct {
	T         string   `json:"t"`
	RoomID    string   `json:"rid"`
	Usernames []string `json:"usernames"`
}
