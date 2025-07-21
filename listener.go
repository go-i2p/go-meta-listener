package meta

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

// Accept implements the net.Listener Accept method.
// It returns the next connection from any of the managed listeners.
func (ml *MetaListener) Accept() (net.Conn, error) {
	// Check if already closed before entering the select loop
	if atomic.LoadInt64(&ml.isClosed) != 0 {
		return nil, ErrListenerClosed
	}

	return ml.waitForConnection()
}

// waitForConnection waits for the next available connection from any managed listener.
func (ml *MetaListener) waitForConnection() (net.Conn, error) {
	for {
		select {
		case result, ok := <-ml.connCh:
			if !ok {
				return nil, ErrListenerClosed
			}
			// Access RemoteAddr() directly on the connection
			return result, nil
		case <-ml.closeCh:
			// Double-check the closed state using atomic operation
			if atomic.LoadInt64(&ml.isClosed) != 0 {
				return nil, ErrListenerClosed
			}
			continue
		}
	}
}

// Close implements the net.Listener Close method.
// It closes all managed listeners and releases resources.
func (ml *MetaListener) Close() error {
	// Use atomic compare-and-swap to ensure we only close once
	if !atomic.CompareAndSwapInt64(&ml.isClosed, 0, 1) {
		return nil
	}

	ml.mu.Lock()
	log.Printf("Closing MetaListener with %d listeners", len(ml.listeners))

	// Signal all goroutines to stop first, before clearing listeners map
	close(ml.closeCh)

	// Close all listeners and collect any errors
	errs := ml.closeAllListeners()

	// Clear the listeners map since they're all closed
	ml.listeners = make(map[string]net.Listener)
	ml.mu.Unlock()

	// Wait for all listener goroutines to exit gracefully
	ml.waitForListenerShutdown()

	return ml.formatCloseErrors(errs)
}

// closeAllListeners closes all managed listeners and returns any errors encountered.
func (ml *MetaListener) closeAllListeners() []error {
	var errs []error
	for id, listener := range ml.listeners {
		if err := listener.Close(); err != nil {
			log.Printf("Error closing %s listener: %v", id, err)
			errs = append(errs, err)
		}
	}
	return errs
}

// waitForListenerShutdown waits for all listener goroutines to exit with a grace period.
func (ml *MetaListener) waitForListenerShutdown() {
	gracePeriod := 100 * time.Millisecond
	done := make(chan struct{})

	go func() {
		ml.listenerWg.Wait()
		close(done)
	}()

	// Wait for either all goroutines to finish or grace period to expire
	select {
	case <-done:
		log.Printf("All listener goroutines exited gracefully")
	case <-time.After(gracePeriod):
		log.Printf("Grace period expired, waiting for remaining goroutines")
		// Wait for all listener goroutines to exit
		ml.listenerWg.Wait()
	}
	log.Printf("All listener goroutines have exited")
}

// formatCloseErrors combines multiple close errors into a single error or returns nil.
func (ml *MetaListener) formatCloseErrors(errs []error) error {
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
