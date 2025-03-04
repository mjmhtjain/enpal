package router

import (
	handlers "github.com/mjmhtjain/enpal/src/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Setup initializes all routes and returns the router
func Setup() *gin.Engine {
	router := gin.Default()

	// Register routes
	router.GET("/health", handlers.NewHealthHandler().Check)

	// Register routes
	router.POST("/calendar/query", handlers.NewAppointmentHandler().Find)

	return router
}
