package mirror

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

// TestConcurrentListenDataRace tests for data race in concurrent Listen() calls
func TestConcurrentListenDataRace(t *testing.T) {
	// Enable race detector with go test -race
	// This test reproduces the data race in Mirror.Listen()

	mirror, err := NewMirror("test-mirror:3001")
	if err != nil {
		t.Fatalf("Failed to create mirror: %v", err)
	}
	defer mirror.Close()

	const numGoroutines = 10
	var wg sync.WaitGroup
	errorCh := make(chan error, numGoroutines)
	// Launch multiple goroutines that call Listen() concurrently
	// This should trigger the data race on ml.Onions and ml.Garlics maps
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Use different ports to avoid "address already in use" errors
			// The race condition occurs on the map access, not the network binding
			port := 4000 + id
			portstr := strconv.Itoa(port)
			addr := "test-" + string(rune('a'+id)) + ":" + portstr

			// This should cause concurrent map access on ml.Onions and ml.Garlics
			listener, err := mirror.Listen(addr, "")
			if err != nil {
				errorCh <- err
				return
			}
			if listener != nil {
				// Clean close if successful
				listener.Close()
			}
		}(i)
	}

	// Wait for all goroutines to complete
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	// Wait with timeout
	select {
	case <-done:
		// Check for any errors
		close(errorCh)
		for err := range errorCh {
			if err != nil {
				t.Errorf("Error in concurrent Listen(): %v", err)
			}
		}
	case <-time.After(30 * time.Second):
		t.Fatal("Test timed out - possible deadlock or hang")
	}
}

// TestSequentialListenWorks verifies that sequential Listen calls work correctly
func TestSequentialListenWorks(t *testing.T) {
	mirror, err := NewMirror("test-sequential:3002")
	if err != nil {
		t.Fatalf("Failed to create mirror: %v", err)
	}
	defer mirror.Close()

	// Sequential calls should work fine with different ports
	listener1, err := mirror.Listen("test-seq-1:3003", "")
	if err != nil {
		t.Fatalf("First Listen() failed: %v", err)
	}
	if listener1 == nil {
		t.Fatal("First Listen() returned nil listener")
	}

	listener2, err := mirror.Listen("test-seq-2:3004", "")
	if err != nil {
		t.Fatalf("Second Listen() failed: %v", err)
	}
	if listener2 == nil {
		t.Fatal("Second Listen() returned nil listener")
	}
}
