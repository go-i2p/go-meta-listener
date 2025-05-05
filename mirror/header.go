package mirror

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
	"net/http"
)

// AddHeaders adds headers to the connection.
// It takes a net.Conn and a map of headers as input.
// It only adds headers if the connection is an HTTP connection.
// It returns a net.Conn with the headers added.
func AddHeaders(conn net.Conn, headers map[string]string) net.Conn {
	// read a request from the connection
	// if the request is an HTTP request, add the headers
	// if the request is not an HTTP request, return the connection as is
	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		log.Println("Error reading request:", err)
		// if the request is not an HTTP request, return the connection as is
		return conn
	}
	log.Println("Adding headers to connection:", req.Method, req.URL)
	for key, value := range headers {
		req.Header.Add(key, value)
		log.Println("Added header:", key, value)
	}
	// write the request back to the connection
	if err := req.Write(conn); err != nil {
		log.Println("Error writing request:", err)
		// if there is an error writing the request, return the connection as is
		return conn
	}
	// If all goes well, return the connection with the headers added
	return conn
}

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
		conn = tlsConn.NetConn()
	}

	host := map[string]string{
		"Host":              ml.MetaListener.Addr().String(),
		"X-Forwarded-For":   conn.RemoteAddr().String(),
		"X-Forwarded-Proto": "http",
	}

	// Add headers to the connection
	conn = AddHeaders(conn, host)

	return conn, nil
}
