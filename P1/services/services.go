package services

import (
	"github.com/VikashKulhari/P1/entities"
	"github.com/VikashKulhari/P1/models"
)

type service struct {
	model models.Model
}

// SignUp implements Service.

type Service interface {
	SignUp(entities.User) error
	SignIn(signinCreds entities.User) (string, error)
	SendEmail(emailReq entities.EmailReq) error 
}

func New(model *models.Model) Service {
	return &service{model: *model}
}
