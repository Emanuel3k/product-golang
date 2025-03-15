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
	r.Get("/{productId}", productHandler.GetById)
	r.Post("/", productHandler.Create)
	r.Delete("/{productId}", productHandler.DeleteById)
	r.Put("/{productId}", productHandler.UpdateById)

	return r
}
