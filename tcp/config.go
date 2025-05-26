// Package tcp provides production hardening for net.TCPListener with minimal overhead.
//
// This package wraps standard Go TCP listeners with essential TCP socket configuration
// for internet-facing services. It applies conservative defaults without requiring
// configuration, making it safe to use in production environments.
//
// Example usage:
//
//	listener, err := net.Listen("tcp", ":8080")
//	if err != nil {
//		log.Fatal(err)
//	}
//	tcpListener := listener.(*net.TCPListener)
//
//	hardenedListener, err := tcp.Config(*tcpListener)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer hardenedListener.Close()
//
//	for {
//		conn, err := hardenedListener.Accept()
//		if err != nil {
//			log.Printf("Accept error: %v", err)
//			continue
//		}
//		go handleConnection(conn)
//	}
package tcp

import (
	"net"
	"time"
)

const (
	// keepAliveInterval sets TCP keep-alive probe interval to 15 seconds.
	// This provides reasonable connection health detection without excessive overhead.
	keepAliveInterval = 15 * time.Second

	// socketBufferSize sets both read and write socket buffers to 64KB.
	// This balances memory usage with throughput for typical web applications.
	socketBufferSize = 64 * 1024
)

// hardenedListener wraps net.TCPListener with production hardening features.
type hardenedListener struct {
	listener net.TCPListener
}

// Config wraps a net.TCPListener with production hardening features.
//
// The wrapped listener applies the following enhancements:
// - TCP keep-alive with 15-second intervals
// - TCP_NODELAY for reduced latency
// - 64KB socket buffer sizes for optimal throughput
//
// These settings are chosen to provide reliability and performance improvements
// while maintaining compatibility with standard net.Listener interface.
func Config(listener net.TCPListener) (net.Listener, error) {
	return &hardenedListener{
		listener: listener,
	}, nil
}

// Accept waits for and returns the next connection with hardening applied.
func (hl *hardenedListener) Accept() (net.Conn, error) {
	conn, err := hl.listener.AcceptTCP()
	if err != nil {
		return nil, err
	}

	// Apply TCP hardening settings
	if err := hl.hardenConnection(conn); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

// hardenConnection applies security and performance settings to a TCP connection.
func (hl *hardenedListener) hardenConnection(conn *net.TCPConn) error {
	// Enable TCP keep-alive to detect dead connections
	if err := conn.SetKeepAlive(true); err != nil {
		return err
	}

	// Set keep-alive interval for timely detection of connection issues
	if err := conn.SetKeepAlivePeriod(keepAliveInterval); err != nil {
		return err
	}

	// Disable Nagle's algorithm for lower latency
	if err := conn.SetNoDelay(true); err != nil {
		return err
	}

	// Set socket buffer sizes for optimal throughput
	if err := conn.SetReadBuffer(socketBufferSize); err != nil {
		return err
	}

	if err := conn.SetWriteBuffer(socketBufferSize); err != nil {
		return err
	}

	return nil
}

// Close stops the listener and prevents new connections.
func (hl *hardenedListener) Close() error {
	return hl.listener.Close()
}

// Addr returns the listener's network address.
func (hl *hardenedListener) Addr() net.Addr {
	return hl.listener.Addr()
}
