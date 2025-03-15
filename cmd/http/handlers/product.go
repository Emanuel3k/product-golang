package handlers

import (
	"encoding/json"
	"github.com/emanuel3k/product-golang/internal/domain/product"
	"github.com/emanuel3k/product-golang/pkg/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
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
		// todo
	}

	if res == nil {
		response.JSON(w, http.StatusNoContent, nil)
	}

	response.JSON(w, http.StatusOK, res)
}

func (ph *productHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body product.BodyRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		// todo
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := ph.productService.Create(body)
	if err != nil {
		// todo
	}

	response.JSON(w, http.StatusCreated, res)
}

func (ph *productHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "productId")

	productId, err := strconv.Atoi(id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid product id")
		return
	}

	if err := ph.productService.DeleteById(productId); err != nil {
		// todo
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func NewProductHandler(productService product.IService) productHandler {
	return productHandler{productService}
}
