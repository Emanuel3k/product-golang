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

func (ps *productService) Create(request BodyRequest) (*Product, error) {
	exists, err := ps.productRepository.GetByCodeValue(request.CodeValue)
	if err != nil {
		// todo
		return nil, nil
	}

	if exists != nil {
		// todo
		return nil, nil
	}

	product := request.toDomain()

	if err := ps.productRepository.Create(&product); err != nil {
		// todo
	}

	return &product, nil
}

func NewService(productRepository IRepository) IService {
	return &productService{productRepository}
}
