package web

import (
	"net/http"

	"github.com/VikashKulhari/P1/handlers"
	ijwt "github.com/VikashKulhari/P1/jwt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func Routers(h *handlers.Handler) http.Handler {
	router := chi.NewRouter()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsHandler.Handler)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	//without authentication
	router.Post("/signin", h.V1.SignIn)
	router.Post("/signup", h.V1.SignUp)

	//with authentication - later on we will make a new service
	router.With(ijwt.Auth).Post("/sendMail", h.V1.SendMail)
	router.With(ijwt.Auth).Get("/getRecievedEmails", h.V1.GetRecievedMails)
	router.With(ijwt.Auth).Get("/getSentEmails", h.V1.GetSentMails)
	router.With(ijwt.Auth).Delete("/deleteMail/{ID}", h.V1.DeleteMail)
	router.With(ijwt.Auth).Patch("/markImp/{ID}", h.V1.MarkImp)
	router.With(ijwt.Auth).Patch("/markSpam/{ID}", h.V1.MarkSpam)
	return router
}
