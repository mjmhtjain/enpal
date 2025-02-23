package router

import (
	handlers "github.com/mjmhtjain/enpal/src/internal/handlers"

	"github.com/gin-gonic/gin"
)

// Setup initializes all routes and returns the router
func Setup() *gin.Engine {
	router := gin.Default()

	// Register routes
	router.GET("/calendar/query", handlers.NewAppointmentHandler().Find)

	return router
}
