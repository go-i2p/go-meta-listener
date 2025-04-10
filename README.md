# go-meta-listener

A Go package that implements a unified network listener interface capable of simultaneously handling connections from multiple underlying transport protocols.

[![Go Reference](https://pkg.go.dev/badge/github.com/go-i2p/go-meta-listener.svg)](https://pkg.go.dev/github.com/go-i2p/go-meta-listener)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

`go-meta-listener` provides a "meta listener" implementation that:
- Manages multiple network listeners through a single interface
- Supports any `net.Listener` implementation (TCP, Unix sockets, TLS, etc.)
- Handles connections and errors from all managed listeners
- Enables graceful shutdown across all listeners

The package also includes a specialized `mirror` implementation for multi-protocol network services supporting:
- TLS-secured clearnet connections
- Tor onion services
- I2P garlic services

## Installation

```bash
# Install core package
go get github.com/go-i2p/go-meta-listener

# For multi-protocol mirror functionality
go get github.com/go-i2p/go-meta-listener/mirror
```

## Basic Usage

```go
package main

import (
    "log"
    "net"
    "net/http"

    "github.com/go-i2p/go-meta-listener"
)

func main() {
    // Create a new meta listener
    metaListener := meta.NewMetaListener()
    defer metaListener.Close()

    // Add a TCP listener
    tcpListener, _ := net.Listen("tcp", ":8080")
    metaListener.AddListener("tcp", tcpListener)

    // Add a TLS listener
    tlsListener, _ := tls.Listen("tcp", ":8443", tlsConfig)
    metaListener.AddListener("tls", tlsListener)

    // Use with standard http server
    http.Serve(metaListener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello from any protocol!"))
    }))
}
```

## Mirror Functionality

The `mirror` package provides a simpler interface for creating services available on clearnet, Tor, and I2P simultaneously:

```go
import "github.com/go-i2p/go-meta-listener/mirror"

// Create a multi-protocol listener
listener, err := mirror.Listen(
    "yourdomain.com",        // Domain name for TLS
    "your.email@example.com", // Email for Let's Encrypt
    "./certs",               // Certificate directory
    false                    // Enable/disable TLS on hidden services
)
defer listener.Close()

// Use with standard library
http.Serve(listener, yourHandler)
```

## Examples

See the [example directory](./example) for complete HTTP server examples and the [mirror/metaproxy directory](./mirror/metaproxy) for multi-protocol connection forwarding.

## License

MIT License - Copyright (c) 2025 I2P For Go
