package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
)

// GetSongByName godoc
// @Summary Retrieves a song by name
// @Description Fetches a song from the database by its name and returns it as an HTML page.
// @Tags Songs
// @Accept json
// @Produce html
// @Param song path string true "Song Name"
// @Success 200 {string} string "HTML page with song details"
// @Failure 404 {object} models.ErrorResponse "Song not found"
// @Failure 500 {object} models.ErrorResponse "Failed to retrieve song"
// @Router /song/{song} [get]
func GetSongByName(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("[INFO] Starting song retrieval process...")

		songName := c.Param("song")
		log.Printf("[INFO] Fetching song with name: %s", songName)

		var song models.Song
		query := "SELECT id, song, \"group\", lyrics FROM songs WHERE song = $1"
		err := db.QueryRow(query, songName).Scan(&song.ID, &song.Song, &song.Group, &song.Lyrics)

		if err == sql.ErrNoRows {
			log.Printf("[INFO] Song with name %s not found", songName)
			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			return
		} else if err != nil {
			log.Printf("[ERROR] Error while fetching song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve song"})
			return
		}

		log.Printf("[INFO] Successfully fetched song: %s", songName)
		c.HTML(http.StatusOK, "song.html", song)
	}
}
