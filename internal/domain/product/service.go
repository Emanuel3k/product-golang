package product

type productService struct {
	productRepository IRepository
}

func (ps *productService) GetAll() ([]*Product, error) {
	res, err := ps.productRepository.GetAll()

	if err != nil {
		// todo
	}

	return res, nil
}

func (ps *productService) GetById(productId int) (*Product, error) {
	res, err := ps.productRepository.GetById(productId)
	if err != nil {
		// todo
	}

	return res, nil
}

func (ps *productService) Create(body CreateBodyRequest) (*Product, error) {
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

func (ps *productService) UpdateById(productId int, body UpdateBodyRequest) (*Product, error) {

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

func NewService(productRepository IRepository) IService {
	return &productService{productRepository}
}
