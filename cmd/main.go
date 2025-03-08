package main

import (

	// _ "github.com/Kocherga38/Library-of-Songs/internal/database/migrations"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./web/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	router.Run(":3000")
}
