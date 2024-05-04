package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"not null type:varchar(100);unique_index"`
	Password string `gorm:"not null type:varchar(100)"`
}
