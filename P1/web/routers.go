package web

import (
	"net/http"

	"github.com/VikashKulhari/P1/handlers"
	ijwt "github.com/VikashKulhari/P1/jwt"
	"github.com/go-chi/chi/v5"
)

func Routers(h *handlers.Handler) http.Handler {
	router := chi.NewRouter()
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
