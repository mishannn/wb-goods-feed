package wildberries

import (
	"fmt"

	"example.org/wbsniper/internal/entities/product"
)

func getProductImages(p Product) []product.Image {
	if p.Pics < 1 {
		return []product.Image{}
	}

	volID := p.ID / 100000
	partID := p.ID / 1000

	var basketID int
	switch {
	case volID <= 143:
		basketID = 1
	case 144 <= volID && volID <= 287:
		basketID = 2
	case 288 <= volID && volID <= 431:
		basketID = 3
	case 432 <= volID && volID <= 719:
		basketID = 4
	case 720 <= volID && volID <= 1007:
		basketID = 5
	case 1008 <= volID && volID <= 1061:
		basketID = 6
	case 1062 <= volID && volID <= 1115:
		basketID = 7
	case 1116 <= volID && volID <= 1169:
		basketID = 8
	case 1170 <= volID && volID <= 1313:
		basketID = 9
	case 1314 <= volID && volID <= 1601:
		basketID = 10
	case 1602 <= volID && volID <= 1655:
		basketID = 11
	case 1656 <= volID && volID <= 1919:
		basketID = 12
	case 1920 <= volID && volID <= 2045:
		basketID = 13
	case 2046 <= volID && volID <= 2189:
		basketID = 14
	case 2190 <= volID && volID <= 2405:
		basketID = 15
	default:
		basketID = 16
	}

	images := make([]product.Image, 0, p.Pics)
	for i := 1; i <= int(p.Pics); i++ {
		images = append(images, product.Image{
			URL: fmt.Sprintf("https://basket-%02d.wbbasket.ru/vol%d/part%d/%d/images/big/%d.jpg", basketID, volID, partID, p.ID, i),
		})
	}

	return images
}
