package routers

import (
	"net/http"

	"github.com/VikashKulhari/P1/handlers"
	ijwt "github.com/VikashKulhari/P1/jwt"
	"github.com/go-chi/chi"
)

func RoutersA(h handlers.Handler, jwtString string) http.Handler {
	r := chi.NewRouter()

	r.With(ijwt.UserAuth(jwtString)).Route("/mail", func(r chi.Router) {
		r.Post("/sendEmail", h.V1.SendMail)
	})
	return r
}

func RouterB(h handlers.Handler) http.Handler {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/signup", h.V1.SignUp)
		r.Post("/signin", h.V1.SignIn)
	})
	return r
}
