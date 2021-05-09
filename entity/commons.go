package entity

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Status string `json:"status"`
}

func NewError(s string) ErrorResponse {
	return ErrorResponse{Error: s}
}
