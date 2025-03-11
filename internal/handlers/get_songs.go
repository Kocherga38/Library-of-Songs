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
		log.Println("[INFO] Starting to fetch all songs...")

		var songs []models.Song

		log.Println("[DEBUG] Executing query to fetch all songs")
		rows, err := db.Query("SELECT id, \"group\", song, lyrics FROM songs")
		if err != nil {
			log.Printf("[ERROR] Error while fetching songs: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		defer rows.Close()

		log.Println("[INFO] Iterating over the rows of songs")
		for rows.Next() {
			var song models.Song
			if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.Lyrics); err != nil {
				log.Printf("[ERROR] Error scanning song: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read songs"})
				return
			}

			log.Printf("[DEBUG] Adding song to the list: %s", song.Song)
			songs = append(songs, song)
		}

		if err := rows.Err(); err != nil {
			log.Printf("[ERROR] Row iteration error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve songs"})
			return
		}

		log.Printf("[INFO] Successfully fetched %d songs", len(songs))
		c.JSON(http.StatusOK, songs)
	}
}
