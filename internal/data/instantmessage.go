package data

import (
	"net/url"
	"strconv"
)

type CreateInstantMessage struct {
	baseReader
	baseValidator

	Username    string   `json:"username" validate:"required_without=Usernames,min=3,max=100"`
	Usernames   []string `json:"usernames" validate:"required_without=Username,min=1"`
	ExcludeSelf bool     `json:"excludeSelf" validate:"omitempty"`
}

type ListInstantMessageMessages struct {
	baseValidator

	RoomID   *string        `json:"roomId" validate:"required_without=Username"`
	Username *string        `json:"username" validate:"required_without=RoomID"`
	Offset   *int           `json:"offset" validate:"omitempty"`
	Count    *int           `json:"count" validate:"omitempty"`
	Sort     map[string]int `json:"sort" validate:"omitempty"`
}

func (l ListInstantMessageMessages) ToQuery() string {
	params := url.Values{}

	if l.RoomID != nil {
		params.Add("roomId", url.QueryEscape(*l.RoomID))
	}

	if l.Username != nil {
		params.Add("username", url.QueryEscape(*l.Username))
	}

	if l.Offset != nil {
		offset := strconv.Itoa(*l.Offset)
		params.Add("offset", url.QueryEscape(offset))
	}

	if l.Count != nil {
		count := strconv.Itoa(*l.Count)
		params.Add("count", url.QueryEscape(count))
	}

	if l.Sort != nil {
		for k, v := range l.Sort {
			params.Add("sort", url.QueryEscape(k))

			val := strconv.Itoa(v)
			params.Add("sort", url.QueryEscape(val))
		}
	}

	return params.Encode()
}
