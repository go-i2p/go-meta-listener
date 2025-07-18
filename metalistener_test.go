package meta

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

// mockListener is a test listener that can simulate errors
type mockListener struct {
	addr      net.Addr
	connCh    chan net.Conn
	closeCh   chan struct{}
	closed    bool
	mu        sync.Mutex
	errorMode bool
}

func newMockListener(addr string) *mockListener {
	return &mockListener{
		addr:    &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080},
		connCh:  make(chan net.Conn, 1),
		closeCh: make(chan struct{}),
	}
}

func (m *mockListener) Accept() (net.Conn, error) {
	m.mu.Lock()
	errorMode := m.errorMode
	m.mu.Unlock()

	if errorMode {
		return nil, fmt.Errorf("permanent error")
	}

	select {
	case conn := <-m.connCh:
		return conn, nil
	case <-m.closeCh:
		return nil, fmt.Errorf("listener closed")
	}
}

func (m *mockListener) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.closed {
		return nil
	}

	m.closed = true
	close(m.closeCh)
	return nil
}

func (m *mockListener) Addr() net.Addr {
	return m.addr
}

func (m *mockListener) setErrorMode(errorMode bool) {
	m.mu.Lock()
	m.errorMode = errorMode
	m.mu.Unlock()
}

// TestListenerRemovalRace tests that listener removal doesn't cause race conditions
func TestListenerRemovalRace(t *testing.T) {
	ml := NewMetaListener()
	defer ml.Close()

	// Add multiple listeners
	listener1 := newMockListener("127.0.0.1:8080")
	listener2 := newMockListener("127.0.0.1:8081")

	err := ml.AddListener("test1", listener1)
	if err != nil {
		t.Fatalf("Failed to add listener1: %v", err)
	}

	err = ml.AddListener("test2", listener2)
	if err != nil {
		t.Fatalf("Failed to add listener2: %v", err)
	}

	// Verify both listeners are present
	if ml.Count() != 2 {
		t.Errorf("Expected 2 listeners, got %d", ml.Count())
	}

	// Simulate permanent error in listener1
	listener1.setErrorMode(true)

	// Wait for listener to be removed due to error
	time.Sleep(100 * time.Millisecond)

	// Verify listener1 was removed
	if ml.Count() != 1 {
		t.Errorf("Expected 1 listener after error, got %d", ml.Count())
	}

	// Verify we can still use RemoveListener on the remaining listener
	err = ml.RemoveListener("test2")
	if err != nil {
		t.Errorf("Failed to remove listener2: %v", err)
	}

	if ml.Count() != 0 {
		t.Errorf("Expected 0 listeners after removal, got %d", ml.Count())
	}
}

// TestConcurrentListenerAccess tests concurrent access to listener map
func TestConcurrentListenerAccess(t *testing.T) {
	ml := NewMetaListener()
	defer ml.Close()

	var wg sync.WaitGroup

	// Add multiple listeners concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			listener := newMockListener(fmt.Sprintf("127.0.0.1:%d", 8080+id))
			err := ml.AddListener(fmt.Sprintf("test%d", id), listener)
			if err != nil {
				t.Errorf("Failed to add listener%d: %v", id, err)
			}
		}(i)
	}

	// Concurrently check listener count
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ml.Count()
			ml.ListenerIDs()
		}()
	}

	wg.Wait()

	if ml.Count() != 10 {
		t.Errorf("Expected 10 listeners, got %d", ml.Count())
	}
}

// TestAcceptRaceCondition tests that Accept() method doesn't have race conditions
func TestAcceptRaceCondition(t *testing.T) {
	ml := NewMetaListener()

	// Add a listener
	listener := newMockListener("127.0.0.1:8080")
	err := ml.AddListener("test", listener)
	if err != nil {
		t.Fatalf("Failed to add listener: %v", err)
	}

	var wg sync.WaitGroup

	// Start multiple goroutines calling Accept()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := ml.Accept()
			// We expect either a valid connection or ErrListenerClosed
			if err != nil && err.Error() != ErrListenerClosed.Error() {
				t.Errorf("Unexpected error from Accept(): %v", err)
			}
		}()
	}

	// Concurrently close the listener
	go func() {
		time.Sleep(10 * time.Millisecond)
		ml.Close()
	}()

	wg.Wait()
}
