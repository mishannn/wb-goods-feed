package product

import "time"

type Product struct {
	Name         string
	Brand        string
	Rating       float32
	ReviewCount  int
	Images       []Image
	PriceHistory []PriceHistoryItem
	Link         string
}

type Image struct {
	URL string
}

type PriceHistoryItem struct {
	Date  time.Time
	Price Price
}

type Price struct {
	RUB int64 // Копейки
}
