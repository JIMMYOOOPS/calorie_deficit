package types

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (error *AppError) Error() string {
	return error.Message
}

// NewAppError creates a new AppError instance
func NewAppError(code int, message string, details error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
