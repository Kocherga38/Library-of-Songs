package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
)

func GetSongs(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var songs []models.Song

		rows, err := db.Query("SELECT id, \"group\", song FROM songs")
		if err != nil {
			log.Printf("Error while fetching songs: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var song models.Song
			if err := rows.Scan(&song.ID, &song.Group, &song.Song); err != nil {
				log.Printf("Error scanning song: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read songs"})
				return
			}

			songs = append(songs, song)
		}

		if err := rows.Err(); err != nil {
			log.Printf("Row iteration error: %v", err) // Логируем ошибку
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve songs"})
			return
		}

		c.JSON(http.StatusOK, songs)
	}
}
