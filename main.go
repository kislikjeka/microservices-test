package main

import (
	"context"
	"github.com/kislikjeka/microservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)
func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(logger)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	server := &http.Server{
		Addr: ":8080",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		logger.Println("Starting server on port 8080")

		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	logger.Println("Recieved terminate, gracful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30* time.Second)
	server.Shutdown(tc)
}