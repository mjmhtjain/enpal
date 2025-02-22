package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/enpal/src/internal/client"
)

// AppointmentHandler handles health check related endpoints
type AppointmentHandler struct {
	dbClient *sql.DB
}

// NewAppointmentHandler creates a new instance of AppointmentHandler
func NewAppointmentHandler() *AppointmentHandler {
	db, err := client.NewDBClient(client.NewDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err.Error())
	}

	return &AppointmentHandler{
		dbClient: db,
	}
}

// Find finds all the appointment opening
func (h *AppointmentHandler) Find(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
