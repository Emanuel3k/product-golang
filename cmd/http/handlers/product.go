package handlers

import (
	"github.com/emanuel3k/product-golang/internal/domain/product"
	"github.com/emanuel3k/product-golang/pkg/web/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type productHandler struct {
	productService product.IService
}

func (ph *productHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	res, err := ph.productService.GetAll()

	if err != nil {
		// todo
	}

	response.JSON(w, http.StatusOK, res)
}

func (ph *productHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "productId")

	productId, err := strconv.Atoi(id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid product id")
		return
	}

	res, err := ph.productService.GetById(productId)
	if err != nil {

	}

	if res == nil {
		response.JSON(w, http.StatusNoContent, nil)
	}

	response.JSON(w, http.StatusNoContent, res)
}

func NewProductHandler(productService product.IService) productHandler {
	return productHandler{productService}
}
