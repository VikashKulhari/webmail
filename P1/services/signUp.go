package services

import (
	"github.com/VikashKulhari/P1/entities"
)

func (s *service) SignUp(signupCreds entities.User) error {

	// userAlreadyExists, _ := s.model.IsExistingUser(signupCreds)
	// if userAlreadyExists {
	// 	return errors.New("exists")
	// }
	if err := s.model.CreateUser(signupCreds); err != nil {
		return err
	}
	return nil
}
