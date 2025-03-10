package routes

import (
	"database/sql"
	"log"

	"github.com/Kocherga38/Library-of-Songs/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	log.Println("[INFO] Setting up routes...")

	router.GET("/", handlers.HomeHandler)
	router.POST("/song", handlers.PostSong(db))
	router.GET("/songs", handlers.GetSongs(db))
	router.DELETE("/song/:song", handlers.DeleteSong(db))
	router.GET("/song/:song", handlers.GetSongByName(db))
	router.PATCH("/song/:song", handlers.UpdateSongByName(db))

	log.Println("[INFO] Routes have been successfully set up")
}
