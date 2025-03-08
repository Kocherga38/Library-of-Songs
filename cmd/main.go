package main

import (

	// _ "github.com/Kocherga38/Library-of-Songs/internal/database/migrations"

	"log"

	"github.com/Kocherga38/Library-of-Songs/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// db, err := config.InitDB()
	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }

	router := setupRouter()

	log.Println("Server is successfully started!")
	router.Run(":3000")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("./web/templates/*")

	routes.SetupRoutes(router)

	return router
}
