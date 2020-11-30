package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/0xkalvin/url-shortener/logger"
)

func setupGracefulShutdown(server *http.Server) {
	logger := log.GetLogger()

	signalListener := make(chan os.Signal)

	signal.Notify(signalListener, syscall.SIGINT, syscall.SIGTERM)
	<-signalListener

	logger.Info("Gracefully shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		logger.Fatal("Server forced to shutdown:", err)
		os.Exit(1)
	}

	logger.Info("Server closed. Exiting process...")

	os.Exit(0)
}

func startServer(server *http.Server) {

	err := server.ListenAndServe()

	if err != nil {
		logger := log.GetLogger()

		logger.Fatal("Server failed to initialize")

		os.Exit(1)
	}
}

// Run app
func Run() {
	var address = fmt.Sprintf(":%s", os.Getenv("PORT"))

	router := initializeRouter()

	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	go startServer(server)

	logger := log.GetLogger()

	logger.Info("Server is up and kicking")

	setupGracefulShutdown(server)
}
