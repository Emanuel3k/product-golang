package request

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrRequestContentTypeNotJSON = errors.New("request content type is not application/json")
	InvalidRequestBody           = errors.New("invalid request body")
)

func JSON(r *http.Request, ptr any) (err error) {

	if r.Header.Get("Content-Type") != "application/json" {
		err = ErrRequestContentTypeNotJSON
		return
	}

	err = json.NewDecoder(r.Body).Decode(ptr)
	if err != nil {
		err = InvalidRequestBody
		return
	}

	return
}
