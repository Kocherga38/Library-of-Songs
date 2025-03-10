package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
)

func PostSong(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var song models.Song

		if err := c.ShouldBindJSON(&song); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		var existingSong models.Song
		query := "SELECT id, \"group\", song FROM songs WHERE song = $1"
		err := db.QueryRow(query, song.Song).Scan(&existingSong.ID, &existingSong.Song, &existingSong.Group)
		if err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Song with this name already exists"})
			return
		}

		if err != sql.ErrNoRows {
			log.Println("Error while checking song existence:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check song existence"})
			return
		}

		insertQuery := "INSERT INTO songs (\"group\", song) VALUES ($1, $2) RETURNING id"
		var newID int
		err = db.QueryRow(insertQuery, song.Group, song.Song).Scan(&newID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create song"})
			return
		}

		song.ID = newID
		c.JSON(http.StatusCreated, song)
	}
}
