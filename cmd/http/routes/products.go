package routes

import (
	"github.com/emanuel3k/product-golang/cmd/http/handlers"
	"github.com/emanuel3k/product-golang/internal/domain/product"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func buildProductsRoutes() http.Handler {
	r := chi.NewRouter()

	productRepository := product.NewRepository()
	productService := product.NewService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	r.Get("/", productHandler.GetAll)

	return r
}
