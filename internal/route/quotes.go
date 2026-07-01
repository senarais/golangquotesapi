package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/senarais/golangquotesapi/internal/middleware"
	"github.com/senarais/golangquotesapi/internal/service"
)

type QuotesRoute struct{}

func(q *QuotesRoute) SetupRoutes(router *gin.Engine) {
	router.Use(middleware.Verify())
	
	router.LoadHTMLGlob("views/*")


	router.GET("/quote", func(c *gin.Context) {
		c.HTML(http.StatusOK, "fetch_quote.html", gin.H{
			"title": "Quote",
		})
	})

	router.POST("/quote", func(c *gin.Context) {
		service := service.QuotesService{}
		quote := service.GetRandomQuote()

		c.JSON(http.StatusOK, quote)
	})
}