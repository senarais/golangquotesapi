package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func(r *RegisterService) Register(request *http.Request) (string, error) {
	userData := r.UserDatabase.GetData()
	user := model.User{}

	err:= json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		return "", fmt.Errorf("failed to decode %w", err)
	}
	fmt.Println(user)

	// Masukin user ke database.
	for _, u := range userData {
		if u.Username == user.Username {
			return "", errors.New("Username already existed")
		}
	}
	r.UserDatabase.AddUser(user)

	// Login User
	login := LoginService{}
	token, err :=  login.Login(request)
	if err != nil {
		return "",errors.New("Login gagal.")
	}

	return token, nil
}

func(l *LoginService) Login(request *http.Request) (string, error) {
	userData := l.UserDatabase.GetData()
	user := model.User{}

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		return "",fmt.Errorf("failed to decode %w", err)
	}

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

	exp := time.Now().Add(24 * time.Hour).Unix()

	key := []byte("jwtsecret")
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, 
	jwt.MapClaims{ 
		"username": user.Username,
		"exp": exp,
	})

	tokenString, err := token.SignedString(key)
	    if err != nil {
        return "", err
    }

	return tokenString, nil
}