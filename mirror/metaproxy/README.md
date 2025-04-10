# metaproxy - Connection Forwarder

A simple utility that forwards connections from a meta listener to a specified host and port. Part of the `go-i2p/go-meta-listener` toolset.
Automatically forwards a local service to a TLS Service, an I2P Eepsite, and a Tor Onion service at the same time.

## Installation

To install the metaproxy utility, use:

```bash
go install github.com/go-i2p/go-meta-listener/mirror/metaproxy@latest
```

## Usage

```bash
metaproxy [options]
```

### Options

- `-host`: Host to forward connections to (default: "localhost")
- `-port`: Port to forward connections to (default: 8080)
- `-domain`: Domain name for TLS listener (default: "i2pgit.org")
- `-email`: Email address for Let's Encrypt registration (default: "example@example.com")
- `-certdir`: Directory for storing certificates (default: "./certs")
- `-hidden-tls`: Enable hidden TLS (default: false)

## Description

metaproxy creates a meta listener that can accept connections from multiple transport types and forwards them to a specified destination (host:port).
It supports TLS with automatic certificate management through Let's Encrypt, I2P EepSites, and Tor Onion Services.

## Examples

Forward connections to a local web server:
```bash
metaproxy -host localhost -port 3000
```

Forward connections with custom TLS settings:
```bash
metaproxy -domain yourdomain.com -email you@example.com -certdir /etc/certs -port 8443
```
