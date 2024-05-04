package ijwt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const (
	ContextEmail    contextKey = "UserIDClaim"
	ContextPassword contextKey = "UserPasswordClaim"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tt := r.Header.Get("Authorization")
		if len(tt) == 0 || tt == "" {
			http.Error(w, "Authorization Token Missing", http.StatusUnauthorized)
			return
		}
		tt = strings.TrimPrefix(tt, "Bearer ")
		tt = strings.TrimPrefix(tt, " ")

		token, err := jwt.Parse(tt, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})
		if err != nil {
			http.Error(w, "error while parsing the authentication token", http.StatusUnauthorized)
		}
		var UserIDClaim string = ""
		var UserPasswordClaim string = ""
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			UserIDClaim = claims["Email"].(string)
			UserPasswordClaim = claims["Password"].(string)
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextEmail, UserIDClaim)
		ctx = context.WithValue(ctx, ContextPassword, UserPasswordClaim)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
func UserAuth(jwtString string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			tt := r.Header.Get("Authorization")
			if len(tt) == 0 {
				res, _ := json.Marshal(map[string]string{
					"message": "Token Not Found",
				})
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Write(res)
				return
			}
			tt1, err := jwt.Parse(tt, func(token *jwt.Token) (interface{}, error) {
				return []byte("key"), nil
			})
			if err != nil {
				res, _ := json.Marshal(map[string]string{
					"message": "Token Invalid",
				})
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Write(res)
				return
			}
			var userIDClaim uint64
			var UserPasswordClaim string
			if claims, ok := tt1.Claims.(jwt.MapClaims); ok && tt1.Valid {
				userIDClaim = uint64(claims["UserID"].(float64))
				UserPasswordClaim = claims["UserName"].(string)
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, ContextEmail, userIDClaim)
			ctx = context.WithValue(ctx, ContextPassword, UserPasswordClaim)
			r = r.WithContext(ctx)
			next.ServeHTTP(rw, r)
		})
	}
}
