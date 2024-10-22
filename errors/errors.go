package errors

import (
	"encoding/json"
	"github.com/sirupsen/logrus" // Add logrus
	"net/http"
)

// ErrorResponse represents a standard structure for sending error responses.
// swagger:models ErrorResponse
type ErrorResponse struct {
	// The HTTP status code of the error.
	// example: 404
	Code int `json:"code"`
	// A detailed error message.
	// example: "Account not found"
	Message string `json:"message"`
}

// SendErrorResponse sends a structured error response in JSON format.
//
// This function is used to send error responses in a standardized format.
//
// Parameters:
//   - w: The ResponseWriter to send the response to.
//   - code: The HTTP status code for the response.
//   - message: A human-readable message explaining the error.
//
// swagger:response ErrorResponse
func SendErrorResponse(w http.ResponseWriter, code int, message string) {
	// Log the error with logrus
	logrus.WithFields(logrus.Fields{
		"status_code": code,
		"message":     message,
	}).Error("Sending error response")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Code: code, Message: message})
}
