// Package handlers
package handlers

import "github.com/PrinceNarteh/pos/internal/services"

type Handlers struct {
	Auth AuthHandler
	User UserHandler
}

func NewHandlers(svc *services.Services) *Handlers {
	return &Handlers{
		Auth: &authHandler{svc: svc},
		User: &userHandler{svc: svc},
	}
}

type Response struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    any    `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
}

func SuccessResponse(statusCode int, data interface{}) Response {
	return Response{
		Status:     "success",
		StatusCode: statusCode,
		Data:       data,
	}
}

func ErrResponse(statusCode int, message string) Response {
	return Response{
		Status:     "error",
		StatusCode: statusCode,
		Message:    message,
	}
}
