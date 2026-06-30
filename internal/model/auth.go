package model


type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type TokenJWT struct {
	Token string
}