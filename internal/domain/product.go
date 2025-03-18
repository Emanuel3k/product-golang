package domain

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"codeValue"`
	IsPublished bool    `json:"isPublished"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type CreateBodyRequest struct {
	Name        string  `json:"name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CodeValue   string  `json:"codeValue" validate:"required"`
	IsPublished bool    `json:"isPublished"`
	Expiration  string  `json:"expiration" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type UpdateBodyRequest struct {
	Name        *string  `json:"name"`
	Quantity    *int     `json:"quantity"`
	CodeValue   *string  `json:"codeValue"`
	IsPublished *bool    `json:"isPublished"`
	Expiration  *string  `json:"expiration"`
	Price       *float64 `json:"price"`
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

type IService interface {
	GetAll() ([]*Product, error)
	GetById(productId int) (*Product, error)
	Create(body CreateBodyRequest) (*Product, error)
	DeleteById(productId int) error
	UpdateById(productId int, body UpdateBodyRequest) (*Product, error)
}

type IRepository interface {
	GetAll() ([]*Product, error)
	GetById(productId int) (*Product, error)
	Create(p *Product) error
	GetByCodeValue(codeValue string) (*Product, error)
	DeleteById(productId int) error
	UpdateById(productId int, body UpdateBodyRequest) (*Product, error)
}
