package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"netforemost/pkg/logger"
)

// Server is a base server configuration.
type server struct {
	*http.Server
	log logger.Logger
}

// newServer initialize a new server with configuration.
//func newServer(listening string, mux http.Handler, log logger.Logger) *server {
func newServer(listening string, mux http.Handler, log logger.Logger) *server {
	s := &http.Server{
		Addr:         ":" + listening,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server{s, log}
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *server) Start() {
	srv.log.Info("starting API cmd")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.log.Error("could not listen on" + srv.Addr + "due to " + err.Error())
		}
	}()
	srv.log.Info("cmd is ready to handle requests ", srv.Addr)
	srv.gracefulShutdown()
}

func (srv *server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	srv.log.Info("cmd is shutting down ", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		srv.log.Error("could not gracefully shutdown the cmd ", err.Error())
	}
	srv.log.Info("cmd stopped")
}
