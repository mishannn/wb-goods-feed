package usecases

import (
	"fmt"

	"example.org/wbsniper/internal/entities/feed"
	"example.org/wbsniper/internal/entities/product"
)

type PostProduct struct {
	fetcher product.Fetcher
	chooser product.Chooser
	poster  feed.Poster
}

func NewPostProduct(fetcher product.Fetcher, chooser product.Chooser, poster feed.Poster) *PostProduct {
	return &PostProduct{
		fetcher: fetcher,
		chooser: chooser,
		poster:  poster,
	}
}

func (pp *PostProduct) Do() error {
	products, err := pp.fetcher.GetProducts()
	if err != nil {
		return fmt.Errorf("can't get products: %w", err)
	}

	product, err := pp.chooser.ChooseProduct(products)
	if err != nil {
		return fmt.Errorf("can't choose product: %w", err)
	}

	post := feed.Post{
		Title:   fmt.Sprintf("%s от %s", product.Name, product.Brand),
		Content: fmt.Sprintf("Рейтинг ⭐️ %.1f на 💬 %d отзывов", product.Rating, product.ReviewCount),
		Link:    product.Link,
	}

	err = pp.poster.PublishPost(post)
	if err != nil {
		return fmt.Errorf("can't publish post: %w", err)
	}

	return nil
}
