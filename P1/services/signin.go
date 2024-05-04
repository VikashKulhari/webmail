package services

import (
	"errors"
	"fmt"

	"github.com/VikashKulhari/P1/entities"
	ijwt "github.com/VikashKulhari/P1/jwt"
)

func (s *service) SignIn(signinCreds entities.User) (string, error) {
	IsExisting, existingUser := s.model.IsExistingUser(signinCreds)
	if !IsExisting {
		fmt.Println(existingUser)
		fmt.Println(IsExisting)
		return "", errors.New("userNotFound")
	}

	//check password
	if string(signinCreds.Password) != existingUser.Password {
		return "", errors.New("incorrectPassword")
	}
	token := ijwt.GenerateToken(signinCreds)

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		fmt.Println("error after converting token to string: ", err)
		return "", errors.New("issue in tokenization")
	}
	return tokenString, nil
}
