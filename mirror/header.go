package mirror

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"time"
)

// AddHeaders adds headers to the connection.
// It takes a net.Conn and a map of headers as input.
// It only adds headers if the connection is an HTTP connection.
// It returns a net.Conn with the headers added.
func AddHeaders(conn net.Conn, headers map[string]string) net.Conn {
	// Create a buffer to store the original request
	var buf bytes.Buffer
	teeReader := io.TeeReader(conn, &buf)

	// Try to read the request, but also save it to our buffer
	req, err := http.ReadRequest(bufio.NewReader(teeReader))
	if err != nil {
		// Not an HTTP request or couldn't parse, return original connection
		return conn
	}

	// Add our headers
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Create a pipe to connect our modified request with the output
	pr, pw := io.Pipe()

	// Write the modified request to one end of the pipe with timeout protection
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("PANIC in header processing goroutine: %v", r)
			}
			pw.Close()
		}()

		// Set a deadline for the entire operation to prevent indefinite blocking
		if deadline, ok := conn.(interface{ SetDeadline(time.Time) error }); ok {
			deadline.SetDeadline(time.Now().Add(30 * time.Second))
		}

		// Write the modified request
		if err := req.Write(pw); err != nil {
			log.Printf("Error writing modified request: %v", err)
			return
		}

		// Copy remaining data with timeout protection
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		done := make(chan error, 1)
		go func() {
			_, err := io.Copy(pw, conn)
			done <- err
		}()

		select {
		case err := <-done:
			if err != nil {
				log.Printf("Error copying connection data: %v", err)
			}
		case <-ctx.Done():
			log.Printf("Header processing goroutine timed out after 30 seconds")
			// Connection will be closed by defer pw.Close()
		}
	}()

	// Return a ReadWriter that reads from our pipe and writes to the original connection
	return &readWriteConn{
		Reader: pr,
		Writer: conn,
		conn:   conn,
	}
}

// readWriteConn implements net.Conn
type readWriteConn struct {
	io.Reader
	io.Writer
	conn net.Conn
}

// Implement the rest of net.Conn interface by delegating to the original connection
func (rwc *readWriteConn) Close() error                       { return rwc.conn.Close() }
func (rwc *readWriteConn) LocalAddr() net.Addr                { return rwc.conn.LocalAddr() }
func (rwc *readWriteConn) RemoteAddr() net.Addr               { return rwc.conn.RemoteAddr() }
func (rwc *readWriteConn) SetDeadline(t time.Time) error      { return rwc.conn.SetDeadline(t) }
func (rwc *readWriteConn) SetReadDeadline(t time.Time) error  { return rwc.conn.SetReadDeadline(t) }
func (rwc *readWriteConn) SetWriteDeadline(t time.Time) error { return rwc.conn.SetWriteDeadline(t) }

// Accept accepts a connection from the listener.
// It takes a net.Listener as input and returns a net.Conn with the headers added.
// It is used to accept connections from the meta listener and add headers to them.
func (ml *Mirror) Accept() (net.Conn, error) {
	// Accept a connection from the listener
	conn, err := ml.MetaListener.Accept()
	if err != nil {
		log.Println("Error accepting connection:", err)
		return nil, err
	}

	// Check if the connection is a TLS connection
	if tlsConn, ok := conn.(*tls.Conn); ok {
		// If it is a TLS connection, perform the handshake
		if err := tlsConn.Handshake(); err != nil {
			log.Println("Error performing TLS handshake:", err)
			return nil, err
		}
		// If the handshake is successful, get the underlying connection
		// conn = tlsConn.NetConn()
	}

	host := map[string]string{
		"Host":              conn.LocalAddr().String(),
		"X-Forwarded-For":   conn.RemoteAddr().String(),
		"X-Forwarded-Proto": "http",
	}

	// Add headers to the connection
	conn = AddHeaders(conn, host)

	return conn, nil
}
