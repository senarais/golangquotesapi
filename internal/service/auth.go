package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/senarais/golangquotesapi/internal/database"
	"github.com/senarais/golangquotesapi/internal/model"
)

type RegisterService struct {
	UserDatabase *database.UserData
}

func(r *RegisterService) Register(request *http.Request) error {
	userData := r.UserDatabase.GetData()
	user := model.User{}

	err:= json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		return fmt.Errorf("failed to decode %w", err)
	}
	fmt.Println(user)

	// Masukin user ke database.
	for _, u := range userData {
		if u.Username == user.Username {
			return errors.New("Username already existed")
		}
	}
	r.UserDatabase.AddUser(user)

	return nil
}