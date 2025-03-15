package product

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"codeValue"`
	IsPublished bool    `json:"isPublished"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type BodyRequest struct {
	Name        string  `json:"name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CodeValue   string  `json:"codeValue" validate:"required"`
	IsPublished bool    `json:"isPublished"`
	Expiration  string  `json:"expiration" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

func (br *BodyRequest) toDomain() Product {
	return Product{
		Name:        br.Name,
		Quantity:    br.Quantity,
		CodeValue:   br.CodeValue,
		IsPublished: br.IsPublished,
		Expiration:  br.Expiration,
		Price:       br.Price,
	}
}

type IService interface {
	GetAll() ([]*Product, error)
	GetById(productId int) (*Product, error)
	Create(body BodyRequest) (*Product, error)
	DeleteById(productId int) error
}

type IRepository interface {
	GetAll() ([]*Product, error)
	GetById(productId int) (*Product, error)
	Create(p *Product) error
	GetByCodeValue(codeValue string) (*Product, error)
	DeleteById(productId int) error
}
