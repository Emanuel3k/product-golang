package product

import (
	"github.com/emanuel3k/product-golang/storage"
	"math/rand"
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

	request.ID = rand.Int()
	pjr.products[request.ID] = request

	if err := pjr.updateJson(); err != nil {
		// todo
	}

	return nil
}

func (pjr *productJsonRepository) DeleteById(productId int) error {
	if err := pjr.loadProducts(); err != nil {
		// todo
	}

	delete(pjr.products, productId)

	if err := pjr.updateJson(); err != nil {
		// todo
	}

	return nil
}

func (pjr *productJsonRepository) UpdateById(productId int, body UpdateBodyRequest) (*Product, error) {
	if err := pjr.loadProducts(); err != nil {
		// todo
	}

	product := pjr.products[productId]

	if body.Name != nil {
		product.Name = *body.Name
	}
	if body.Quantity != nil {
		product.Quantity = *body.Quantity
	}
	if body.CodeValue != nil {
		product.CodeValue = *body.CodeValue
	}
	if body.IsPublished != nil {
		product.IsPublished = *body.IsPublished
	}
	if body.Expiration != nil {
		product.Expiration = *body.Expiration
	}
	if body.Price != nil {
		product.Price = *body.Price
	}

	if err := pjr.updateJson(); err != nil {
		// todo
	}

	return product, nil
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

func (pjr *productJsonRepository) updateJson() error {
	var products []*Product
	for _, p := range pjr.products {
		products = append(products, p)
	}
	if err := storage.WriteJson(path, products); err != nil {
		// todo
	}

	return nil
}
