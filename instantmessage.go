package rocketchat

import (
	"context"
	"fmt"

	"github.com/goccy/go-json"
)

type InstantMessageService struct {
	client *Client
}

func NewInstantMessageServiceeeee(client *Client) *InstantMessageService {
	return &InstantMessageService{
		client: client,
	}
}

func (dm InstantMessageService) Create(ctx context.Context, d CreateInstantMessage) (*Room, error) {
	if err := d.Validate(); err != nil {
		return nil, err
	}

	payload, err := d.ToReader()
	if err != nil {
		return nil, err
	}

	data, err := dm.client.Request(ctx, "POST", "/api/v1/im.create", payload)
	if err != nil {
		return nil, err
	}

	roomData, err := json.Marshal(data["room"])
	if err != nil {
		return nil, err
	}

	var room Room
	if err := json.Unmarshal(roomData, &room); err != nil {
		return nil, err
	}

	return &room, nil
}

func (dm InstantMessageService) ListMessages(ctx context.Context, d ListInstantMessageMessages) (*MessageList, error) {
	if err := d.Validate(); err != nil {
		return nil, err
	}

	data, err := dm.client.Request(ctx, "GET", "/api/v1/im.messages?"+d.ToQuery(), nil)
	if err != nil {
		return nil, err
	}

	messageData, err := json.Marshal(data["messages"])
	if err != nil {
		return nil, err
	}

	var messages []Message
	if err := json.Unmarshal(messageData, &messages); err != nil {
		return nil, err
	}

	count, ok := data["count"].(int)
	if !ok {
		return nil, fmt.Errorf("could not parse count")
	}

	offset, ok := data["offset"].(int)
	if !ok {
		return nil, fmt.Errorf("could not parse offset")
	}

	total, ok := data["total"].(int)
	if !ok {
		return nil, fmt.Errorf("could not parse total")
	}

	return &MessageList{
		Messages: messages,
		Count:    count,
		Offset:   offset,
		Total:    total,
	}, nil
}
