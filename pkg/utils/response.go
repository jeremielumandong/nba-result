// Package utils provides shared utility functions for HTTP responses and common operations.
package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

// APIResponse represents a standardized API response structure.
type APIResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Message   string      `json:"message,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// APIError represents an API error response.
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// SendSuccessResponse sends a successful JSON response.
func SendSuccessResponse(w http.ResponseWriter, data interface{}, message string) {
	response := APIResponse{
		Success:   true,
		Data:      data,
		Message:   message,
		Timestamp: time.Now().UTC(),
	}

	sendJSONResponse(w, http.StatusOK, response)
}

// SendErrorResponse sends an error JSON response.
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := APIResponse{
		Success:   false,
		Error:     message,
		Timestamp: time.Now().UTC(),
	}

	sendJSONResponse(w, statusCode, response)
}

// SendValidationErrorResponse sends a validation error response.
func SendValidationErrorResponse(w http.ResponseWriter, errors []string) {
	response := APIResponse{
		Success:   false,
		Error:     "Validation failed",
		Data:      map[string][]string{"validation_errors": errors},
		Timestamp: time.Now().UTC(),
	}

	sendJSONResponse(w, http.StatusBadRequest, response)
}

// sendJSONResponse is a helper function to send JSON responses.
func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		LogErrorf("Failed to encode JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// IsValidJSON checks if a string is valid JSON.
func IsValidJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// PrettyPrintJSON returns a pretty-printed JSON string.
func PrettyPrintJSON(data interface{}) (string, error) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}