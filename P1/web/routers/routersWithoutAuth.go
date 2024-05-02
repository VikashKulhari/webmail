package routers

// import (
// 	"net/http"

// 	"github.com/VikashKulhari/P1/handlers"
// 	ijwt "github.com/VikashKulhari/P1/jwt"
// 	"github.com/go-chi/chi/v5"
// )

// func Routers(h handlers.Handler, jwtString string) http.Handler {
// 	r := chi.NewRouter()
// 	r.With(ijwt.UserAuth(jwtString)).Route("/",func(r chi.Router){
// 		r.Post("/signIn",h.SignIn)

// 	} )
// 	return r

// }
