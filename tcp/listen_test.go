package tcp

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestListen_ValidNetworks(t *testing.T) {
	testCases := []struct {
		network string
		address string
	}{
		{"tcp", ":0"},
		{"tcp4", ":0"},
		{"tcp6", ":0"},
		{"tcp", "127.0.0.1:0"},
		{"tcp4", "127.0.0.1:0"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%s", tc.network, tc.address), func(t *testing.T) {
			listener, err := Listen(tc.network, tc.address)
			if err != nil {
				t.Fatalf("Listen(%s, %s) failed: %v", tc.network, tc.address, err)
			}
			defer listener.Close()

			// Verify we can get the actual address
			addr := listener.Addr()
			if addr == nil {
				t.Fatal("Listener address is nil")
			}

			// Verify the address has expected network type
			if !strings.HasPrefix(addr.Network(), "tcp") {
				t.Errorf("Expected tcp network, got %s", addr.Network())
			}
		})
	}
}

func TestListen_InvalidNetworks(t *testing.T) {
	invalidNetworks := []string{"udp", "unix", "ip", "", "TCP", "tcp1"}

	for _, network := range invalidNetworks {
		t.Run(network, func(t *testing.T) {
			_, err := Listen(network, ":0")
			if err == nil {
				t.Fatalf("Listen(%s, :0) should have failed", network)
			}

			expectedMsg := "invalid network type"
			if !strings.Contains(err.Error(), expectedMsg) {
				t.Errorf("Error should contain %q, got: %v", expectedMsg, err)
			}
		})
	}
}

func TestListen_InvalidAddresses(t *testing.T) {
	invalidAddresses := []string{
		":::",
		":::8080",
		"256.256.256.256:8080",
		"invalid:address",
		"127.0.0.1:99999",
	}

	for _, addr := range invalidAddresses {
		t.Run(addr, func(t *testing.T) {
			_, err := Listen("tcp", addr)
			if err == nil {
				t.Fatalf("Listen(tcp, %s) should have failed", addr)
			}

			expectedMsg := "invalid address"
			if !strings.Contains(err.Error(), expectedMsg) {
				t.Errorf("Error should contain %q, got: %v", expectedMsg, err)
			}
		})
	}
}

func TestListen_AcceptConnections(t *testing.T) {
	listener, err := Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Test accepting multiple connections
	const numConnections = 5
	var wg sync.WaitGroup

	// Start accepting connections
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numConnections; i++ {
			conn, err := listener.Accept()
			if err != nil {
				t.Errorf("Accept failed: %v", err)
				return
			}
			conn.Close()
		}
	}()

	// Create client connections
	for i := 0; i < numConnections; i++ {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatalf("Failed to dial listener: %v", err)
		}
		conn.Close()
	}

	wg.Wait()
}

func TestListen_RapidRestarts(t *testing.T) {
	// Test that SO_REUSEADDR prevents "address already in use" errors
	address := ":0"

	// Create and close listener to get a port
	initialListener, err := Listen("tcp", address)
	if err != nil {
		t.Fatalf("Failed to create initial listener: %v", err)
	}

	// Get the actual assigned port
	actualAddr := initialListener.Addr().String()
	initialListener.Close()

	// Rapidly restart listeners on the same port
	for i := 0; i < 3; i++ {
		listener, err := Listen("tcp", actualAddr)
		if err != nil {
			t.Fatalf("Restart %d failed: %v", i, err)
		}
		listener.Close()

		// Brief pause to simulate real restart scenario
		time.Sleep(10 * time.Millisecond)
	}
}

func TestListen_ConcurrentAccess(t *testing.T) {
	listener, err := Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()
	const numGoroutines = 10
	var wg sync.WaitGroup

	// Start multiple goroutines that will connect simultaneously
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", addr)
			if err != nil {
				t.Errorf("Goroutine %d: dial failed: %v", id, err)
				return
			}
			defer conn.Close()

			// Write some data to verify connection works
			testData := fmt.Sprintf("test-%d", id)
			if _, err := conn.Write([]byte(testData)); err != nil {
				t.Errorf("Goroutine %d: write failed: %v", id, err)
			}
		}(i)
	}

	// Accept all connections
	acceptedCount := 0
	done := make(chan bool)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			acceptedCount++
			conn.Close()

			if acceptedCount >= numGoroutines {
				done <- true
				return
			}
		}
	}()

	// Wait for all connections with timeout
	select {
	case <-done:
		// Success
	case <-time.After(5 * time.Second):
		t.Fatalf("Timeout waiting for connections, accepted %d/%d", acceptedCount, numGoroutines)
	}

	wg.Wait()
}

func TestListen_ErrorMessages(t *testing.T) {
	testCases := []struct {
		network     string
		address     string
		expectedErr string
	}{
		{"udp", ":8080", "invalid network type"},
		{"tcp", ":::", "invalid address"},
		{"", ":8080", "invalid network type"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%s", tc.network, tc.address), func(t *testing.T) {
			_, err := Listen(tc.network, tc.address)
			if err == nil {
				t.Fatalf("Expected error for Listen(%s, %s)", tc.network, tc.address)
			}

			if !strings.Contains(err.Error(), tc.expectedErr) {
				t.Errorf("Error should contain %q, got: %v", tc.expectedErr, err)
			}

			// Verify error message includes the problematic values
			errMsg := err.Error()
			if tc.network != "" && !strings.Contains(errMsg, tc.network) {
				t.Errorf("Error should include network %q: %v", tc.network, err)
			}
			if tc.address != "" && !strings.Contains(errMsg, tc.address) {
				t.Errorf("Error should include address %q: %v", tc.address, err)
			}
		})
	}
}

// Benchmark the listener creation overhead
func BenchmarkListen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		listener, err := Listen("tcp", ":0")
		if err != nil {
			b.Fatalf("Listen failed: %v", err)
		}
		listener.Close()
	}
}

func TestConfigureSocket_NoResourceLeak(t *testing.T) {
	// Create a listener
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to create base listener: %v", err)
	}
	defer listener.Close()

	tcpListener := listener.(*net.TCPListener)

	// Configure the socket - this should not invalidate the listener
	err = configureSocket(tcpListener)
	if err != nil {
		t.Fatalf("Socket configuration failed: %v", err)
	}

	// Verify the listener is still usable by accepting a connection
	addr := listener.Addr().String()

	// Test connection in a goroutine
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			conn.Close()
		}
	}()

	// This should succeed if the listener wasn't invalidated
	conn, err := listener.Accept()
	if err != nil {
		t.Fatalf("Listener became unusable after configuration: %v", err)
	}
	conn.Close()
}
