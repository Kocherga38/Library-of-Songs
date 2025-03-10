package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
)

func DeleteSong(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		songName := c.Param("song")
		if songName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing song parametr"})
			return
		}

		var song models.Song
		query := "SELECT id, song, \"group\" FROM songs WHERE song = $1"
		err := db.QueryRow(query, songName).Scan(&song.ID, &song.Song, &song.Group)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			return
		} else if err != nil {
			log.Printf("Error while finding song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check song existence"})
			return
		}

		deleteQuery := "DELETE FROM songs WHERE song = $1"
		_, err = db.Exec(deleteQuery, songName)
		if err != nil {
			log.Printf("Error while deleting song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete song"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully"})
	}
}
