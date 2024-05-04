package ijwt

import (
	"github.com/VikashKulhari/P1/entities"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(req entities.User) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":    req.Email,
		"Password": req.Password,
	})
	return token
}
