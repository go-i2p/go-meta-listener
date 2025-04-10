# Mirror Listener

A network listener implementation that simultaneously listens on clearnet (TLS), Tor onion services, and I2P garlic services.

## Overview

Mirror Listener is a wrapper around the [go-meta-listener](https://github.com/go-i2p/go-meta-listener) package that provides a simplified interface for setting up multi-protocol listeners. It automatically configures:

- TLS-secured clearnet connections with Let's Encrypt certificates
- Tor onion service endpoints
- I2P garlic service endpoints

This allows you to run a single service that's accessible through multiple network layers and protocols.

## Installation

```bash
go get github.com/go-i2p/go-meta-listener/mirror
```

## Usage

```go
import (
    "github.com/go-i2p/go-meta-listener/mirror"
    "net/http"
)

func main() {
    // Create a multi-protocol listener
    listener, err := mirror.Listen(
        "yourdomain.com",        // Domain name for TLS
        "your.email@example.com", // Email for Let's Encrypt
        "./certs",               // Certificate directory
        false,                   // Enable/disable TLS on hidden services
    )
    if err != nil {
        panic(err)
    }
    defer listener.Close()
    
    // Use with standard library
    http.Serve(listener, yourHandler)
}
```

## Configuration Options

- **Domain Name**: Required for TLS certificate issuance through Let's Encrypt
- **Email Address**: Used for Let's Encrypt registration
- **Certificate Directory**: Where TLS certificates will be stored
- **Hidden TLS**: When set to true, enables TLS for Tor and I2P services as well

## Example: Connection Forwarding

See the [example directory](./example) for a complete example of using Mirror Listener to forward connections to a local service.