package models

import (
	"github.com/VikashKulhari/P1/entities"
)

func (m *model) DeleteMail(mailID uint, emailid string) error {

	if err := m.DB.Where(`"to" = ? AND "deleted_at" IS NULL AND "id"=?`, emailid, mailID).Delete(&entities.Email{}).Error; err != nil {
		return err
	}
	return nil
}
