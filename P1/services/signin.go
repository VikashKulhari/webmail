package services

import (
	"errors"
	"fmt"

	"github.com/VikashKulhari/P1/entities"
	ijwt "github.com/VikashKulhari/P1/jwt"
)

func (s *service) SignIn(signinCreds entities.User) (string, error) {
	// signinCreds1 := entities.User{
	// 	Email:    signinCreds.Email,
	// 	Password: signinCreds.Password,
	// }
	IsExisting, existingUser := s.model.IsExistingUser(signinCreds)
	if !IsExisting {
		fmt.Println(existingUser)
		fmt.Println(IsExisting)
		return "", errors.New("userNotFound")
	}
	if string(signinCreds.Password) != existingUser.Password {
		return "", errors.New("incorrectPassword")
	}
	token := ijwt.GenerateToken(signinCreds)
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	fmt.Println("tokenString: ", tokenString)
	if err != nil {
		fmt.Println("error after converting token to string: ", err)
		return "", errors.New("issue in tokenization")
	}
	return tokenString, nil
}
