package handlers

import (
	"github.com/emanuel3k/product-golang/internal/domain/product"
	"github.com/emanuel3k/product-golang/pkg/web/response"
	"net/http"
)

type productHandler struct {
	productService product.IService
}

func (ph productHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	res, err := ph.productService.GetAll()

	if err != nil {
		// todo
	}

	if len(res) == 0 {
		response.JSON(w, http.StatusNoContent, nil)
		return
	}

	response.JSON(w, http.StatusOK, res)
}

func NewProductHandler(productService product.IService) productHandler {
	return productHandler{productService}
}
