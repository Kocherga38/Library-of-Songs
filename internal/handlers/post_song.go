package handlers

import (
	"net/http"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostSong(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var song models.Song

		if err := c.ShouldBindJSON(&song); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		var existingSong models.Song
		if err := db.Where("song = ?", song.Song).First(&existingSong).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Song with this name already exists"})
			return
		}

		if err := db.Create(&song).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create song"})
			return
		}

		c.JSON(http.StatusCreated, song)
	}
}
