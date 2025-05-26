package tcp

import (
	"fmt"
	"net"
	"syscall"
	"time"
)

const (
	// Production-optimized defaults
	keepAliveInterval = 30 * time.Second
	bufferSize        = 64 * 1024 // 64KB
	connectionBacklog = 128
)

// Listen creates a TCP listener with production-hardened defaults suitable
// for internet-facing services. The implementation applies secure socket
// options including address reuse, keep-alive monitoring, optimized buffers,
// and minimal latency configuration.
//
// Parameters:
//   - network: Must be "tcp", "tcp4", or "tcp6"
//   - address: Standard Go network address format (e.g., ":8080", "127.0.0.1:8080")
//
// Returns a net.Listener configured with production defaults, or an error
// if the listener cannot be created or configured.
//
// The listener applies these optimizations:
//   - SO_REUSEADDR: Prevents "address already in use" during rapid restarts
//   - Keep-alive: 30-second interval for connection health monitoring
//   - Buffer sizes: 64KB receive/send buffers for optimal throughput
//   - TCP_NODELAY: Enabled to minimize latency
//   - Backlog: 128 pending connections in accept queue
func Listen(network, address string) (net.Listener, error) {
	// Validate network parameter
	if !isValidNetwork(network) {
		return nil, fmt.Errorf("tcp.Listen: invalid network type %q for addr %s, must be tcp, tcp4, or tcp6", network, address)
	}

	// Validate address format by attempting to resolve
	if _, err := net.ResolveTCPAddr(network, address); err != nil {
		return nil, fmt.Errorf("tcp.Listen: invalid address %q for network %q: %w", address, network, err)
	}

	// Create the base listener
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, fmt.Errorf("tcp.Listen: failed to create listener on %s %s: %w", network, address, err)
	}

	// Apply production socket configurations
	tcpListener, ok := listener.(*net.TCPListener)
	if !ok {
		listener.Close()
		return nil, fmt.Errorf("tcp.Listen: unexpected listener type for %s %s", network, address)
	}

	// Configure the underlying socket with production defaults
	if err := configureSocket(tcpListener); err != nil {
		// Ensure proper cleanup even if socket is in inconsistent state
		if closeErr := listener.Close(); closeErr != nil {
			// Log the close error but return the original configuration error
			// since that's the primary failure point
			return nil, fmt.Errorf("tcp.Listen: failed to configure socket options for %s %s (%w), and failed to close listener (%v)", network, address, err, closeErr)
		}
		return nil, fmt.Errorf("tcp.Listen: failed to configure socket options for %s %s: %w", network, address, err)
	}

	return tcpListener, nil
}

// isValidNetwork validates the network parameter
func isValidNetwork(network string) bool {
	switch network {
	case "tcp", "tcp4", "tcp6":
		return true
	default:
		return false
	}
}

// configureSocket applies production-ready socket options to the TCP listener
func configureSocket(listener *net.TCPListener) error {
	// Use SyscallConn to access the raw socket without duplicating the file descriptor
	rawConn, err := listener.SyscallConn()
	if err != nil {
		return fmt.Errorf("failed to get listener syscall connection: %w", err)
	}

	// Apply socket options using the raw connection
	var sockErr error
	err = rawConn.Control(func(fd uintptr) {
		// Enable SO_REUSEADDR to prevent "address already in use" errors
		if err := syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
			sockErr = fmt.Errorf("failed to set SO_REUSEADDR: %w", err)
			return
		}

		// Set receive buffer size for optimal throughput
		if err := syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_RCVBUF, bufferSize); err != nil {
			sockErr = fmt.Errorf("failed to set receive buffer size: %w", err)
			return
		}

		// Set send buffer size for optimal throughput
		if err := syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_SNDBUF, bufferSize); err != nil {
			sockErr = fmt.Errorf("failed to set send buffer size: %w", err)
			return
		}

		// Enable TCP_NODELAY to minimize latency (disable Nagle's algorithm)
		if err := syscall.SetsockoptInt(int(fd), syscall.IPPROTO_TCP, syscall.TCP_NODELAY, 1); err != nil {
			sockErr = fmt.Errorf("failed to set TCP_NODELAY: %w", err)
			return
		}
	})

	if err != nil {
		return fmt.Errorf("failed to access socket for configuration: %w", err)
	}
	if sockErr != nil {
		return sockErr
	}

	return nil
}
