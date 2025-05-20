// internal/constants/errors.go
package constants

import (
	"net/http"
)

// Form the error code and message according to the error type
// AppError defines a standard application error structure.
type AppError struct {
	Code    int
	Message string
}

// Form the error code and message according to the error type
var (
	// ErrInternalServerError is an error for internal server error
	ErrInternalServerError = AppError{
		Code:    http.StatusInternalServerError,
		Message: LogMessages.General.InternalServerError,
	}
	// ErrInvalidRequest is an error for invalid request
	ErrInvalidRequest = AppError{
		Code:    http.StatusBadRequest,
		Message: LogMessages.General.InvalidRequest,
	}
	// ErrNotFound is an error for not found
	ErrNotFound = AppError{
		Code:    http.StatusNotFound,
		Message: LogMessages.General.NotFound,
	}
)
