package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

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
		query := "SELECT id, \"group\", song, lyrics FROM songs WHERE song = $1"
		err := db.QueryRow(query, song.Song).Scan(&existingSong.ID, &existingSong.Song, &existingSong.Group, &existingSong.Lyrics)
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

		log.Printf("[DEBUG] Received song data: %+v", song)
		log.Printf("[DEBUG] Inserting song: Group=%s, Song=%s, Lyrics=%s", song.Group, song.Song, song.Lyrics)
		insertQuery := "INSERT INTO songs (\"group\", song, lyrics) VALUES ($1, $2, $3) RETURNING id"
		var newID int
		err = db.QueryRow(insertQuery, song.Group, song.Song, song.Lyrics).Scan(&newID)
		if err != nil {
			log.Printf("[ERROR] Failed to create song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create song"})
			return
		}

		song.ID = newID
		log.Printf("[INFO] Song created successfully with ID %d", song.ID)

		log.Printf("[INFO] Creating song page")
		err = createSongPage(song)
		if err != nil {
			log.Printf("[ERROR] Failed to create song page: %v", err)
		}
		log.Printf("[INFO] Song page created successfully")

		c.JSON(http.StatusCreated, song)
	}
}

// TODO: write info and update logs
func createSongPage(song models.Song) error {
	template, err := template.ParseFiles("templates/song.html")
	if err != nil {
		return err
	}

	outputDir := "public/songs"
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	filename := filepath.Join(outputDir, fmt.Sprintf("%s.html", song.Song))
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = template.Execute(file, gin.H{
		"Song":   song.Song,
		"Group":  song.Group,
		"Lyrics": song.Lyrics,
	})

	if err != nil {
		return err
	}

	return nil
}
