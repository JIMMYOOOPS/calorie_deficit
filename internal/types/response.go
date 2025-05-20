package types

// Response represents a standard API response structure
type SuccessResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"` // Pointer to Meta struct
}

type Meta struct {
	Page       uint `json:"page"`
	PageSize   uint `json:"page_size"`
	TotalCount uint `json:"total_count"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
