package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null type:varchar(100);unique_index"`
	Password string `json:"password" gorm:"not null type:varchar(100)"`
}
