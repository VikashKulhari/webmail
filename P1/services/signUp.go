package services

import (
	"fmt"

	"github.com/VikashKulhari/P1/entities"
)

func (s *service) SignUp(signupCreds entities.User) error {
	if err := s.model.CreateUser(&signupCreds); err != nil {
		return err
	}
	fmt.Println(signupCreds)
	return nil
}
