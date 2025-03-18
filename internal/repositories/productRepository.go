package repositories

import (
	"github.com/emanuel3k/product-golang/internal/domain"
	"github.com/emanuel3k/product-golang/internal/domain/product"
)

func NewRepository() domain.IRepository {
	return &product.JsonRepository{}
}
