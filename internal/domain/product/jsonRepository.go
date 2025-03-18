package product

import (
	"github.com/emanuel3k/product-golang/internal/domain"
	"github.com/emanuel3k/product-golang/storage"
	"math/rand"
)

const path = "./storage/json/products.json"

type JsonRepository struct {
	products map[int]*domain.Product
}

func (pjr *JsonRepository) GetAll() ([]*domain.Product, error) {
	if err := pjr.loadProducts(); err != nil {
		// todo
	}

	products := make([]*domain.Product, 0, len(pjr.products))
	for _, p := range pjr.products {
		products = append(products, p)
	}

	return products, nil
}

func (pjr *JsonRepository) GetById(productId int) (*domain.Product, error) {
	if err := pjr.loadProducts(); err != nil {
		// todo
	}
	return pjr.products[productId], nil
}

func (pjr *JsonRepository) GetByCodeValue(codeValue string) (*domain.Product, error) {
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

func (pjr *JsonRepository) Create(request *domain.Product) error {
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

func (pjr *JsonRepository) DeleteById(productId int) error {
	if err := pjr.loadProducts(); err != nil {
		// todo
	}

	delete(pjr.products, productId)

	if err := pjr.updateJson(); err != nil {
		// todo
	}

	return nil
}

func (pjr *JsonRepository) UpdateById(productId int, body domain.UpdateBodyRequest) (*domain.Product, error) {
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

func (pjr *JsonRepository) loadProducts() error {
	tmp, err := storage.ReadJson[domain.Product](path)
	if err != nil {
		// todo
	}

	pjr.products = make(map[int]*domain.Product)
	for _, p := range tmp {
		pjr.products[p.ID] = p
	}

	return nil
}

func (pjr *JsonRepository) updateJson() error {
	var products []*domain.Product
	for _, p := range pjr.products {
		products = append(products, p)
	}
	if err := storage.WriteJson(path, products); err != nil {
		// todo
	}

	return nil
}
