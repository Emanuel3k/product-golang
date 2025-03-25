package services

import (
	"github.com/emanuel3k/product-golang/internal/domain"
	"github.com/emanuel3k/product-golang/pkg/appError"
)

type productService struct {
	productRepository domain.IRepository
}

func (ps *productService) GetAll() ([]*domain.Product, error) {
	res, err := ps.productRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ps *productService) GetById(productId int) (*domain.Product, error) {
	res, err := ps.productRepository.GetById(productId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ps *productService) Create(body domain.CreateBodyRequest) (*domain.Product, error) {
	exists, err := ps.productRepository.GetByCodeValue(body.CodeValue)
	if err != nil {
		return nil, err
	}

	if exists != nil {
		return nil, appError.Conflict("code value already in use")
	}

	product := body.CreateToDomain()

	if err := ps.productRepository.Create(&product); err != nil {
		return nil, err
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
			return nil, appError.Conflict("code value already in use")
		}
	}

	return ps.productRepository.UpdateById(productId, body)
}

func NewService(productRepository domain.IRepository) domain.IService {
	return &productService{productRepository}
}
