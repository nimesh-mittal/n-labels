package entity

// ErrorResponse represents common error object
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse represents common success object
type SuccessResponse struct {
	Status string `json:"status"`
}

// NewError creates new object of ErrorResponse
func NewError(s string) ErrorResponse {
	return ErrorResponse{Error: s}
}
