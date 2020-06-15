package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hfantin/goblank/config"
	"github.com/hfantin/goblank/server"
)

func main() {
	// Load environment configuration
	config.Load()

	// Init server
	srv := server.New()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("Failed to start server on port", config.Env.ServerPort)
			log.Println("Error: ", err.Error())
		}
	}()
	log.Println("Running on port", config.Env.ServerPort)

	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		log.Println("Good bye")
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown the server:%+v", err)
	}
	log.Println("Stopping the server")
}
