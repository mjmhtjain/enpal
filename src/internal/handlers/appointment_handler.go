package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/enpal/src/internal/dto"
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
	var calendarQueryBody dto.CalendarQueryRequestBody
	err := json.NewDecoder(c.Request.Body).Decode(&calendarQueryBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	calendarQueryObj, err := calendarQueryBody.GetDomainObject()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	res, err := h.appointmentService.FindFreeSlots(calendarQueryObj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
