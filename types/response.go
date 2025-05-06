package types

// Response represents a standard API response structure
type SuccessResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"` // Pointer to Meta struct
}

type Meta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
