package router

import (
	handlers "github.com/mjmhtjain/enpal/src/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Setup initializes all routes and returns the router
func Setup() *gin.Engine {
	router := gin.Default()

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler()

	// Register routes
	router.GET("/health", healthHandler.Check)

	return router
}
