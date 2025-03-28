package product

import (
	"database/sql"
	"errors"
	"github.com/emanuel3k/product-golang/internal/domain"
)

type PostgresRepository struct {
	Conn *sql.DB
}

func (pr PostgresRepository) GetAll() ([]*domain.Product, error) {
	rows, err := pr.Conn.Query("Select * from products")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	products := make([]*domain.Product, 0)
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (pr PostgresRepository) GetById(productId int) (*domain.Product, error) {
	var p domain.Product
	row := pr.Conn.QueryRow("Select * from products p where p.id = $1", productId)

	if err := row.Scan(&p.ID, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &p, nil
}

func (pr PostgresRepository) Create(body *domain.Product) error {
	panic("implement me")
}

func (pr PostgresRepository) GetByCodeValue(codeValue string) (*domain.Product, error) {
	var p domain.Product
	row := pr.Conn.QueryRow("Select * from products p where p.code_value = $1", codeValue)

	if err := row.Scan(&p.ID, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &p, nil
}

func (pr PostgresRepository) DeleteById(productId int) error {
	if _, err := pr.Conn.Exec("Delete from products p where p.id = $1", productId); err != nil {
		return err
	}

	return nil
}

func (pr PostgresRepository) UpdateById(productId int, body domain.UpdateBodyRequest) (*domain.Product, error) {
	panic("implement me")
}
