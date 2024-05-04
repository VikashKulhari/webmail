package models

import (
	"errors"
	"fmt"

	"github.com/VikashKulhari/P1/entities"
)

func (m *model) GetEmailsSentByEmailId(emailid string) ([]entities.Email, error) {
	var emails []entities.Email
	fmt.Println("emails is : ", emailid)
	if err := m.DB.Find(&entities.Email{}, `"from" = ?  AND "is_draft" = ?`, emailid, false).Find(&emails).Error; err != nil {
		return emails, errors.New("error fetching emails")
	}
	return emails, nil
}
