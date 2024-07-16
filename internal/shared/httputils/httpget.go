package httputils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HttpGet[T any](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't do request: %w", err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server sent unexpected code %d, body: %s", resp.StatusCode, respBodyBytes)
	}

	var respBody T
	err = json.Unmarshal(respBodyBytes, &respBody)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal response body: %w", err)
	}

	return &respBody, nil
}
