package models

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/VikashKulhari/P1/entities"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func (m *model) IsExistingUser(userSignCreds entities.User) (bool, entities.User) {
	var existingUser entities.User
	if err := m.DB.Where("Email = ?", userSignCreds.Email).First(&existingUser).Error; err != nil {
		fmt.Println("err: ", err)
		return false, existingUser
	}
	return true, existingUser
}

func (m *model) IsExistingUserByEMailID(emailID string) bool {
	var existingUser entities.User

	if err := m.DB.Where("Email = ?", emailID).First(&existingUser).Error; err == nil {
		logger.Print(err)
		return true

	}
	return false
}

func (m *model) CreateUser(userSignCreds entities.User) error {
	fmt.Println("email: ", userSignCreds.Email)
	userSignCreds.CreatedAt = time.Now()
	fmt.Println("password: ", userSignCreds)
	if err := m.DB.Create(&userSignCreds).Error; err != nil {
		return err
	}
	return nil
}
