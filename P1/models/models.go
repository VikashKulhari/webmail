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
	GetEmailsRecievedByEmailId(emailid string) ([]entities.Email, error)
	GetEmailsSentByEmailId(emailid string) ([]entities.Email, error)
	DeleteMail(mailID uint, emailid string) error
	MarkImportant(mailID uint, emailid string) error
	MarkSpam(mailID uint, emailid string) error
}

func New(db *gorm.DB) Model {
	return &model{DB: db}
}
