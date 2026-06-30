package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)


type AuthMiddleware struct {}

func(a *AuthMiddleware) VerifyUser(request *http.Request) {
	token := request.Header.Get("Authorization")

	
}