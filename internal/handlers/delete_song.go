package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
)

// PostSong godoc
// @Summary Deletes a song
// @Description Deletes a song from the database by its name.
// @Tags Songs
// @Accept json
// @Produce json
// @Param song path string true "Song Name"
// @Success 200 {object} map[string]string "Song deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Missing song parameter"
// @Failure 404 {object} models.ErrorResponse "Song not found"
// @Failure 500 {object} models.ErrorResponse "Failed to delete song"
// @Router /song/{song} [delete]
func DeleteSong(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("[INFO] Starting song deletion process...")

		songName := c.Param("song")
		if songName == "" {
			log.Println("[INFO] Missing song parameter in request")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing song parameter"})
			return
		}

		var song models.Song
		log.Printf("[DEBUG] Querying song with name: %s", songName)
		query := "SELECT id, song, \"group\", verses FROM songs WHERE song = $1"
		err := db.QueryRow(query, songName).Scan(&song.ID, &song.Song, &song.Group, &song.Verses)
		if err == sql.ErrNoRows {
			log.Printf("[INFO] Song with name %s not found", songName)
			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			return
		} else if err != nil {
			log.Printf("[ERROR] Error while finding song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check song existence"})
			return
		}

		log.Printf("[DEBUG] Deleting song with name: %s", songName)
		deleteQuery := "DELETE FROM songs WHERE song = $1"
		_, err = db.Exec(deleteQuery, songName)
		if err != nil {
			log.Printf("[ERROR] Error while deleting song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete song"})
			return
		}

		log.Printf("[INFO] Song with name %s deleted successfully", songName)
		c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully"})
	}
}
