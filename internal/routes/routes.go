package routes

import (
	"github.com/Kocherga38/Library-of-Songs/internal/handlers"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/", handlers.HomeHandler)
	router.POST("/song", handlers.PostSong(db))
	router.GET("/songs", handlers.GetSongs(db))
	router.GET("/song/:song", handlers.GetSongByName(db))
	router.DELETE("/song/:song", handlers.DeleteSong(db))
	router.PATCH("/song/:song", handlers.UpdateSongByName(db))
}
