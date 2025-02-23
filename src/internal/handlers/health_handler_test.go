package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler_Check(t *testing.T) {
	// Switch to test mode
	gin.SetMode(gin.TestMode)

	// Create test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Create handler
	handler := NewHealthHandler()

	// Call the handler
	handler.Check(c)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "healthy")
}
