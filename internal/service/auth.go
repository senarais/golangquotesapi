package service

import (
	"errors"
	"time"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/senarais/golangquotesapi/internal/database"
	"github.com/senarais/golangquotesapi/internal/model"
)

type RegisterService struct {
	UserDatabase *database.UserData
}

type LoginService struct {
	UserDatabase *database.UserData
}

func(r *RegisterService) Register(user model.User) (string, error) {
	userData := r.UserDatabase.GetData()

	// Masukin user ke database.
	for _, u := range userData {
		if u.Username == user.Username {
			return "", errors.New("Username already existed")
		}
	}
	r.UserDatabase.AddUser(user)

	// Login User
	login := LoginService{}
	token, err :=  login.Login(user)
	if err != nil {
		return "",errors.New("Login gagal.")
	}

	return token, nil
}

func(l *LoginService) Login(user model.User) (string, error) {
	userData := l.UserDatabase.GetData()

	for _, u := range userData {
		if u.Username == user.Username && u.Password == user.Password {
			token, err := l.createJwtToken(user)
			if err != nil {
				return "", err
			}
			return token, nil
		}
	}

	return "", errors.New("User Not Found.")
}

func(l *LoginService) createJwtToken(user model.User) (string, error) {
	claims := model.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	
	key := []byte("jwtsecret")
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, 
	claims)

	tokenString, err := token.SignedString(key)
	    if err != nil {
        return "", err
    }

	return tokenString, nil
}