package db

import (
	"fmt"

	"github.com/VikashKulhari/P1/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	var db *gorm.DB
	var err error
	destination := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable search_path= public"
	fmt.Println("Connecting to database...Please wait")
	db, err = gorm.Open(postgres.Open(destination), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Email{})
	fmt.Println("Connected!")

	return db
}
