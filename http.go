package rocketchat

import (
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

func handleResponseBody(resp *http.Response) (map[string]any, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	success, ok := data["success"].(bool)
	if !ok {
		return nil, fmt.Errorf("success key not found in response")
	}

	if !success {
		return nil, fmt.Errorf("request failed: %s", data["error"])
	}

	return data, nil
}
