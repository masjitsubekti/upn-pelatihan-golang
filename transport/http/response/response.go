package response

import (
	"encoding/json"
	"net/http"

	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/shared/logger"
)

// Base is the base object of all responses
type Base struct {
	Data    *interface{} `json:"data,omitempty"`
	Error   *string      `json:"error,omitempty"`
	Message *string      `json:"message,omitempty"`
	Success *bool        `json:"success,omitempty"`
}

// NoContent sends a response without any content
func NoContent(w http.ResponseWriter) {
	respond(w, http.StatusNoContent, nil)
}

// WithMessage sends a response with a simple text message
func WithMessage(w http.ResponseWriter, code int, message string) {
	respond(w, code, Base{Message: &message})
}

// WithMessage sends a response with a simple text message
func WithStatusMessage(w http.ResponseWriter, code int, status bool, message string) {
	respond(w, code, Base{
		Message: &message,
		Success: &status,
	})
}

// WithJSON sends a response containing a JSON object
func WithJSON(w http.ResponseWriter, code int, jsonPayload interface{}) {
	respond(w, code, Base{Data: &jsonPayload})
}

// WithError sends a response with an error message
func WithError(w http.ResponseWriter, err error) {
	code := failure.GetCode(err)
	errMsg := err.Error()
	respond(w, code, Base{Error: &errMsg})
}

// WithPreparingShutdown sends a default response for when the server is preparing to shut down
func WithPreparingShutdown(w http.ResponseWriter) {
	WithMessage(w, http.StatusServiceUnavailable, "SERVER PREPARING TO SHUT DOWN")
}

// WithUnhealthy sends a default response for when the server is unhealthy
func WithUnhealthy(w http.ResponseWriter) {
	WithMessage(w, http.StatusServiceUnavailable, "SERVER UNHEALTHY")
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		logger.ErrorWithStack(err)
	}
}

// API Ekternal
type BaseApi struct {
	Data    *interface{} `json:"data,omitempty"`
	Success *bool        `json:"success,omitempty"`
	Message *string      `json:"message,omitempty"`
}

// WithJSON sends a response containing a JSON object
func WithJSONApi(w http.ResponseWriter, code int, jsonPayload interface{}) {
	respond(w, code, jsonPayload)
}
