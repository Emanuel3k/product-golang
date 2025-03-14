package product

type productService struct {
	productRepository IRepository
}

func (ps productService) GetAll() ([]*Product, error) {
	res, err := ps.productRepository.GetAll()

	if err != nil {
		// todo
	}

	return res, nil
}

func NewService(productRepository IRepository) IService {
	return productService{productRepository}
}
