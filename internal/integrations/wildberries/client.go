package wildberries

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
}

type Response[T any] struct {
	Metadata       Metadata `json:"metadata"`
	State          int64    `json:"state"`
	Version        int64    `json:"version"`
	PayloadVersion int64    `json:"payloadVersion"`
	Data           T        `json:"data"`
}

type GetProductsResponseData struct {
	Products []Product `json:"products"`
}

func (*Client) GetProducts() (*Response[GetProductsResponseData], error) {
	qs := url.Values{}
	qs.Set("ab_testing", "false")
	qs.Set("appType", "1")
	qs.Set("curr", "rub")
	qs.Set("dest", "-1257786")
	qs.Set("page", "1")
	qs.Set("query", "0")
	qs.Set("resultset", "catalog")
	qs.Set("spp", "30")
	qs.Set("suppressSpellcheck", "false")

	req, err := http.NewRequest("GET", fmt.Sprintf("https://recom.wb.ru/personal/ru/common/v5/search?%s", qs.Encode()), nil)
	if err != nil {
		return nil, fmt.Errorf("can't create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9")

	resp, err := http.DefaultClient.Do(req)
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

	var respBody Response[GetProductsResponseData]
	err = json.Unmarshal(respBodyBytes, &respBody)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal response body: %w", err)
	}

	return &respBody, nil
}
