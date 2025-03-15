package product

import (
	"github.com/emanuel3k/product-golang/storage"
)

const path = "./storage/json/products.json"

type productJsonRepository struct {
	products map[int]*Product
}

func (pjr *productJsonRepository) GetAll() ([]*Product, error) {
	pjr.loadProducts()

	products := make([]*Product, 0, len(pjr.products))

	for _, p := range pjr.products {
		products = append(products, p)
	}

	return products, nil
}

func (pjr *productJsonRepository) GetById(productId int) (*Product, error) {
	return pjr.products[productId], nil
}

func (pjr *productJsonRepository) loadProducts() {

	tmp, err := storage.ReadJson[Product](path)
	if err != nil {
		// todo
	}

	pjr.products = make(map[int]*Product)
	for _, p := range tmp {
		pjr.products[p.ID] = p
	}

}
