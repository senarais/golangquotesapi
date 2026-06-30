package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/senarais/golangquotesapi/internal/service"
)

type QuotesRoute struct{}

func(q *QuotesRoute) SetupRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("views/*")

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/login")
	})

	router.GET("/quote", func(c *gin.Context) {
		service := service.QuotesService{}
		quote := service.GetRandomQuote()

		c.JSON(http.StatusOK, quote)
	})
}