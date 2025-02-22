package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/enpal/src/internal/repository"
)

// AppointmentHandler handles health check related endpoints
type AppointmentHandler struct {
	appointmentRepo repository.IAppointmentRepo
}

// NewAppointmentHandler creates a new instance of AppointmentHandler
func NewAppointmentHandler() *AppointmentHandler {
	return &AppointmentHandler{
		appointmentRepo: repository.NewAppointmentRepo(),
	}
}

// Find finds all the appointment opening
func (h *AppointmentHandler) Find(c *gin.Context) {

	res, err := h.appointmentRepo.FindFreeSlots("2024-05-03")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, res)
}
