package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (err AppError) AsMsg() *AppError {
	return &AppError{Message: err.Message}
}

func NewNotFoundError(msg string) *AppError {
	return &AppError{Message: msg, Code: http.StatusNotFound}
}

func NewUnexpectedError(msg string) *AppError {
	return &AppError{Message: msg, Code: http.StatusInternalServerError}
}

func ValidationError(msg string) *AppError {
	return &AppError{Message: msg, Code: http.StatusUnprocessableEntity}
}

func NewUnavailable(msg string) *AppError {
	return &AppError{Message: msg, Code: http.StatusServiceUnavailable}
}
