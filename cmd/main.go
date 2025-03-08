package main

import (

	// _ "github.com/Kocherga38/Library-of-Songs/internal/database/migrations"

	"github.com/Kocherga38/Library-of-Songs/internal/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./web/templates/*")

	routes.SetupRoutes(router)

	router.Run(":3000")
}
