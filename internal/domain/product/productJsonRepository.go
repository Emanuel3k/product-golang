package product

import (
	"github.com/emanuel3k/product-golang/storage"
)

const path = "./storage/json/products.json"

type productJsonRepository struct {
	products []*Product
}

func (pjr productJsonRepository) GetAll() ([]*Product, error) {
	var err error
	pjr.products, err = storage.ReadJson[Product](path)

	if err != nil {
		// todo
	}

	return pjr.products, nil
}
