package v1

import (
	"net/http"

	"github.com/VikashKulhari/P1/services"
)

type handlerV1 struct {
	Service services.Service
}
type HandlerV1 interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	SendMail(w http.ResponseWriter, r *http.Request)
}

func New(s services.Service) HandlerV1 {
	return &handlerV1{Service: s}
}
