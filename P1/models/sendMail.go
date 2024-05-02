package models

import (
	"fmt"

	"github.com/VikashKulhari/P1/entities"
	"go.uber.org/zap"
)

func (m *model) SendEmail(emailReq entities.Email) error {
	if err := m.DB.Create(&emailReq).Error; err != nil {
		zap.S().Warnw("Error Saving Mail", "err", err)
		fmt.Println("Error saving mail...modelLayer")
		return err
	}
	return nil
}
