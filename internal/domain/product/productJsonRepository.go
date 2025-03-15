package product

import (
	"github.com/emanuel3k/product-golang/storage"
)

const path = "./storage/json/products.json"

type productJsonRepository struct {
	products map[int]*Product
}

func (pjr *productJsonRepository) GetAll() ([]*Product, error) {
	if err := pjr.loadProducts(); err != nil {
		// todo
	}

	products := make([]*Product, 0, len(pjr.products))
	for _, p := range pjr.products {
		products = append(products, p)
	}

	return products, nil
}

func (pjr *productJsonRepository) GetById(productId int) (*Product, error) {
	if err := pjr.loadProducts(); err != nil {
		// todo
	}
	return pjr.products[productId], nil
}

func (pjr *productJsonRepository) GetByCodeValue(codeValue string) (*Product, error) {
	if err := pjr.loadProducts(); err != nil {
		// todo
	}

	for _, p := range pjr.products {
		if p.CodeValue == codeValue {
			return p, nil
		}
	}

	return nil, nil
}

func (pjr *productJsonRepository) Create(request *Product) error {
	if err := pjr.loadProducts(); err != nil {
		// todo
	}

	request.ID = len(pjr.products) + 1
	pjr.products[request.ID] = request

	products := make([]*Product, 0, len(pjr.products))
	for _, p := range pjr.products {
		products = append(products, p)
	}

	if err := storage.WriteJson(path, products); err != nil {
		// todo
	}

	return nil
}

func (pjr *productJsonRepository) loadProducts() error {
	tmp, err := storage.ReadJson[Product](path)
	if err != nil {
		// todo
	}

	pjr.products = make(map[int]*Product)
	for _, p := range tmp {
		pjr.products[p.ID] = p
	}

	return nil
}
