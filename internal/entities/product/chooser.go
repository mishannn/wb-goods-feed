package product

import (
	"fmt"
	"math/rand"
)

type Chooser interface {
	ChooseProduct(products []Product) (*Product, error)
}

type DefaultChooser struct{}

func (dcs *DefaultChooser) ChooseProduct(products []Product) (*Product, error) {
	if len(products) == 0 {
		return nil, fmt.Errorf("no products to choose from")
	}

	return &products[rand.Intn(len(products))], nil
}
