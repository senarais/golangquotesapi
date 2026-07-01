package route

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/senarais/golangquotesapi/internal/model"
	"github.com/senarais/golangquotesapi/internal/service"
)

type AuthRoutes struct{}

func(a *AuthRoutes) SetupRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("views/*")


	// GET Requests
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/login")
	})

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
		user := model.User{}

		err:= json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		service := service.RegisterService{}
		token, err := service.Register(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		
		c.SetCookie("token", token, 60*60*24, "/", "localhost", false, true)
		c.Redirect(http.StatusFound, "/quote")
	})

	router.POST("/login", func(c *gin.Context) {
		request := c.Request
		user := model.User{}

		err:= json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		service := service.LoginService{}
		token, err := service.Login(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		
		c.SetCookie("token", token, 60*60*24, "/", "localhost", false, true)
		c.Redirect(http.StatusFound, "/quote")
	})
}