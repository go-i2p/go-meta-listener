package meta

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/samber/oops"
)

var (
	// ErrListenerClosed is returned when attempting to accept on a closed listener
	ErrListenerClosed = oops.Errorf("listener is closed")
	// ErrNoListeners is returned when the meta listener has no active listeners
	ErrNoListeners = oops.Errorf("no active listeners")
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
		connCh:    make(chan ConnResult, 100), // Larger buffer for high connection volume
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
