package main

import (
	"database/sql"
	"log"

	"github.com/Kocherga38/Library-of-Songs/internal/database"
	"github.com/Kocherga38/Library-of-Songs/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	router := setupRouter(db)

	log.Println("Server is successfully started!")
	router.Run(":3000")
}

func setupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("./web/templates/*")

	routes.SetupRoutes(router, db)

	return router
}
