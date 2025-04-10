package meta

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	// ErrListenerClosed is returned when attempting to accept on a closed listener
	ErrListenerClosed = errors.New("listener is closed")
	// ErrNoListeners is returned when the meta listener has no active listeners
	ErrNoListeners = errors.New("no active listeners")
)

// MetaListener implements the net.Listener interface and manages multiple
// underlying network listeners as a unified interface.
type MetaListener struct {
	// listeners is a map of registered listeners with their unique identifiers
	listeners map[string]net.Listener
	// listenerWg tracks active listener goroutines for graceful shutdown
	listenerWg sync.WaitGroup
	// connCh is used to receive connections from all managed listeners
	connCh chan ConnResult
	// errCh is used to receive errors from all managed listeners
	errCh chan error
	// closeCh signals all goroutines to stop
	closeCh chan struct{}
	// isClosed indicates whether the meta listener has been closed
	isClosed bool
	// mu protects concurrent access to the listener's state
	mu sync.RWMutex
}

// ConnResult represents a connection received from a listener
type ConnResult struct {
	net.Conn
	src string // source listener ID
}

// NewMetaListener creates a new MetaListener instance ready to manage multiple listeners.
func NewMetaListener() *MetaListener {
	return &MetaListener{
		listeners: make(map[string]net.Listener),
		connCh:    make(chan ConnResult),
		errCh:     make(chan error, 1), // Buffered to prevent blocking
		closeCh:   make(chan struct{}),
	}
}

// AddListener adds a new listener with the specified ID.
// Returns an error if a listener with the same ID already exists or if the
// provided listener is nil.
func (ml *MetaListener) AddListener(id string, listener net.Listener) error {
	if listener == nil {
		return errors.New("cannot add nil listener")
	}

	ml.mu.Lock()
	defer ml.mu.Unlock()

	if ml.isClosed {
		return ErrListenerClosed
	}

	if _, exists := ml.listeners[id]; exists {
		return fmt.Errorf("listener with ID '%s' already exists", id)
	}

	ml.listeners[id] = listener

	// Start a goroutine to handle connections from this listener
	ml.listenerWg.Add(1)
	go ml.handleListener(id, listener)

	return nil
}

// RemoveListener stops and removes the listener with the specified ID.
// Returns an error if no listener with that ID exists.
func (ml *MetaListener) RemoveListener(id string) error {
	ml.mu.Lock()
	defer ml.mu.Unlock()

	listener, exists := ml.listeners[id]
	if !exists {
		return fmt.Errorf("no listener with ID '%s' exists", id)
	}

	// Close the specific listener
	err := listener.Close()
	delete(ml.listeners, id)

	return err
}

// handleListener runs in a separate goroutine for each added listener
// and forwards accepted connections to the connCh channel.
func (ml *MetaListener) handleListener(id string, listener net.Listener) {
	defer ml.listenerWg.Done()

	for {
		conn, err := listener.Accept()

		select {
		case <-ml.closeCh:
			// Meta listener is being closed, exit
			return
		default:
			// Continue processing
		}

		if err != nil {
			// Check if this is a temporary error
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				// For temporary errors, wait a bit and try again
				time.Sleep(100 * time.Millisecond)
				continue
			}

			// For non-temporary errors, check if listener was closed
			ml.mu.RLock()
			_, stillExists := ml.listeners[id]
			ml.mu.RUnlock()

			if stillExists {
				// Only report error if this listener hasn't been removed
				select {
				case ml.errCh <- fmt.Errorf("listener %s error: %w", id, err):
				default:
					// Don't block if no one is reading errors
				}
			}
			return
		}

		// Send the accepted connection to the connection channel
		select {
		case ml.connCh <- ConnResult{Conn: conn, src: id}:
			// Connection forwarded successfully
		case <-ml.closeCh:
			// If we're closing and got a connection, close it
			conn.Close()
			return
		}
	}
}

// Accept implements the net.Listener Accept method.
// It waits for and returns the next connection from any of the managed listeners.
func (ml *MetaListener) Accept() (net.Conn, error) {
	ml.mu.RLock()
	if ml.isClosed {
		ml.mu.RUnlock()
		return nil, ErrListenerClosed
	}

	if len(ml.listeners) == 0 {
		ml.mu.RUnlock()
		return nil, ErrNoListeners
	}
	ml.mu.RUnlock()

	// Wait for either a connection, an error, or close signal
	select {
	case result := <-ml.connCh:
		return result.Conn, nil
	case err := <-ml.errCh:
		return nil, err
	case <-ml.closeCh:
		return nil, ErrListenerClosed
	}
}

// Close implements the net.Listener Close method.
// It closes all managed listeners and releases resources.
func (ml *MetaListener) Close() error {
	ml.mu.Lock()

	if ml.isClosed {
		ml.mu.Unlock()
		return nil
	}

	ml.isClosed = true

	// Signal all goroutines to stop
	close(ml.closeCh)

	// Close all listeners
	var errs []error
	for id, listener := range ml.listeners {
		if err := listener.Close(); err != nil {
			errs = append(errs, fmt.Errorf("failed to close listener %s: %w", id, err))
		}
	}

	ml.mu.Unlock()

	// Wait for all listener goroutines to exit
	ml.listenerWg.Wait()

	// Return combined errors if any
	if len(errs) > 0 {
		return fmt.Errorf("errors closing listeners: %v", errs)
	}

	return nil
}

// Addr implements the net.Listener Addr method.
// It returns a MetaAddr representing all managed listeners.
func (ml *MetaListener) Addr() net.Addr {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	addresses := make([]net.Addr, 0, len(ml.listeners))
	for _, listener := range ml.listeners {
		addresses = append(addresses, listener.Addr())
	}

	return &MetaAddr{addresses: addresses}
}

// ListenerIDs returns the IDs of all active listeners.
func (ml *MetaListener) ListenerIDs() []string {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	ids := make([]string, 0, len(ml.listeners))
	for id := range ml.listeners {
		ids = append(ids, id)
	}

	return ids
}

// Count returns the number of active listeners.
func (ml *MetaListener) Count() int {
	ml.mu.RLock()
	defer ml.mu.RUnlock()

	return len(ml.listeners)
}

// WaitForShutdown blocks until all listener goroutines have exited.
// This is useful for ensuring clean shutdown in server applications.
func (ml *MetaListener) WaitForShutdown(ctx context.Context) error {
	done := make(chan struct{})

	go func() {
		ml.listenerWg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// MetaAddr implements the net.Addr interface for a meta listener.
type MetaAddr struct {
	addresses []net.Addr
}

// Network returns the name of the network.
func (ma *MetaAddr) Network() string {
	return "meta"
}

// String returns a string representation of all managed addresses.
func (ma *MetaAddr) String() string {
	if len(ma.addresses) == 0 {
		return "meta(empty)"
	}

	result := "meta("
	for i, addr := range ma.addresses {
		if i > 0 {
			result += ", "
		}
		result += addr.String()
	}
	result += ")"

	return result
}
