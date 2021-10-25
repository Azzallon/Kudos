package errs

import (
	"net/http"
)

type ApiError struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message"`
}

func (a ApiError) Error() string {
	return a.Message
}

func (a ApiError) AsMessage() *ApiError {
	return &ApiError{
		Message: a.Error(),
	}
}

func NewBadRequestError(message string) *ApiError {
	return &ApiError{http.StatusBadRequest, message}
}

func NewNotFoundError(message string) *ApiError {
	return &ApiError{http.StatusNotFound, message}
}

func NewUnauthorizedError(message string) *ApiError {
	return &ApiError{http.StatusUnauthorized, message}
}

func NewForbiddenError(message string) *ApiError {
	return &ApiError{http.StatusForbidden, message}
}

func NewUnexpectedError(message string) *ApiError {
	return &ApiError{http.StatusInternalServerError, message}
}

func NewValidationError(message string) *ApiError {
	return &ApiError{http.StatusUnprocessableEntity, message}
}
