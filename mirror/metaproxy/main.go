package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-i2p/go-meta-listener/mirror"
)

const (
	maxConcurrentConnections = 100 // Limit concurrent connections
	connectionTimeout        = 30 * time.Second
	shutdownTimeout          = 5 * time.Second
)

// connectionPool manages concurrent connections with proper lifecycle
type connectionPool struct {
	semaphore   chan struct{}
	activeConns sync.WaitGroup
	ctx         context.Context
	cancel      context.CancelFunc
}

func newConnectionPool(maxConns int) *connectionPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &connectionPool{
		semaphore: make(chan struct{}, maxConns),
		ctx:       ctx,
		cancel:    cancel,
	}
}

func (cp *connectionPool) handleConnection(clientConn net.Conn, targetHost string, targetPort int) {
	// Acquire semaphore slot or block
	select {
	case cp.semaphore <- struct{}{}:
		// Got slot, continue
	case <-cp.ctx.Done():
		clientConn.Close()
		return
	}

	// Track active connection
	cp.activeConns.Add(1)

	// Handle connection in separate goroutine
	go func() {
		defer func() {
			<-cp.semaphore // Release semaphore slot
			cp.activeConns.Done()
			clientConn.Close()
		}()

		// Set connection timeout
		clientConn.SetDeadline(time.Now().Add(connectionTimeout))

		// Connect to target with timeout
		serverConn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", targetHost, targetPort), 10*time.Second)
		if err != nil {
			log.Printf("Failed to connect to target %s:%d: %v", targetHost, targetPort, err)
			return
		}
		defer serverConn.Close()

		// Set timeout on server connection
		serverConn.SetDeadline(time.Now().Add(connectionTimeout))

		// Create context for this connection
		connCtx, connCancel := context.WithCancel(cp.ctx)
		defer connCancel()

		// Forward data bidirectionally with proper error handling
		var wg sync.WaitGroup
		wg.Add(2)

		// Client to server
		go func() {
			defer wg.Done()
			if _, err := copyWithContext(connCtx, serverConn, clientConn); err != nil && err != io.EOF {
				log.Printf("Error copying client to server: %v", err)
			}
			// Close server write side to signal completion
			if tcpConn, ok := serverConn.(*net.TCPConn); ok {
				tcpConn.CloseWrite()
			}
		}()

		// Server to client
		go func() {
			defer wg.Done()
			if _, err := copyWithContext(connCtx, clientConn, serverConn); err != nil && err != io.EOF {
				log.Printf("Error copying server to client: %v", err)
			}
			// Close client write side to signal completion
			if tcpConn, ok := clientConn.(*net.TCPConn); ok {
				tcpConn.CloseWrite()
			}
		}()

		// Wait for either copy operation to complete or context cancellation
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			// Normal completion
		case <-connCtx.Done():
			// Context cancelled, connections will be closed by defers
		}
	}()
}

func (cp *connectionPool) shutdown() {
	cp.cancel()
	cp.activeConns.Wait()
}

// copyWithContext copies data between connections with context cancellation support
func copyWithContext(ctx context.Context, dst, src net.Conn) (int64, error) {
	// Use a small buffer for responsive cancellation
	buf := make([]byte, 32*1024)
	var written int64

	for {
		select {
		case <-ctx.Done():
			return written, ctx.Err()
		default:
		}

		// Set short read timeout for responsiveness
		src.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = fmt.Errorf("invalid write count")
				}
			}
			written += int64(nw)
			if ew != nil {
				return written, ew
			}
			if nr != nw {
				return written, io.ErrShortWrite
			}
		}
		if er != nil {
			if netErr, ok := er.(net.Error); ok && netErr.Timeout() {
				continue // Retry on timeout
			}
			if er != io.EOF {
				return written, er
			}
			break
		}
	}
	return written, nil
}

// main function sets up a meta listener that forwards connections to a specified host and port.
// It listens for incoming connections and forwards them to the specified destination.
func main() {
	host := flag.String("host", "localhost", "Host to forward connections to")
	port := flag.Int("port", 8080, "Port to forward connections to")
	listenPort := flag.Int("listen-port", 3002, "Port to listen for incoming connections")
	domain := flag.String("domain", "i2pgit.org", "Domain name for TLS listener")
	email := flag.String("email", "", "Email address for Let's Encrypt registration")
	certDir := flag.String("certdir", "./certs", "Directory for storing certificates")
	hiddenTls := flag.Bool("hidden-tls", false, "Enable hidden TLS")
	maxConns := flag.Int("max-conns", maxConcurrentConnections, "Maximum concurrent connections")
	flag.Parse()

	mirror.CERT_DIR = *certDir
	mirror.HIDDEN_TLS = *hiddenTls
	addr := net.JoinHostPort(*domain, fmt.Sprintf("%d", *listenPort))

	// Create connection pool with specified limits
	pool := newConnectionPool(*maxConns)
	defer pool.shutdown()

	// Create a new meta listener
	metaListener, err := mirror.Listen(addr, *email)
	if err != nil {
		log.Fatalf("Failed to create meta listener: %v", err)
	}
	defer metaListener.Close()

	// Set up graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	log.Printf("Proxy server starting, forwarding to %s:%d (max concurrent connections: %d)", *host, *port, *maxConns)

	// Start accepting connections in a separate goroutine
	go func() {
		for {
			conn, err := metaListener.Accept()
			if err != nil {
				// Check if this is due to shutdown
				select {
				case <-pool.ctx.Done():
					log.Println("Shutting down connection accept loop")
					return
				default:
					log.Printf("Error accepting connection: %v", err)
					continue
				}
			}

			log.Printf("Accepted connection from %s", conn.RemoteAddr())
			pool.handleConnection(conn, *host, *port)
		}
	}()

	// Wait for shutdown signal
	<-sigCh
	log.Println("Shutdown signal received, stopping proxy...")

	// Close listener to stop accepting new connections
	metaListener.Close()

	// Shutdown connection pool with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdownCancel()

	done := make(chan struct{})
	go func() {
		pool.shutdown()
		close(done)
	}()

	select {
	case <-done:
		log.Println("All connections closed gracefully")
	case <-shutdownCtx.Done():
		log.Println("Shutdown timeout exceeded, forcing exit")
	}

	log.Println("Proxy server stopped")
}
