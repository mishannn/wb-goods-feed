package wildberries

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"example.org/wbsniper/internal/entities/product"
)

type PriceHistoryItem struct {
	Dt    int64 `json:"dt"`
	Price Price `json:"price"`
}

type Price struct {
	RUB int64 `json:"RUB"`
}

func getProductPriceHistory(wbProduct Product) ([]product.PriceHistoryItem, error) {
	basketURL := getProductBasketURL(wbProduct)

	resp, err := http.Get(fmt.Sprintf("%s/info/price-history.json", basketURL))
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

	var respBody []PriceHistoryItem
	err = json.Unmarshal(respBodyBytes, &respBody)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal response body: %w", err)
	}

	result := make([]product.PriceHistoryItem, 0, len(respBody))
	for _, item := range respBody {
		result = append(result, product.PriceHistoryItem{
			Date:  time.Unix(item.Dt, 0),
			Price: product.Price(item.Price),
		})
	}

	return result, nil
}