package services

import (
	"github.com/emanuel3k/product-golang/internal/domain"
	"github.com/emanuel3k/product-golang/pkg/appError"
)

type productService struct {
	productRepository domain.IRepository
}

func (ps *productService) GetAll() ([]*domain.ResponseBody, error) {
	products, err := ps.productRepository.GetAll()

	if err != nil {
		return nil, err
	}

	res := make([]*domain.ResponseBody, 0, len(products))

	for _, product := range products {
		res = append(res, product.ToResponse())
	}

	return res, nil
}

func (ps *productService) GetById(productId int) (*domain.ResponseBody, error) {
	res, err := ps.productRepository.GetById(productId)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.ToResponse(), nil
}

func (ps *productService) Create(body domain.CreateBodyRequest) (*domain.ResponseBody, error) {
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

	return product.ToResponse(), nil
}

func (ps *productService) DeleteById(productId int) error {
	return ps.productRepository.DeleteById(productId)
}

func (ps *productService) UpdateById(productId int, body domain.UpdateBodyRequest) (*domain.ResponseBody, error) {
	exists, err := ps.productRepository.GetById(productId)
	if err != nil {
		return nil, err
	}

	if exists == nil {
		return nil, appError.NotFound("product not found")
	}

	if body.CodeValue != nil {
		exists, err := ps.productRepository.GetByCodeValue(*body.CodeValue)
		if err != nil {
			return nil, err
		}

		if exists != nil && exists.ID != productId {
			return nil, appError.Conflict("code value already in use")
		}
	}

	product := body.UpdateToDomain()

	if err := ps.productRepository.UpdateById(productId, product); err != nil {
		return nil, err
	}

	return product.ToResponse(), nil
}

func NewService(productRepository domain.IRepository) domain.IService {
	return &productService{productRepository}
}
