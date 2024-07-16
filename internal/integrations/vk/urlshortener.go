package vk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type URLShortener struct {
	accessToken string
}

func NewURLShortener(accessToken string) *URLShortener {
	return &URLShortener{
		accessToken: accessToken,
	}
}

type VKAPIResponse[T any] struct {
	Response *T          `json:"response"`
	Error    *VKAPIError `json:"error"`
}

type VKAPIError struct {
	ErrorCode    int64  `json:"error_code"`
	ErrorMessage string `json:"error_msg"`
}

type GetShortURLResponse struct {
	Key      string `json:"key"`
	ShortURL string `json:"short_url"`
	URL      string `json:"url"`
}

func (us *URLShortener) GetShortURL(longURL string) (string, error) {
	form := url.Values{}
	form.Add("url", longURL)
	form.Add("private", "0")
	form.Add("access_token", us.accessToken)
	form.Add("v", "5.199")

	resp, err := http.Post("https://api.vk.com/method/utils.getShortLink", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("can't do request: %w", err)
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("can't read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("server sent unexpected code %d, body: %s", resp.StatusCode, respBodyBytes)
	}

	var respBody VKAPIResponse[GetShortURLResponse]
	err = json.Unmarshal(respBodyBytes, &respBody)
	if err != nil {
		return "", fmt.Errorf("can't unmarshal response body: %w", err)
	}

	if respBody.Error != nil {
		return "", fmt.Errorf("server sent error: %s", respBody.Error.ErrorMessage)
	}

	if respBody.Response == nil {
		return "", fmt.Errorf("server did not send either an error or a response")
	}

	return respBody.Response.ShortURL, nil
}
