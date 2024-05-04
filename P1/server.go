package main

import (
	"log"
	"net/http"

	"github.com/VikashKulhari/P1/db"
	"github.com/VikashKulhari/P1/handlers"
	"github.com/VikashKulhari/P1/models"
	"github.com/VikashKulhari/P1/services"
	"github.com/VikashKulhari/P1/web"
)

func main() {

	db := db.Init()
	model := models.New(db)
	service := services.New(&model)
	handlers := handlers.New(service)
	r := web.Routers(handlers)
	port := ":8080"
	log.Printf("Server started on 8080")
	log.Fatal(http.ListenAndServe(port, r))
}
