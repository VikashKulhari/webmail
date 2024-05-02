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

	// hUserName := string(helpers.HashSHA256(userSignCreds.Password))
	// hEmail := string(helpers.HashSHA256(userSignCreds.Email))
	// if err := m.DB.Where("Username = ?", hUserName).First(&existingUser).Error; err == nil {
	// 	return true, existingUser
	// }
	if err := m.DB.Where("Email = ?", userSignCreds.Email).First(&existingUser).Error; err != nil {
		fmt.Println("err: ", err)
		return false, existingUser
	}
	return true, existingUser
}
func (m *model) IsExistingUserByEMailID(emailID string) bool {
	var existingUser entities.User

	// hUserName := string(helpers.HashSHA256(userSignCreds.Password))
	// hEmail := string(helpers.HashSHA256(emailID))
	// if err := m.DB.Where("Username = ?", hUserName).First(&existingUser).Error; err == nil {
	// 	return true, existingUser
	// }

	if err := m.DB.Where("Email = ?", emailID).First(&existingUser).Error; err == nil {
		logger.Print(err)
		// fmt.Println("IsExistingUser: ", err.Error())
		return true

	}
	return false
}
func (m *model) CreateUser(userSignCreds entities.User) error {
	fmt.Println("email: ", userSignCreds.Email)
	userSignCreds.CreatedAt = time.Now()
	fmt.Println("password: ", userSignCreds)
	if err := m.DB.Create(&userSignCreds).Error; err != nil {
		// zap.S().Warnw("Error while creating new user: sigUp", "err", err)
		return err
	}
	return nil
}