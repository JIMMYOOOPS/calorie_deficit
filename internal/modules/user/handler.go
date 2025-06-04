package user

import (
	"calorie_deficit/internal/constants"
	"calorie_deficit/internal/types"
	//"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

var (
	// ErrInvalidRequest is an error for invalid request
	errInvalidRequest = types.ErrorResponse{
		Code:    400,
		Message: constants.LogMessages.General.InvalidRequest,
	}
)

// TODO: Implement HTTP handlers for user
// CreateUser handles the creation of a new User record

// GetUser handles the retrieval of a User record by ID

// GetUserList handles the retrieval of all User records

// UpdateUser handles the update of an existing User record
