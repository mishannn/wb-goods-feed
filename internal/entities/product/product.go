package product

type Image struct {
	URL string
}

type Product struct {
	Name        string
	Brand       string
	Rating      float32
	ReviewCount int
	Images      []Image
	Link        string
}
