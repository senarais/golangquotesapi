package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/senarais/golangquotesapi/internal/model"
)


func Verify() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
        if err != nil {
            c.Redirect(http.StatusTemporaryRedirect, "/login")
            return
        }
	
		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (any, error) {
			return []byte("jwtsecret"), nil
		})
	
		if err != nil || !token.Valid {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return
		} else if claims, ok := token.Claims.(*model.Claims); ok {
			c.Set("username", claims.Username)
			c.Next()
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return 
		}
	}
}