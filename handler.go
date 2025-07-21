package meta

import (
	"net"
	"strings"
	"sync/atomic"
	"time"
)

// handleListener runs in a separate goroutine for each added listener
// and forwards accepted connections to the connCh channel.
func (ml *MetaListener) handleListener(id string, listener net.Listener) {
	defer ml.recoverAndCleanup(id)

	for {
		if ml.shouldStopListener(id) {
			return
		}

		ml.setAcceptDeadline(listener)

		conn, err := listener.Accept()
		if err != nil {
			if ml.handleAcceptError(id, err) {
				continue
			}
			return
		}

		log.Printf("Listener %s accepted connection from %s", id, conn.RemoteAddr())
		ml.forwardConnection(id, conn)
	}
}

// recoverAndCleanup handles panic recovery and ensures proper cleanup for listener goroutines.
func (ml *MetaListener) recoverAndCleanup(id string) {
	if r := recover(); r != nil {
		log.Printf("PANIC in listener goroutine for %s: %v", id, r)
	}
	log.Printf("Listener goroutine for %s exiting", id)
	ml.listenerWg.Done()
}

// shouldStopListener checks if the MetaListener is closed and should stop processing.
func (ml *MetaListener) shouldStopListener(id string) bool {
	select {
	case <-ml.closeCh:
		log.Printf("MetaListener closed, stopping %s listener", id)
		return true
	default:
		return false
	}
}

// setAcceptDeadline sets a deadline for Accept to prevent blocking indefinitely.
func (ml *MetaListener) setAcceptDeadline(listener net.Listener) {
	if deadline, ok := listener.(interface{ SetDeadline(time.Time) error }); ok {
		deadline.SetDeadline(time.Now().Add(1 * time.Second))
	}
}

// handleAcceptError processes errors from listener.Accept() and determines if processing should continue.
// Returns true if the listener should continue processing, false if it should stop.
func (ml *MetaListener) handleAcceptError(id string, err error) bool {
	// Check if this is a timeout error (which we expect due to our deadline)
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}

	// For other network errors, check if it's retryable based on common patterns
	errStr := err.Error()
	if strings.Contains(errStr, "connection reset") || strings.Contains(errStr, "broken pipe") ||
		strings.Contains(errStr, "resource temporarily unavailable") {
		log.Printf("Retryable error in %s listener: %v, retrying in 100ms", id, err)
		time.Sleep(100 * time.Millisecond)
		return true
	}

	// Check if the listener was closed (expected during shutdown)
	if atomic.LoadInt64(&ml.isClosed) != 0 {
		log.Printf("Listener %s closed during shutdown", id)
		return false
	}

	log.Printf("Permanent error in %s listener: %v, stopping", id, err)
	ml.signalListenerRemoval(id)
	return false
}

// signalListenerRemoval attempts to signal that a listener should be removed.
func (ml *MetaListener) signalListenerRemoval(id string) {
	select {
	case ml.removeListenerCh <- id:
		// Successfully signaled for removal
	case <-ml.closeCh:
		// MetaListener is closing, no need to signal removal
	}
}

// forwardConnection attempts to forward a connection through the connection channel.
func (ml *MetaListener) forwardConnection(id string, conn net.Conn) {
	select {
	case ml.connCh <- ConnResult{Conn: conn, src: id}:
		log.Printf("Connection from %s successfully forwarded via %s", conn.RemoteAddr(), id)
	case <-ml.closeCh:
		log.Printf("MetaListener closing while forwarding connection, closing connection")
		conn.Close()
	case <-time.After(5 * time.Second):
		// If we can't forward within 5 seconds, something is seriously wrong
		log.Printf("WARNING: Connection forwarding timed out, closing connection from %s", conn.RemoteAddr())
		conn.Close()
	}
}
