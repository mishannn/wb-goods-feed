package product

type Fetcher interface {
	GetProducts() ([]Product, error)
}
