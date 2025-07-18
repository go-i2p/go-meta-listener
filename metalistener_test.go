package meta

import (
	"context"
	"fmt"
	"io"
	"net"
	"strings"
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

// TestShutdownRaceCondition tests that shutdown doesn't cause race conditions
func TestShutdownRaceCondition(t *testing.T) {
	ml := NewMetaListener()

	// Add multiple listeners
	for i := 0; i < 5; i++ {
		listener := newMockListener(fmt.Sprintf("127.0.0.1:%d", 8080+i))
		err := ml.AddListener(fmt.Sprintf("test%d", i), listener)
		if err != nil {
			t.Fatalf("Failed to add listener%d: %v", i, err)
		}
	}

	var wg sync.WaitGroup

	// Start multiple goroutines that will try to accept connections
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			conn, err := ml.Accept()
			if err != nil {
				// During shutdown, we expect ErrListenerClosed
				if err.Error() != ErrListenerClosed.Error() {
					t.Errorf("Goroutine %d: unexpected error: %v", id, err)
				}
				return
			}
			if conn != nil {
				conn.Close()
			}
		}(i)
	}

	// Allow some time for goroutines to start
	time.Sleep(10 * time.Millisecond)

	// Close the listener
	err := ml.Close()
	if err != nil {
		t.Errorf("Error closing MetaListener: %v", err)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Verify the listener is properly closed
	if ml.Count() != 0 {
		t.Errorf("Expected 0 listeners after close, got %d", ml.Count())
	}
}

// TestWaitGroupSynchronization tests that WaitGroup is properly synchronized
func TestWaitGroupSynchronization(t *testing.T) {
	ml := NewMetaListener()
	defer ml.Close()

	// Test 1: Verify panic recovery in handleListener doesn't break WaitGroup
	listener := newMockListener("127.0.0.1:8080")

	// Add a listener
	err := ml.AddListener("panic-test", listener)
	if err != nil {
		t.Fatalf("Failed to add listener: %v", err)
	}

	// Test 2: Verify we can't add listeners during shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Start WaitForShutdown in background
	go func() {
		ml.WaitForShutdown(ctx)
	}()

	// Give WaitForShutdown a moment to set the shutdown flag
	time.Sleep(10 * time.Millisecond)

	// Try to add another listener - should fail
	listener2 := newMockListener("127.0.0.1:8081")
	err = ml.AddListener("shutdown-test", listener2)
	if err == nil {
		t.Error("Expected error when adding listener during shutdown, got nil")
	}
	if !strings.Contains(err.Error(), "shutdown") {
		t.Errorf("Expected shutdown error, got: %v", err)
	}

	// Close the MetaListener to finish the test cleanly
	ml.Close()
}

// TestWaitGroupPanicRecovery tests that panics in handleListener are recovered
func TestWaitGroupPanicRecovery(t *testing.T) {
	ml := NewMetaListener()
	defer ml.Close()

	// Create a custom listener that will cause a panic
	panicListener := &panicMockListener{
		mockListener: newMockListener("127.0.0.1:8080"),
	}

	// Add the panic-inducing listener
	err := ml.AddListener("panic-test", panicListener)
	if err != nil {
		t.Fatalf("Failed to add listener: %v", err)
	}

	// Trigger the panic by sending a connection
	conn := &mockConn{}
	panicListener.connCh <- conn

	// Wait a bit to let the panic occur and be recovered
	time.Sleep(100 * time.Millisecond)

	// The MetaListener should still be functional
	// Add another listener to verify WaitGroup is not broken
	normalListener := newMockListener("127.0.0.1:8081")
	err = ml.AddListener("normal-test", normalListener)
	if err != nil {
		t.Fatalf("Failed to add normal listener after panic: %v", err)
	}

	// Verify we can still close cleanly
	err = ml.Close()
	if err != nil {
		t.Errorf("Error closing MetaListener after panic: %v", err)
	}
}

// panicMockListener is a mock listener that panics when Accept is called
type panicMockListener struct {
	*mockListener
}

func (p *panicMockListener) Accept() (net.Conn, error) {
	// First call the normal Accept to get a connection
	_, err := p.mockListener.Accept()
	if err != nil {
		return nil, err
	}

	// Then panic to test panic recovery
	panic("test panic in handleListener")
}

// mockConn is a minimal implementation of net.Conn for testing
type mockConn struct{}

func (m *mockConn) Read(b []byte) (n int, err error)  { return 0, io.EOF }
func (m *mockConn) Write(b []byte) (n int, err error) { return len(b), nil }
func (m *mockConn) Close() error                      { return nil }
func (m *mockConn) LocalAddr() net.Addr {
	return &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080}
}

func (m *mockConn) RemoteAddr() net.Addr {
	return &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8081}
}
func (m *mockConn) SetDeadline(t time.Time) error      { return nil }
func (m *mockConn) SetReadDeadline(t time.Time) error  { return nil }
func (m *mockConn) SetWriteDeadline(t time.Time) error { return nil }
