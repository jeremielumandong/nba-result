package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Create a request
	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	assert.NoError(t, err)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	HealthCheck(rr, req)

	// Check status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check content type
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Parse response
	var response HealthResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Check response fields
	assert.Equal(t, "healthy", response.Status)
	assert.Equal(t, "1.0.0", response.Version)
	assert.NotZero(t, response.Timestamp)
}