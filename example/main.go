package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-i2p/go-meta-listener"
)

func main() {
	// Create a new meta listener
	metaListener := meta.NewMetaListener()
	defer metaListener.Close()

	// Create and add TCP listener
	tcpListener, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		log.Fatalf("Failed to create TCP listener: %v", err)
	}
	if err := metaListener.AddListener("tcp", tcpListener); err != nil {
		log.Fatalf("Failed to add TCP listener: %v", err)
	}
	log.Println("Added TCP listener on 127.0.0.1:8082")

	// Create and add a Unix socket listener (on Unix systems)
	socketPath := "/tmp/example.sock"
	os.Remove(socketPath) // Clean up from previous runs
	unixListener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Printf("Failed to create Unix socket listener: %v", err)
	} else {
		if err := metaListener.AddListener("unix", unixListener); err != nil {
			log.Printf("Failed to add Unix socket listener: %v", err)
		} else {
			log.Println("Added Unix socket listener on", socketPath)
		}
	}
	log.Println("Starting http server...")

	// Create a simple HTTP server using the meta listener
	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello from MetaListener! You connected via: %s\n", r.Proto)
		}),
	}
	log.Println("Server is ready to accept connections...")

	// Handle server shutdown gracefully
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Server starting, listening on multiple transports")
		if err := server.Serve(metaListener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	log.Println("Shutting down server...")

	// Create a deadline for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shut down the HTTP server
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	// Wait for all listener goroutines to exit
	if err := metaListener.WaitForShutdown(ctx); err != nil {
		log.Printf("Timed out waiting for listener shutdown: %v", err)
	}

	log.Println("Server stopped")
}
