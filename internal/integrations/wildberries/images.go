package wildberries

import (
	"fmt"

	"github.com/mishannn/wb-goods-feed/internal/entities/product"
)

func getProductImages(wbProduct Product) []product.Image {
	if wbProduct.Pics < 1 {
		return []product.Image{}
	}

	basketURL := getProductBasketURL(wbProduct)

	images := make([]product.Image, 0, wbProduct.Pics)
	for i := 1; i <= int(wbProduct.Pics); i++ {
		images = append(images, product.Image{
			URL: fmt.Sprintf("%s/images/big/%d.jpg", basketURL, i),
		})
	}

	return images
}
