package routes

import (
	"github.com/emanuel3k/product-golang/cmd/http/handlers"
	"github.com/emanuel3k/product-golang/internal/repositories"
	"github.com/emanuel3k/product-golang/internal/services"
	"github.com/emanuel3k/product-golang/storage/postgres"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func buildProductsRoutes() http.Handler {
	r := chi.NewRouter()

	psql, _ := postgres.Config()
	productRepository := repositories.NewRepository(psql)
	productService := services.NewService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	r.Get("/", productHandler.GetAll)
	r.Get("/{productId}", productHandler.GetById)
	r.Post("/", productHandler.Create)
	r.Delete("/{productId}", productHandler.DeleteById)
	r.Put("/{productId}", productHandler.UpdateById)

	return r
}
