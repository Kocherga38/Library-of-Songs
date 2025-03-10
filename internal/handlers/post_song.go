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
		log.Println("[INFO] Starting the song creation process...")

		var song models.Song

		log.Println("[DEBUG] Binding incoming JSON to song struct")
		if err := c.ShouldBindJSON(&song); err != nil {
			log.Printf("[ERROR] Invalid JSON format: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		var existingSong models.Song
		log.Printf("[INFO] Checking if song %s already exists", song.Song)
		query := "SELECT id, \"group\", song FROM songs WHERE song = $1"
		err := db.QueryRow(query, song.Song).Scan(&existingSong.ID, &existingSong.Song, &existingSong.Group)
		if err == nil {
			log.Printf("[INFO] Song with name %s already exists", song.Song)
			c.JSON(http.StatusConflict, gin.H{"error": "Song with this name already exists"})
			return
		}

		if err != sql.ErrNoRows {
			log.Printf("[ERROR] Error while checking song existence: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check song existence"})
			return
		}

		log.Printf("[DEBUG] Inserting new song: %s", song.Song)
		insertQuery := "INSERT INTO songs (\"group\", song) VALUES ($1, $2) RETURNING id"
		var newID int
		err = db.QueryRow(insertQuery, song.Group, song.Song).Scan(&newID)
		if err != nil {
			log.Printf("[ERROR] Failed to create song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create song"})
			return
		}

		song.ID = newID
		log.Printf("[INFO] Song created successfully with ID %d", song.ID)

		c.JSON(http.StatusCreated, song)
	}
}
