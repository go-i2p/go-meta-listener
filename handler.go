package meta

import (
	"net"
	"time"
)

// handleListener runs in a separate goroutine for each added listener
// and forwards accepted connections to the connCh channel.
func (ml *MetaListener) handleListener(id string, listener net.Listener) {
	defer func() {
		log.Printf("Listener goroutine for %s exiting", id)
		ml.listenerWg.Done()
	}()

	for {
		// First check if the MetaListener is closed
		select {
		case <-ml.closeCh:
			log.Printf("MetaListener closed, stopping %s listener", id)
			return
		default:
		}

		// Set a deadline for Accept to prevent blocking indefinitely
		if deadline, ok := listener.(interface{ SetDeadline(time.Time) error }); ok {
			deadline.SetDeadline(time.Now().Add(1 * time.Second))
		}

		conn, err := listener.Accept()
		if err != nil {
			// Check if this is a timeout error (which we expect due to our deadline)
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}

			// Check if this is any other temporary error
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				log.Printf("Temporary error in %s listener: %v, retrying in 100ms", id, err)
				time.Sleep(100 * time.Millisecond)
				continue
			}

			log.Printf("Permanent error in %s listener: %v, stopping", id, err)
			select {
			case ml.removeListenerCh <- id:
				// Successfully signaled for removal
			case <-ml.closeCh:
				// MetaListener is closing, no need to signal removal
			}
			return
		}

		// If we reach here, we have a valid connection
		log.Printf("Listener %s accepted connection from %s", id, conn.RemoteAddr())

		// Try to forward the connection, but don't block indefinitely
		select {
		case ml.connCh <- ConnResult{Conn: conn, src: id}:
			log.Printf("Connection from %s successfully forwarded via %s", conn.RemoteAddr(), id)
		case <-ml.closeCh:
			log.Printf("MetaListener closing while forwarding connection, closing connection")
			conn.Close()
			return
		case <-time.After(5 * time.Second):
			// If we can't forward within 5 seconds, something is seriously wrong
			log.Printf("WARNING: Connection forwarding timed out, closing connection from %s", conn.RemoteAddr())
			conn.Close()
		}
	}
}
