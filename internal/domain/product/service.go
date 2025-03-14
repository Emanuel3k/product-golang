package product

type productService struct {
	productRepository IRepository
}

func newService(productRepository IRepository) IService {
	return productService{productRepository}
}
