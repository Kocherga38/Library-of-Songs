package main

import (

	// _ "github.com/Kocherga38/Library-of-Songs/internal/database/migrations"

	"log"

	"github.com/Kocherga38/Library-of-Songs/internal/config"
	"github.com/Kocherga38/Library-of-Songs/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	router := gin.Default()
	router.LoadHTMLGlob("./web/templates/*")

	routes.SetupRoutes(router)

	log.Println("Server is successfully started!")
	router.Run(":3000")
}
