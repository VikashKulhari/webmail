package models

import (
	"github.com/VikashKulhari/P1/entities"
	"gorm.io/gorm"
)

type model struct {
	DB *gorm.DB
}
type Model interface {
	IsExistingUser(userSignCreds entities.User) (bool, entities.User)
	CreateUser(userSignCreds entities.User) error
	IsExistingUserByEMailID(emailID string) bool
	SendEmail(emailReq entities.Email) error
}

// SignUp implements services.Service.

func New(db *gorm.DB) Model {
	return &model{DB: db}
}
