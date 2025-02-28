package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/enpal/src/internal/dto"
	"github.com/mjmhtjain/enpal/src/internal/handlers"
	"github.com/mjmhtjain/enpal/src/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockService := mocks.NewAppointmentService(t)
		handler := &handlers.AppointmentHandler{
			AppointmentService: mockService,
		}

		calendarQuery := dto.CalendarQueryRequestBody{
			Date:     "2023-01-01",
			Products: `["SolarPanels", "Heatpumps"]`,
			Language: "German",
			Rating:   "Bronze",
		}

		mockService.On("FindFreeSlots", mock.Anything).Return([]dto.CalendarQueryResponse{}, nil)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body, _ := json.Marshal(calendarQuery)
		c.Request, _ = http.NewRequest(http.MethodPost, "/calendar/query", bytes.NewBuffer(body))

		handler.Find(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("invalid date format", func(t *testing.T) {
		mockService := mocks.NewAppointmentService(t)
		handler := &handlers.AppointmentHandler{
			AppointmentService: mockService,
		}

		calendarQuery := dto.CalendarQueryRequestBody{
			Date:     "01-01-2023", // Invalid date format, should be YYYY-MM-DD
			Products: `["SolarPanels", "Heatpumps"]`,
			Language: "German",
			Rating:   "Bronze",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body, _ := json.Marshal(calendarQuery)
		c.Request, _ = http.NewRequest(http.MethodPost, "/calendar/query", bytes.NewBuffer(body))

		handler.Find(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("invalid product values", func(t *testing.T) {
		mockService := mocks.NewAppointmentService(t)
		handler := &handlers.AppointmentHandler{
			AppointmentService: mockService,
		}

		calendarQuery := dto.CalendarQueryRequestBody{
			Date:     "2023-01-01",
			Products: `["InvalidProduct"]`, // Product not supported
			Language: "German",
			Rating:   "Bronze",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body, _ := json.Marshal(calendarQuery)
		c.Request, _ = http.NewRequest(http.MethodPost, "/calendar/query", bytes.NewBuffer(body))

		handler.Find(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("malformed product JSON", func(t *testing.T) {
		mockService := mocks.NewAppointmentService(t)
		handler := &handlers.AppointmentHandler{
			AppointmentService: mockService,
		}

		calendarQuery := dto.CalendarQueryRequestBody{
			Date:     "2023-01-01",
			Products: `[SolarPanels]`, // Malformed JSON
			Language: "German",
			Rating:   "Bronze",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body, _ := json.Marshal(calendarQuery)
		c.Request, _ = http.NewRequest(http.MethodPost, "/calendar/query", bytes.NewBuffer(body))

		handler.Find(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("invalid language", func(t *testing.T) {
		mockService := mocks.NewAppointmentService(t)
		handler := &handlers.AppointmentHandler{
			AppointmentService: mockService,
		}

		calendarQuery := dto.CalendarQueryRequestBody{
			Date:     "2023-01-01",
			Products: `["SolarPanels", "Heatpumps"]`,
			Language: "InvalidLanguage", // Unsupported language
			Rating:   "Bronze",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body, _ := json.Marshal(calendarQuery)
		c.Request, _ = http.NewRequest(http.MethodPost, "/calendar/query", bytes.NewBuffer(body))

		handler.Find(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("invalid rating", func(t *testing.T) {
		mockService := mocks.NewAppointmentService(t)
		handler := &handlers.AppointmentHandler{
			AppointmentService: mockService,
		}

		calendarQuery := dto.CalendarQueryRequestBody{
			Date:     "2023-01-01",
			Products: `["SolarPanels", "Heatpumps"]`,
			Language: "German",
			Rating:   "InvalidRating", // Unsupported rating
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body, _ := json.Marshal(calendarQuery)
		c.Request, _ = http.NewRequest(http.MethodPost, "/calendar/query", bytes.NewBuffer(body))

		handler.Find(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("service layer error", func(t *testing.T) {
		mockService := mocks.NewAppointmentService(t)
		handler := &handlers.AppointmentHandler{
			AppointmentService: mockService,
		}

		calendarQuery := dto.CalendarQueryRequestBody{
			Date:     "2023-01-01",
			Products: `["SolarPanels", "Heatpumps"]`,
			Language: "German",
			Rating:   "Bronze",
		}

		// Mock service to return an internal error
		mockService.On("FindFreeSlots", mock.Anything).Return(nil, assert.AnError)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body, _ := json.Marshal(calendarQuery)
		c.Request, _ = http.NewRequest(http.MethodPost, "/calendar/query", bytes.NewBuffer(body))

		handler.Find(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("missing required fields", func(t *testing.T) {
		mockService := mocks.NewAppointmentService(t)
		handler := &handlers.AppointmentHandler{
			AppointmentService: mockService,
		}

		calendarQuery := dto.CalendarQueryRequestBody{
			// Missing Date
			Products: `["SolarPanels", "Heatpumps"]`,
			Language: "German",
			Rating:   "Bronze",
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body, _ := json.Marshal(calendarQuery)
		c.Request, _ = http.NewRequest(http.MethodPost, "/calendar/query", bytes.NewBuffer(body))

		handler.Find(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
