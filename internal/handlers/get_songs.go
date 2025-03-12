package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
)

// GetSongs godoc
// @Summary Gets all songs
// @Description This endpoint allows you to get all songs from the database.
// @Tags Songs
// @Accept json
// @Produce json
// @Success 200 {array} models.Song "List of songs"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /songs [get]
func GetSongs(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("[INFO] Starting to fetch all songs...")

		var songs []models.Song

		log.Println("[DEBUG] Executing query to fetch all songs")
		rows, err := db.Query("SELECT id, \"group\", song, verses FROM songs")
		if err != nil {
			log.Printf("[ERROR] Error while fetching songs: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		defer rows.Close()

		log.Println("[INFO] Iterating over the rows of songs")
		for rows.Next() {
			var song models.Song
			var verses string

			if err := rows.Scan(&song.ID, &song.Group, &song.Song, &verses); err != nil {
				log.Printf("[ERROR] Error scanning song: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read songs"})
				return
			}

			song.Verses = strings.Split(verses, "\n")

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
