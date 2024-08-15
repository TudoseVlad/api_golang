package api

import (
	"context"
	"fmt"
	"golang_api/src/data"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitApi() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}
	http.HandleFunc("/cuvinte", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			POSTHandler(w, r)
		case http.MethodGet:
			GETHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	// Start the server in a goroutine
	go func() {
		fmt.Println("Listener started")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()

	<-stopChan
	fmt.Println("\nShutting down server...")
	data.DumpData()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exiting")
}
