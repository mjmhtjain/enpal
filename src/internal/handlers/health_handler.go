package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check related endpoints
type HealthHandler struct{}

// NewHealthHandler creates a new instance of HealthHandler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Check handles the health check endpoint
func (h *HealthHandler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
