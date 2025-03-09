package handlers

import (
	"net/http"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteSong(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		songName := c.Param("song")
		if songName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing song parametr"})
			return
		}

		var song models.Song
		if err := db.Where("song = ?", songName).First(&song).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if err := db.Delete(&song).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully"})
	}
}
