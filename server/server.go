package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setupGracefulShutdown(server *http.Server) {
	signalListener := make(chan os.Signal)

	signal.Notify(signalListener, syscall.SIGINT, syscall.SIGTERM)
	<-signalListener

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
		os.Exit(1)
	}

	log.Println("Server exiting...")

	os.Exit(0)
}

func startServer(server *http.Server) {
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Server failed to initialize")

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

	log.Println("Server is up and kicking")

	setupGracefulShutdown(server)
}
