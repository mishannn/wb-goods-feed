package wildberries

import (
	"fmt"
	"time"

	"github.com/mishannn/wb-goods-feed/internal/entities/product"
	"github.com/mishannn/wb-goods-feed/internal/shared/httputils"
)

type priceHistoryItem struct {
	Dt    int64 `json:"dt"`
	Price price `json:"price"`
}

type price struct {
	RUB int64 `json:"RUB"`
}

func getProductPriceHistory(wbProduct Product) ([]product.PriceHistoryItem, error) {
	basketURL := getProductBasketURL(wbProduct)

	respBody, err := httputils.HttpGet[[]priceHistoryItem](fmt.Sprintf("%s/info/price-history.json", basketURL))
	if err != nil {
		return nil, fmt.Errorf("can't get price history: %w", err)
	}

	result := make([]product.PriceHistoryItem, 0, len(*respBody))
	for _, item := range *respBody {
		result = append(result, product.PriceHistoryItem{
			Date:  time.Unix(item.Dt, 0),
			Price: product.Price(item.Price),
		})
	}

	if len(wbProduct.Sizes) != 0 {
		result = append(result, product.PriceHistoryItem{
			Date: time.Now(),
			Price: product.Price{
				RUB: wbProduct.Sizes[0].Price.Total,
			},
		})
	}

	return result, nil
}
