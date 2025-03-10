package main

import (
	"database/sql"
	"log"

	"github.com/Kocherga38/Library-of-Songs/internal/database"
	"github.com/Kocherga38/Library-of-Songs/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("[INFO] Starting application...")

	log.Println("[INFO] Initializing database connection...")
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect to database: %v", err)
	}
	defer func() {
		log.Println("[INFO] Closing database connection...")
		db.Close()
	}()

	log.Println("[DEBUG] Database connection established successfully.")

	log.Println("[INFO] Setting up the router...")
	router := setupRouter(db)

	log.Println("[INFO] Server is successfully started on port 3000!")

	if err := router.Run(":3000"); err != nil {
		log.Fatalf("[ERROR] Server failed to start: %v", err)
	}
}

func setupRouter(db *sql.DB) *gin.Engine {
	log.Println("[INFO] Loading HTML templates for the router...")

	router := gin.Default()
	router.LoadHTMLGlob("./web/templates/*")

	log.Println("[INFO] Setting up routes...")
	routes.SetupRoutes(router, db)

	log.Println("[DEBUG] Router setup complete.")

	return router
}
