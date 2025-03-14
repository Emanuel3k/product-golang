package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type router struct {
}

func (router *router) MapRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/products", buildProductsRoutes())
	})

	return r
}

func NewRouter() *router {
	return &router{}
}
