package web

import (
	"net/http"

	"github.com/VikashKulhari/P1/handlers"
	ijwt "github.com/VikashKulhari/P1/jwt"
	"github.com/go-chi/chi/v5"
)

// func Routers(r chi.Router, h v1.HandlerV1, jwtString string) http.Handler {
// 	// r := chi.NewRouter()
// 	// r.Post("/signup", routers.RouterB(*h))
// 	logGroup := r.Group(nil)
// 	logGroup1 := r.Group(nil)
// 	logGroup1.Group(func(r chi.Router) {
// 		r.Route("/v2", func(r chi.Router) {
// 			r.Mount("/", routers.RouterB(*h))
// 		})

//		})
//		logGroup.Group(func(r chi.Router) {
//			r.Route("/v1", func(r chi.Router) {
//				r.Mount("/", routers.RoutersA(*h, jwtString))
//			})
//		})
//		return r
//	}
//
//	func Routers(h v1.HandlerV1, jwtString string) http.Handler {
//		r := chi.NewRouter()
//		r.With(ijwt.UserAuth(jwtString)).Route("/", func(r chi.Router) {
//			r.Post()
//		})
//	}
//
//	func RouteA(h *handlers.Handler) http.Handler {
//		router := chi.NewRouter()
//		router.Post("/signin", h.V1.SignIn)
//		router.Post("/signup", h.V1.SignUp)
//		return router
//	}
func Routers(h *handlers.Handler) http.Handler {
	router := chi.NewRouter()
	// router.Use(ijwt.Auth)
	router.Post("/signin", h.V1.SignIn)
	router.Post("/signup", h.V1.SignUp)
	router.With(ijwt.Auth).Post("/sendMail", h.V1.SendMail)
	// router.Post("/sendMail", h.V1.SendMail)
	// router.Post("/sendMail", handlers.Handler.SendMail)
	return router
}

// /func Routers(h *handlers.Handler) http.HandlerFunc
