package routes

import (
	"github.com/Kocherga38/Library-of-Songs/internal/handlers"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/", handlers.HomeHandler)
	router.POST("/song", handlers.CreateSong(db))
}
