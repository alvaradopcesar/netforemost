package http

import (
	"netforemost/internal/config"
	"netforemost/pkg/logger"
)

// Start sets server's handler.
func Start(conf *config.Config, log logger.Logger) error {
	r := routes(log)

	server := newServer(conf.Port, r, log)

	server.Start()

	return nil
}
