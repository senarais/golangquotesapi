package model

import "github.com/golang-jwt/jwt/v5"


type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenJWT struct {
	Token string
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}