package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create default gin router
	router := gin.Default()

	// Register health check endpoint
	router.GET("/health", healthCheck)

	// Start the server on port 8080
	router.Run(":8080")
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
