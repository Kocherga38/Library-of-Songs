package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/Kocherga38/Library-of-Songs/internal/models"
	"github.com/gin-gonic/gin"
)

func UpdateSongByName(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		songName := c.Param("song")
		if songName == "" {
			log.Println("[INFO] Missing song parameter")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing song parameter"})
			return
		}

		log.Printf("[INFO] Fetching song with name: %s", songName)

		var existingSong models.Song
		query := "SELECT id, song, \"group\" FROM songs WHERE song = $1"
		err := db.QueryRow(query, songName).Scan(&existingSong.ID, &existingSong.Song, &existingSong.Group)
		if err == sql.ErrNoRows {
			log.Printf("[INFO] Song %s not found", songName)
			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			return
		} else if err != nil {
			log.Printf("[ERROR] Error while fetching song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve song"})
			return
		}

		log.Println("[INFO] Binding incoming JSON to update song fields")
		var updateData map[string]interface{}
		if err := c.ShouldBindJSON(&updateData); err != nil {
			log.Printf("[ERROR] Invalid JSON format: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		setValues := []string{}
		params := []interface{}{}
		paramCount := 1

		log.Println("[DEBUG] Preparing update query fields")
		if group, ok := updateData["group"]; ok {
			setValues = append(setValues, "\"group\" = $"+strconv.Itoa(paramCount))
			params = append(params, group)
			paramCount++
		}

		var newSongName string
		if song, ok := updateData["song"]; ok {
			setValues = append(setValues, "song = $"+strconv.Itoa(paramCount))
			newSongName = song.(string)
			params = append(params, song)
			paramCount++
		}

		if len(setValues) == 0 {
			log.Println("[INFO] No valid fields to update")
			c.JSON(http.StatusBadRequest, gin.H{"error": "No valid fields to update"})
			return
		}

		updateQuery := "UPDATE songs SET " + stringJoin(setValues, ", ") + " WHERE song = $" + strconv.Itoa(paramCount)
		params = append(params, songName)

		log.Printf("[DEBUG] Executing query: %s with parameters: %v", updateQuery, params)

		_, err = db.Exec(updateQuery, params...)
		if err != nil {
			log.Printf("[ERROR] Error while updating song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update song"})
			return
		}

		if newSongName == "" {
			newSongName = songName
		}

		log.Printf("[INFO] Fetching updated song with name: %s", newSongName)
		if err := db.QueryRow(query, newSongName).Scan(&existingSong.ID, &existingSong.Song, &existingSong.Group); err != nil {
			log.Printf("[ERROR] Error while fetching updated song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated song"})
			return
		}

		log.Printf("[INFO] Song updated successfully: %+v", existingSong)
		c.JSON(http.StatusOK, existingSong)
	}
}

func stringJoin(a []string, sep string) string {
	result := ""
	for i, s := range a {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}
