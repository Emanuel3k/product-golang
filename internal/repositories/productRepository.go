package repositories

import (
	"database/sql"
	"github.com/emanuel3k/product-golang/internal/domain"
	"github.com/emanuel3k/product-golang/internal/domain/product"
)

func NewRepository(postgres *sql.DB) domain.IRepository {
	return &product.PostgresRepository{
		Conn: postgres,
	}
}
