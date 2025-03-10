package handlers

import (
	"net/http"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateSongByName(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		songName := c.Param("song")
		if songName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing song parameter"})
			return
		}

		var existingSong models.Song
		if err := db.Where("song = ?", songName).First(&existingSong).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve song"})
			}
			return
		}

		var updateData map[string]interface{}
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		if err := db.Model(&existingSong).Updates(updateData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update song"})
			return
		}

		c.JSON(http.StatusOK, existingSong)
	}
}
