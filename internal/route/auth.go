package route

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/senarais/golangquotesapi/internal/service"
)

type AuthRoutes struct{}

func(a *AuthRoutes) SetupRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("views/*")


	// GET Requests
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
		})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "login",
		})
	})

	// POST Request
	router.POST("/register", func(c *gin.Context) {
		request := c.Request

		service := service.RegisterService{}
		token, err := service.Register(request)
		if err != nil {
			c.Error(err)
		}
		
		c.SetCookie("token", token, 60*60*24, "/", "localhost", true, true)
	})
}