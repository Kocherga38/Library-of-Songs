package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
)

func GetSongByName(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		songName := c.Param("song")

		var song models.Song
		query := "SELECT id, song, \"group\" FROM songs WHERE song = $1"
		err := db.QueryRow(query, songName).Scan(&song.ID, &song.Song, &song.Group)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			return
		} else if err != nil {
			log.Printf("Error while fetching song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve song"})
			return
		}

		c.JSON(http.StatusOK, song)
	}
}
