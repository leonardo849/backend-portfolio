package dto

type MessageResponseDTO struct {
	Message string `json:"message"`
}

type ErrorResponseDTO struct {
	Error string `json:"error"`
}