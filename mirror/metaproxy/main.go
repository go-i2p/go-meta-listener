package main

import (
	"flag"
	"fmt"
	"io"
	"net"

	"github.com/go-i2p/go-meta-listener/mirror"
)

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
	flag.Parse()
	addr := net.JoinHostPort(*domain, fmt.Sprintf("%d", *listenPort))
	// Create a new meta listener
	metaListener, err := mirror.Listen(addr, *email, *certDir, *hiddenTls)
	if err != nil {
		panic(err)
	}
	defer metaListener.Close()
	// forward all connections recieved on the meta listener to a local host:port
	for {
		conn, err := metaListener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			localConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
			if err != nil {
				panic(err)
			}
			defer localConn.Close()
			go io.Copy(localConn, conn)
			io.Copy(conn, localConn)
		}()
	}
}
