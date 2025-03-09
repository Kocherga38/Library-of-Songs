package handlers

import (
	"net/http"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetSongByName(db gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		songName := c.Param("song")

		var song models.Song
		if err := db.Where("song = ?", songName).First(&song).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve song"})
			}
			return
		}

		c.JSON(http.StatusOK, song)
	}
}
