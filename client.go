package rocketchat

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	authToken  string
	userID     string
	baseURL    string
	httpClient *http.Client

	userService *UserService
}

func NewClient(baseURL, authToken, userID string, httpClient *http.Client) *Client {
	if httpClient != nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		authToken:  authToken,
		userID:     userID,
		baseURL:    baseURL,
		httpClient: httpClient,
	}

	client.userService = NewUserService(client)

	return client
}

func (c *Client) Request(ctx context.Context, method, path string, body io.Reader) (map[string]any, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Auth-Token", c.authToken)
	req.Header.Set("X-User-Id", c.userID)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := handleResponseBody(resp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) Users() *UserService {
	return c.userService
}
