package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/senarais/golangquotesapi/internal/service"
)

type QuotesRoute struct{}

func(q *QuotesRoute) SetupRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"message": "API is active, try /quote endpoint."})
	})

	router.GET("/quote", func(c *gin.Context) {
		service := service.QuotesService{}
		quote := service.GetRandomQuote()

		c.JSON(http.StatusOK, quote)
	})
}