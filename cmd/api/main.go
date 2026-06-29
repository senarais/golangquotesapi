package main

import (
	"fmt"
	
	"github.com/gin-gonic/gin"

)

func main()  {
	r:= gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "This API is active.",
		})
	})

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
	
	r.Run("localhost:5000")

}