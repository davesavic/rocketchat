package rocketchat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserService struct {
	client *Client
}

func NewUserService(client *Client) *UserService {
	return &UserService{
		client: client,
	}
}

func (u UserService) Create(ctx context.Context, d CreateUser) (*User, error) {
	if err := d.Validate(); err != nil {
		return nil, err
	}

	payload, err := d.ToReader()
	if err != nil {
		return nil, err
	}

	data, err := u.client.Request(ctx, http.MethodPost, "/api/v1/users.create", payload)
	if err != nil {
		return nil, err
	}

	userData, err := json.Marshal(data["user"])
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(userData, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserService) Login(ctx context.Context, d LoginUser) (*AuthenticatedUser, error) {
	if err := d.Validate(); err != nil {
		return nil, err
	}

	payload, err := d.ToReader()
	if err != nil {
		return nil, err
	}

	data, err := u.client.Request(ctx, http.MethodPost, "/api/v1/login", payload)
	if err != nil {
		return nil, err
	}

	userData, err := json.Marshal(data["me"])
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(userData, &user); err != nil {
		return nil, err
	}

	authToken, ok := data["authToken"].(string)
	if !ok {
		return nil, fmt.Errorf("could not parse authToken")
	}

	userID, ok := data["userId"].(string)
	if !ok {
		return nil, fmt.Errorf("could not parse userId")
	}

	return &AuthenticatedUser{
		AuthToken: authToken,
		UserID:    userID,
		User:      user,
	}, nil
}
