package handlers

import (
	"encoding/json"
	"errors"
	"github.com/emanuel3k/product-golang/internal/domain"
	"github.com/emanuel3k/product-golang/pkg/appError"
	"github.com/emanuel3k/product-golang/pkg/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"net/http"
	"strconv"
)

var (
	InvalidRequestId   = errors.New("invalid product id")
	InvalidRequestBody = errors.New("invalid request body")
)

type productHandler struct {
	productService domain.IService
}

func (ph *productHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	res, err := ph.productService.GetAll()

	if err != nil {
		var errResponse *appError.AppError
		if errors.As(err, &errResponse) {
			response.Error(w, errResponse.StatusCode(), errResponse.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, res)
}

func (ph *productHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "productId")

	productId, err := strconv.Atoi(id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, InvalidRequestId.Error())
		return
	}

	res, err := ph.productService.GetById(productId)
	if err != nil {
		var errResponse *appError.AppError
		if errors.As(err, &errResponse) {
			response.Error(w, errResponse.StatusCode(), errResponse.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if res == nil {
		response.JSON(w, http.StatusNoContent, nil)
	}

	response.JSON(w, http.StatusOK, res)
}

func (ph *productHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body domain.CreateBodyRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, InvalidRequestBody.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := ph.productService.Create(body)
	if err != nil {
		var errResponse *appError.AppError
		if errors.As(err, &errResponse) {
			response.Error(w, errResponse.StatusCode(), errResponse.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, res)
}

func (ph *productHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "productId")

	productId, err := strconv.Atoi(id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, InvalidRequestId.Error())
		return
	}

	if err := ph.productService.DeleteById(productId); err != nil {
		var errResponse *appError.AppError
		if errors.As(err, &errResponse) {
			response.Error(w, errResponse.StatusCode(), errResponse.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func (ph *productHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "productId")

	productId, err := strconv.Atoi(id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, InvalidRequestId.Error())
		return
	}

	var body domain.UpdateBodyRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, InvalidRequestBody.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := ph.productService.UpdateById(productId, body)
	if err != nil {
		var errResponse *appError.AppError
		if errors.As(err, &errResponse) {
			response.Error(w, errResponse.StatusCode(), errResponse.Error())
			return
		}

		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, res)
}

func NewProductHandler(productService domain.IService) productHandler {
	return productHandler{productService}
}
