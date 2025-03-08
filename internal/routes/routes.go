package routes

import (
	"github.com/Kocherga38/Library-of-Songs/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", handlers.HomeHandler)
}
