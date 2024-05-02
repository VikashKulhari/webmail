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
	// r := chi.NewRouter()
	// r.Use(cors.Default().Handler)
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Group(web.RouteA())
	// r.Post("/signup", handlers.Handler.V1.)
	// r.Post("/signin", handlers.Handler.V1.SignIn)
	db := db.Init()
	model := models.New(db)
	service := services.New(&model)
	handlers := handlers.New(service)
	r := web.Routers(handlers)
	// r.Group(web.Route)
	// r.Group(web.Route(handlers1))
	// r.Post("/signup", h.)
	// model := models.New(db.Init())
	// service := services.New(&model)
	// handlers1 := handlers.New(service)
	// rs := web.Routers(r, &handlers1, "")
	// go func() {
	// routerX := mux.NewRouter()
	// 	routerX.HandlerFunc("/signup",handlers.V1.SignUp())
	port := ":8080"
	log.Printf("Server started on 8080")
	log.Fatal(http.ListenAndServe(port, r))
	// }()
	// r.Post("/signUp", handlers.Handler.SignUp)

}
