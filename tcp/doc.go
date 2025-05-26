// Package tcp provides a minimal, production-hardened TCP listener implementation
// with secure cross-platform defaults suitable for internet-facing services.
//
// This package exports a single function Listen() that creates TCP listeners
// with optimized settings for production workloads including connection reuse,
// keep-alive monitoring, optimized buffer sizes, and minimal latency configuration.
//
// The implementation uses only Go standard library components to ensure
// compatibility across Windows, Linux, macOS, and BSD systems without
// requiring platform-specific code or external dependencies.
//
// Example usage:
//
//	listener, err := tcp.Listen("tcp", ":8080")
//	if err != nil {
//		log.Fatalf("Failed to create listener: %v", err)
//	}
//	defer listener.Close()
//
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			log.Printf("Accept error: %v", err)
//			continue
//		}
//		go handleConnection(conn)
//	}
package tcp
