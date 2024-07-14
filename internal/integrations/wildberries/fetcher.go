package wildberries

import (
	"fmt"

	"example.org/wbsniper/internal/entities/product"
)

type Fetcher struct {
	client *Client
}

func NewFetcher() *Fetcher {
	client := &Client{}

	return &Fetcher{
		client: client,
	}
}

func (f *Fetcher) GetProducts() ([]product.Product, error) {
	wbProducts, err := f.client.GetProducts()
	if err != nil {
		return nil, fmt.Errorf("can't get products: %w", err)
	}

	products := make([]product.Product, 0, len(wbProducts.Data.Products))
	for _, wbProduct := range wbProducts.Data.Products {
		products = append(products, product.Product{
			Name:        wbProduct.Name,
			Brand:       wbProduct.Brand,
			Rating:      float32(wbProduct.ReviewRating),
			ReviewCount: int(wbProduct.Feedbacks),
			Link:        fmt.Sprintf("https://www.wildberries.ru/catalog/%d/detail.aspx", wbProduct.ID),
		})
	}

	return products, nil
}
