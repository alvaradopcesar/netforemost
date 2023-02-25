package main

import (
	"netforemost/internal/config"
	"netforemost/internal/http"
	"netforemost/pkg/logger"
	"os"
)

func main() {

	log := logger.New("NoteBook", false)

	log.Info("starting NoteBook service")

	conf := config.GetConfig()

	if err := http.Start(conf, log); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
