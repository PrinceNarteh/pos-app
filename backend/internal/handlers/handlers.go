// Package handlers
package handlers

import "github.com/PrinceNarteh/pos/internal/services"

type Handlers struct {
	Auth     AuthHandler
	User     UserHandler
	Category CategoryHandler
}

func NewHandlers(svc *services.Services) *Handlers {
	return &Handlers{
		Auth:     &authHandler{svc: svc},
		User:     &userHandler{svc: svc},
		Category: &categoryHandler{svc: svc},
	}
}

type Response struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    any    `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
}

func SuccessResponse(statusCode int, data any) Response {
	return Response{
		Status:     "success",
		StatusCode: statusCode,
		Data:       data,
	}
}

func ErrResponse(statusCode int, err any) Response {
	return Response{
		Status:     "error",
		StatusCode: statusCode,
		Message:    err,
	}
}
