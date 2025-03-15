package product

func NewRepository() IRepository {
	return &productJsonRepository{}
}
