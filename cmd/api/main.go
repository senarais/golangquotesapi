package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/senarais/golangquotesapi/internal/route"
)

func main()  {
	router := gin.Default()

	// Import Routes
	authRoutes := route.AuthRoutes{}
	quotesRoute := route.QuotesRoute{}

	// Setup Routes
	authRoutes.SetupRoutes(router)
	quotesRoute.SetupRoutes(router)

	fmt.Printf(`Server running on port 5000
	    (  )   (   )  )
     ) (   )  (  (
     ( )  (    ) )
     _____________
    <_____________> ___
    |             |/ _ \
    |               | | |
    |               |_| |
 ___|             |\___/
/    \___________/    \
\_____________________/

	`)

	router.Run("localhost:5000")
}