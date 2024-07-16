package wildberries

import (
	"fmt"
	"net/url"

	"github.com/mishannn/wb-goods-feed/internal/shared/httputils"
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

	productsURL := fmt.Sprintf("https://recom.wb.ru/personal/ru/common/v5/search?%s", qs.Encode())

	respBody, err := httputils.HttpGet[Response[GetProductsResponseData]](productsURL)
	if err != nil {
		return nil, fmt.Errorf("can't get recommended products: %w", err)
	}

	return respBody, nil
}
