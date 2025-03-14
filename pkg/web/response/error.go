package response

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	defaultStatusCode := http.StatusInternalServerError

	if statusCode > 299 && statusCode < 600 {
		defaultStatusCode = statusCode
	}

	body := errorResponse{
		http.StatusText(defaultStatusCode),
		message,
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(defaultStatusCode)
	_, _ = w.Write(bytes)
}
