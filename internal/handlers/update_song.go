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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing song parameter"})
			return
		}

		// Check if the song exists in the database
		var existingSong models.Song
		query := "SELECT id, song, \"group\" FROM songs WHERE song = $1"
		err := db.QueryRow(query, songName).Scan(&existingSong.ID, &existingSong.Song, &existingSong.Group)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			return
		} else if err != nil {
			log.Printf("Error while fetching song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve song"})
			return
		}

		// Parse the incoming JSON body
		var updateData map[string]interface{}
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		// Prepare SQL update query dynamically
		setValues := []string{}
		params := []interface{}{}
		paramCount := 1

		// Check if the group field needs to be updated
		if group, ok := updateData["group"]; ok {
			setValues = append(setValues, "\"group\" = $"+strconv.Itoa(paramCount))
			params = append(params, group)
			paramCount++
		}

		// Check if the song field needs to be updated
		var newSongName string
		if song, ok := updateData["song"]; ok {
			setValues = append(setValues, "song = $"+strconv.Itoa(paramCount))
			newSongName = song.(string)
			params = append(params, song)
			paramCount++
		}

		// If there are no valid fields to update, return an error
		if len(setValues) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No valid fields to update"})
			return
		}

		// Build the UPDATE query
		updateQuery := "UPDATE songs SET " + stringJoin(setValues, ", ") + " WHERE song = $" + strconv.Itoa(paramCount)
		params = append(params, songName) // Append the old songName as the last parameter for WHERE clause

		// Log the query and parameters for debugging
		log.Printf("Executing query: %s with parameters: %v", updateQuery, params)

		// Execute the update query
		_, err = db.Exec(updateQuery, params...)
		if err != nil {
			log.Printf("Error while updating song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update song"})
			return
		}

		// Fetch the updated song using the new song name (if it was changed)
		if newSongName == "" {
			// If the song name wasn't updated, use the original songName
			newSongName = songName
		}

		// Retrieve the updated song using the new song name
		if err := db.QueryRow(query, newSongName).Scan(&existingSong.ID, &existingSong.Song, &existingSong.Group); err != nil {
			log.Printf("Error while fetching updated song: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated song"})
			return
		}

		c.JSON(http.StatusOK, existingSong)
	}
}

// Utility function to join a slice of strings with a separator
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
