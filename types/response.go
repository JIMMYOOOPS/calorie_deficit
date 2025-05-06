package types

// Response represents a standard API response structure
type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   Error  `json:"error,omitempty"`
}

// Error represents an error response structure
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
