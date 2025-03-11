package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /

// PingExample godoc
// @Summary Home page
// @Description Home page
// @Tags Home
// @Produce json
// @Success 200
// @Router / [get]
func HomeHandler(c *gin.Context) {
	log.Println("[INFO] Handling request for the home page")

	log.Println("[DEBUG] Rendering index.html with title 'Home Page'")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
	})

	log.Println("[INFO] Home page rendered successfully")
}
