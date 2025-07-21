package mirror

import (
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

// TestFixedConcurrentListenDataRace tests that the data race fix works
func TestFixedConcurrentListenDataRace(t *testing.T) {
	// Disable actual network listeners to focus on testing the map race condition
	os.Setenv("DISABLE_TOR", "true")
	os.Setenv("DISABLE_I2P", "true")
	defer func() {
		os.Unsetenv("DISABLE_TOR")
		os.Unsetenv("DISABLE_I2P")
	}()

	// Create a Mirror instance
	mirror, err := NewMirror("test-fixed:3003")
	if err != nil {
		t.Fatalf("Failed to create mirror: %v", err)
	}
	defer mirror.Close()

	const numGoroutines = 10
	var wg sync.WaitGroup
	errorCh := make(chan error, numGoroutines)

	// Launch multiple goroutines that call Listen() concurrently
	// With the fix, this should NOT trigger data races
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			port := 5000 + id
			portstr := strconv.Itoa(port)
			// Use different addresses to avoid network conflicts
			addr := "test-fixed-" + string(rune('a'+id)) + ":" + portstr

			// This should now be safe thanks to mutex protection
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
		t.Log("Test passed - no data races detected with fix")
	case <-time.After(30 * time.Second):
		t.Fatal("Test timed out - possible deadlock or hang")
	}
}

// TestSequentialOperationsStillWork verifies sequential operations work after the fix
func TestSequentialOperationsStillWork(t *testing.T) {
	// Disable actual network listeners
	os.Setenv("DISABLE_TOR", "true")
	os.Setenv("DISABLE_I2P", "true")
	defer func() {
		os.Unsetenv("DISABLE_TOR")
		os.Unsetenv("DISABLE_I2P")
	}()

	mirror, err := NewMirror("test-sequential:3004")
	if err != nil {
		t.Fatalf("Failed to create mirror: %v", err)
	}
	defer mirror.Close()

	// Sequential calls should work fine
	listener1, err := mirror.Listen("test-seq-1:3004", "")
	if err != nil {
		t.Fatalf("First Listen() failed: %v", err)
	}
	if listener1 != nil {
		listener1.Close()
	}

	listener2, err := mirror.Listen("test-seq-2:3004", "")
	if err != nil {
		t.Fatalf("Second Listen() failed: %v", err)
	}
	if listener2 != nil {
		listener2.Close()
	}

	t.Log("Sequential operations work correctly after fix")
}
