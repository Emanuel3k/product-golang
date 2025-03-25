package appError

import "net/http"

type AppError struct {
	statusCode int
	message    string
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) StatusCode() int {
	return e.statusCode
}

func Conflict(message string) error {
	return &AppError{
		statusCode: http.StatusConflict,
		message:    message,
	}
}

func NotFound(message string) error {
	return &AppError{
		statusCode: http.StatusNotFound,
		message:    message,
	}
}
