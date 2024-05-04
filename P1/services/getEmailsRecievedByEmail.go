package services

import (
	"github.com/VikashKulhari/P1/entities"
)

func (s *service) GetEmailsRecievedById(emailid string) ([]entities.Email, error) {

	emails, err := s.model.GetEmailsRecievedByEmailId(emailid)
	if err != nil {
		return emails, err
	}
	return emails, nil

}
