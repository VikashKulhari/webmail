package models

import (
	"github.com/VikashKulhari/P1/entities"
)

func (m *model) MarkImportant(mailID uint, emailid string) error {

	if err := m.DB.Model(&entities.Email{}).
		Where(`"to" = ? AND "deleted_at" IS NULL AND "id"=?`, emailid, mailID).
		Update("is_important", true).Error; err != nil {
		return err
	}
	return nil
}