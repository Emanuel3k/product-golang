package domain

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"codeValue"`
	IsPublished bool      `json:"isPublished"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

type CreateBodyRequest struct {
	Name        string    `json:"name" validate:"required"`
	Quantity    int       `json:"quantity" validate:"required"`
	CodeValue   string    `json:"codeValue" validate:"required"`
	IsPublished bool      `json:"isPublished"`
	Expiration  time.Time `json:"expiration" validate:"required"`
	Price       float64   `json:"price" validate:"required"`
}

type UpdateBodyRequest struct {
	Name        *string    `json:"name"`
	Quantity    *int       `json:"quantity"`
	CodeValue   *string    `json:"codeValue"`
	IsPublished *bool      `json:"isPublished"`
	Expiration  *time.Time `json:"expiration"`
	Price       *float64   `json:"price"`
}

type ResponseBody struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"codeValue"`
	IsPublished bool      `json:"isPublished"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

func (cbr *CreateBodyRequest) CreateToDomain() Product {
	return Product{
		Name:        cbr.Name,
		Quantity:    cbr.Quantity,
		CodeValue:   cbr.CodeValue,
		IsPublished: cbr.IsPublished,
		Expiration:  cbr.Expiration,
		Price:       cbr.Price,
	}
}

func (cbr *UpdateBodyRequest) UpdateToDomain() *Product {
	product := &Product{}

	if cbr.Name != nil {
		product.Name = *cbr.Name
	}
	if cbr.Quantity != nil {
		product.Quantity = *cbr.Quantity
	}
	if cbr.CodeValue != nil {
		product.CodeValue = *cbr.CodeValue
	}
	if cbr.IsPublished != nil {
		product.IsPublished = *cbr.IsPublished
	}
	if cbr.Expiration != nil {
		product.Expiration = *cbr.Expiration
	}
	if cbr.Price != nil {
		product.Price = *cbr.Price
	}

	return product
}

func (p *Product) ToResponse() *ResponseBody {
	return &ResponseBody{
		ID:          p.ID,
		Name:        p.Name,
		Quantity:    p.Quantity,
		CodeValue:   p.CodeValue,
		IsPublished: p.IsPublished,
		Expiration:  p.Expiration,
		Price:       p.Price,
	}
}

type IService interface {
	GetAll() ([]*ResponseBody, error)
	GetById(productId int) (*ResponseBody, error)
	Create(body CreateBodyRequest) (*ResponseBody, error)
	DeleteById(productId int) error
	UpdateById(productId int, body UpdateBodyRequest) (*ResponseBody, error)
}

type IRepository interface {
	GetAll() ([]*Product, error)
	GetById(productId int) (*Product, error)
	Create(body *Product) error
	GetByCodeValue(codeValue string) (*Product, error)
	DeleteById(productId int) error
	UpdateById(productId int, body *Product) error
}
