package handlers

import (
	"net/http"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetSongs(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var songs []models.Song

		if err := db.Find(&songs).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve songs"})
			return
		}

		c.JSON(http.StatusOK, songs)
	}
}
