package services

import (
	"github.com/emanuel3k/product-golang/internal/domain"
)

type productService struct {
	productRepository domain.IRepository
}

func (ps *productService) GetAll() ([]*domain.Product, error) {
	res, err := ps.productRepository.GetAll()

	if err != nil {
		// todo
	}

	return res, nil
}

func (ps *productService) GetById(productId int) (*domain.Product, error) {
	res, err := ps.productRepository.GetById(productId)
	if err != nil {
		// todo
	}

	return res, nil
}

func (ps *productService) Create(body domain.CreateBodyRequest) (*domain.Product, error) {
	exists, err := ps.productRepository.GetByCodeValue(body.CodeValue)
	if err != nil {
		// todo
		return nil, nil
	}

	if exists != nil {
		// todo
		return nil, nil
	}

	product := body.CreateToDomain()

	if err := ps.productRepository.Create(&product); err != nil {
		// todo
	}

	return &product, nil
}

func (ps *productService) DeleteById(productId int) error {
	return ps.productRepository.DeleteById(productId)
}

func (ps *productService) UpdateById(productId int, body domain.UpdateBodyRequest) (*domain.Product, error) {

	if body.CodeValue != nil {
		exists, err := ps.productRepository.GetByCodeValue(*body.CodeValue)
		if err != nil {
			// todo
			return nil, nil
		}

		if exists != nil && exists.ID != productId {
			// todo
		}
	}

	return ps.productRepository.UpdateById(productId, body)
}

func NewService(productRepository domain.IRepository) domain.IService {
	return &productService{productRepository}
}
