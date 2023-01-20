package rest

import (
	"fmt"
	"godb/src/godb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	E *gin.Engine
}

func NewRest(db *godb.GODB) *Rest {
	a := new(Rest)

	a.E = gin.Default()

	a.E.GET("/ping", func(c *gin.Context) {

		result := db.Ping()

		fmt.Println(result)

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// a.e.Run(":8201") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return a
}
