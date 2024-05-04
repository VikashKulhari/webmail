package handlers

import (
	v1 "github.com/VikashKulhari/P1/handlers/v1"
	"github.com/VikashKulhari/P1/services"
)

type Handler struct {
	V1 v1.HandlerV1
}

func New(service services.Service) *Handler {
	return &Handler{
		V1: v1.New(service),
	}
}
