package http

import (
	"net/http"

	"netforemost/pkg/logger"
	"netforemost/pkg/response"
)

// Handler handles status requests.
type Handler struct {
	log logger.Logger
}

// HealthCheckHandler returns the status of the server dependencies.
func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// here we have to validate db connection / grp connection / cache connection
	_ = response.JSON(w, r, http.StatusOK, response.Map{"status": "ok"})
}

// SayHelloHandler returns a greeting.
func (h *Handler) SayHelloHandler(w http.ResponseWriter, r *http.Request) {
	_ = response.JSON(w, r, http.StatusOK, response.Map{"status": "Hello From the Server!"})
}

// New returns a new http handler.
func New(log logger.Logger) *Handler {
	return &Handler{
		log: log,
	}
}
