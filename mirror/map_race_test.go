package mirror

import (
	"sync"
	"testing"
	"time"

	"github.com/go-i2p/onramp"
)

// TestConcurrentMapAccessDataRace tests for data race in concurrent map access
func TestConcurrentMapAccessDataRace(t *testing.T) {
	// This test directly reproduces the data race on the maps
	// Create a Mirror with empty maps like in Listen()
	mirror := &Mirror{
		Onions:  make(map[string]*onramp.Onion),
		Garlics: make(map[string]*onramp.Garlic),
	}

	const numGoroutines = 20
	var wg sync.WaitGroup

	// Simulate the exact race condition from Listen() method
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Use the same port to force the race condition
			port := "8080"

			// This replicates the exact problematic code from Listen():
			// Simulate the read-check-write pattern from Listen()
			if mirror.Onions[port] == nil && !DisableTor() {
				onion, _ := onramp.NewOnion("test-onion")
				mirror.Onions[port] = onion // Concurrent write to map
			}

			if mirror.Garlics[port] == nil && !DisableI2P() {
				garlic, _ := onramp.NewGarlic("test-garlic", "127.0.0.1:7656", onramp.OPT_WIDE)
				mirror.Garlics[port] = garlic // Concurrent write to map
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
		// Test completed - race detector should catch any issues
	case <-time.After(10 * time.Second):
		t.Fatal("Test timed out")
	}
}
