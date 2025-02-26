package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/enpal/src/internal/domain"
	"github.com/mjmhtjain/enpal/src/internal/repository"
	"github.com/mjmhtjain/enpal/src/internal/service"
)

// AppointmentHandler handles health check related endpoints
type AppointmentHandler struct {
	appointmentService service.IAppointmentService
}

// NewAppointmentHandler creates a new instance of AppointmentHandler
func NewAppointmentHandler() *AppointmentHandler {
	appointmentRepo := repository.NewAppointmentRepo()
	appointmentService := service.NewAppointmentService(appointmentRepo)

	return &AppointmentHandler{
		appointmentService: appointmentService,
	}
}

// Find finds all the appointment opening
func (h *AppointmentHandler) Find(c *gin.Context) {
	res, err := h.appointmentService.FindFreeSlots(domain.CalendarQueryDomain{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return // Add return statement to prevent executing the next line after error
	}

	c.JSON(http.StatusOK, res)
}
