package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
)

type Auth struct {
	secretKey string
}

func New(secretKey string) (Auth, error) {
	auth := Auth{
		secretKey: secretKey,
	}
	return auth, nil
}

func (a Auth) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var secretKey = []byte(a.secretKey)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a Auth) ValidateToken(token string) (jwt.StandardClaims, error) {
	var claims jwt.StandardClaims
	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.secretKey), nil
	})
	if err != nil {
		return jwt.StandardClaims{}, fmt.Errorf("parsing token: %w", err)
	}

	if !t.Valid {
		return jwt.StandardClaims{}, errors.New("invalid token")
	}
	return claims, nil

}
