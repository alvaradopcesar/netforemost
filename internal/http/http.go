package http

import (
	//"gitlab.com/prettytechnical/oryx-backend-core/internal/config"
	//"gitlab.com/prettytechnical/oryx-backend-core/internal/db/postgres"
	//"gitlab.com/prettytechnical/oryx-backend-core/internal/db/redis"
	//"gitlab.com/prettytechnical/oryx-backend-core/pkg/logger"
	//"gitlab.com/prettytechnical/oryx-backend-core/pkg/notification"
	"netforemost/internal/config"
	"netforemost/pkg/logger"
)

// Start sets server's handler.
func Start(conf *config.Config, log logger.Logger) error {
	r := routes(conf, log)

	server := newServer(conf.Port, r, log)

	server.Start()

	return nil
}
