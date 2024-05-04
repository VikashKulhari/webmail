package services

import (
	"github.com/VikashKulhari/P1/entities"
)

func (s *service) GetEmailsSentById(emailid string) ([]entities.Email, error) {

	emails, err := s.model.GetEmailsSentByEmailId(emailid)
	if err != nil {
		return emails, err
	}
	return emails, nil

}
