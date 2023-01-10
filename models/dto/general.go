package dto

type MessageResponse struct {
	Message string `json:"message" binding:"required"`
}

type ErrorResponse struct {
	Error string `json:"error" binding:"required"`
}
