package main

import (
	"fmt"

	"github.com/senarais/golangquotesapi/internal/route"
)

func main()  {
	quotesRoute := route.QuotesRoute{}
	routes := quotesRoute.SetupRoutes()

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

	routes.Run("localhost:5000")
}