package mirror

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/go-i2p/go-meta-listener"
	"github.com/go-i2p/go-meta-listener/tcp"
	"github.com/go-i2p/onramp"

	wileedot "github.com/opd-ai/wileedot"
)

type Mirror struct {
	*meta.MetaListener
	mu      sync.RWMutex // protects Onions and Garlics maps
	Onions  map[string]*onramp.Onion
	Garlics map[string]*onramp.Garlic
}

var _ net.Listener = &Mirror{}

func (m *Mirror) Close() error {
	log.Println("Closing Mirror")
	if err := m.MetaListener.Close(); err != nil {
		log.Println("Error closing MetaListener:", err)
	} else {
		log.Println("MetaListener closed")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	for _, onion := range m.Onions {
		if err := onion.Close(); err != nil {
			log.Println("Error closing Onion:", err)
		} else {
			log.Println("Onion closed")
		}
	}
	for _, garlic := range m.Garlics {
		if err := garlic.Close(); err != nil {
			log.Println("Error closing Garlic:", err)
		} else {
			log.Println("Garlic closed")
		}
	}

	// Clear the maps to prevent reuse of closed instances
	m.Onions = make(map[string]*onramp.Onion)
	m.Garlics = make(map[string]*onramp.Garlic)

	log.Println("Mirror closed")
	return nil
}

func NewMirror(name string) (*Mirror, error) {
	log.Println("Creating new Mirror")
	inner := meta.NewMetaListener()
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, " ", "")
	if name == "" {
		name = "mirror"
	}
	log.Printf("Creating new MetaListener with name: '%s'\n", name)
	_, port, err := net.SplitHostPort(name)
	if err != nil {
		port = "3000"
	}
	onions := make(map[string]*onramp.Onion)
	if !DisableTor() {
		onion, err := onramp.NewOnion("metalistener-" + name)
		if err != nil {
			return nil, err
		}
		log.Println("Created new Onion manager")
		onions[port] = onion
	}
	garlics := make(map[string]*onramp.Garlic)
	if !DisableI2P() {
		garlic, err := onramp.NewGarlic("metalistener-"+name, "127.0.0.1:7656", onramp.OPT_WIDE)
		if err != nil {
			return nil, err
		}
		log.Println("Created new Garlic manager")
		garlics[port] = garlic
	}

	ml := &Mirror{
		MetaListener: inner,
		Onions:       onions,
		Garlics:      garlics,
	}
	log.Printf("Mirror created with name: '%s' and port: '%s', '%s'\n", name, port, ml.MetaListener.Addr().String())
	return ml, nil
}

func (ml Mirror) Listen(name, addr string) (net.Listener, error) {
	log.Println("Starting Mirror Listener")
	// get the port:
	_, port, err := net.SplitHostPort(name)
	if err != nil {
		// check if host is an IP address
		if net.ParseIP(name) == nil {
			// host = "127.0.0.1"
		}
		port = "3000"
	}
	hiddenTls := hiddenTls(port)
	log.Printf("Actual args: name: '%s' addr: '%s' certDir: '%s' hiddenTls: '%t'\n", name, addr, certDir(), hiddenTls)
	localAddr := net.JoinHostPort("127.0.0.1", port)
	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to create TCP listener on %s: %w", localAddr, err)
	}
	tcpListener := listener.(*net.TCPListener)
	hardenedListener, err := tcp.Config(*tcpListener)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TCP listener created on %s\n", localAddr)
	if err := ml.AddListener(port, hardenedListener); err != nil {
		return nil, err
	}
	log.Printf("HTTP Local listener added http://%s\n", tcpListener.Addr())
	log.Println("Checking for existing onion and garlic listeners")
	listenerId := fmt.Sprintf("metalistener-%s-%s", name, port)
	log.Println("Listener ID:", listenerId)

	// Protect map access with mutex to prevent data race
	ml.mu.Lock()

	// Check if onion and garlic listeners already exist
	if ml.Onions[port] == nil && !DisableTor() {
		// make a new onion listener
		// and add it to the map
		log.Println("Creating new onion listener")
		onion, err := onramp.NewOnion(listenerId)
		if err != nil {
			ml.mu.Unlock()
			return nil, err
		}
		log.Println("Onion listener created for port", port)
		ml.Onions[port] = onion
	}
	if ml.Garlics[port] == nil && !DisableI2P() {
		// make a new garlic listener
		// and add it to the map
		log.Println("Creating new garlic listener")
		garlic, err := onramp.NewGarlic(listenerId, "127.0.0.1:7656", onramp.OPT_WIDE)
		if err != nil {
			ml.mu.Unlock()
			return nil, err
		}
		log.Println("Garlic listener created for port", port)
		ml.Garlics[port] = garlic
	}

	ml.mu.Unlock()
	if hiddenTls {
		// make sure an onion and a garlic listener exist at ml.Onions[port] and ml.Garlics[port]
		// and listen on them, check existence first
		if !DisableTor() {
			ml.mu.RLock()
			onionInstance := ml.Onions[port]
			ml.mu.RUnlock()

			if onionInstance == nil {
				return nil, fmt.Errorf("no onion instance found for port %s", port)
			}
			onionListener, err := onionInstance.ListenTLS()
			if err != nil {
				return nil, err
			}
			oid := fmt.Sprintf("onion-%s", onionListener.Addr().String())
			if err := ml.AddListener(oid, onionListener); err != nil {
				return nil, err
			}
			log.Printf("OnionTLS listener added https://%s\n", onionListener.Addr())
		}
		if !DisableI2P() {
			ml.mu.RLock()
			garlicInstance := ml.Garlics[port]
			ml.mu.RUnlock()

			if garlicInstance == nil {
				return nil, fmt.Errorf("no garlic instance found for port %s", port)
			}
			garlicListener, err := garlicInstance.ListenTLS()
			if err != nil {
				return nil, err
			}
			gid := fmt.Sprintf("garlic-%s", garlicListener.Addr().String())
			if err := ml.AddListener(gid, garlicListener); err != nil {
				return nil, err
			}
			log.Printf("GarlicTLS listener added https://%s\n", garlicListener.Addr())
		}
	} else {
		if !DisableTor() {
			ml.mu.RLock()
			onionInstance := ml.Onions[port]
			ml.mu.RUnlock()

			if onionInstance == nil {
				return nil, fmt.Errorf("no onion instance found for port %s", port)
			}
			onionListener, err := onionInstance.Listen()
			if err != nil {
				return nil, err
			}
			oid := fmt.Sprintf("onion-%s", onionListener.Addr().String())
			if err := ml.AddListener(oid, onionListener); err != nil {
				return nil, err
			}
			log.Printf("Onion listener added http://%s\n", onionListener.Addr())
		}
		if !DisableI2P() {
			ml.mu.RLock()
			garlicInstance := ml.Garlics[port]
			ml.mu.RUnlock()

			if garlicInstance == nil {
				return nil, fmt.Errorf("no garlic instance found for port %s", port)
			}
			garlicListener, err := garlicInstance.Listen()
			if err != nil {
				return nil, err
			}
			gid := fmt.Sprintf("garlic-%s", garlicListener.Addr().String())
			if err := ml.AddListener(gid, garlicListener); err != nil {
				return nil, err
			}
			log.Printf("Garlic listener added http://%s\n", garlicListener.Addr())
		}
	}
	if addr != "" {
		cfg := wileedot.Config{
			Domain:         name,
			AllowedDomains: []string{name},
			CertDir:        certDir(),
			Email:          addr,
		}
		tlsListener, err := wileedot.New(cfg)
		if err != nil {
			return nil, err
		}
		tid := fmt.Sprintf("tls-%s", tlsListener.Addr().String())
		if err := ml.AddListener(tid, tlsListener); err != nil {
			return nil, err
		}
		log.Printf("TLS listener added https://%s\n", tlsListener.Addr())
	}
	return &ml, nil
}

// Listen creates a new Mirror instance and sets up listeners for TLS, Onion, and Garlic.
// It returns the Mirror instance and any error encountered during setup.
// name is the domain name used for the TLS listener, required for Let's Encrypt.
// addr is the email address used for Let's Encrypt registration.
// It is recommended to use a valid email address for production use.
func Listen(name, addr string) (net.Listener, error) {
	ml, err := NewMirror(name)
	if err != nil {
		return nil, err
	}
	return ml.Listen(name, addr)
}

func DisableTor() bool {
	val := os.Getenv("DISABLE_TOR")
	if val == "1" || strings.ToLower(val) == "true" {
		log.Println("Tor is disabled by environment variable DISABLE_TOR")
		return true
	}
	return false
}

func DisableI2P() bool {
	val := os.Getenv("DISABLE_I2P")
	if val == "1" || strings.ToLower(val) == "true" {
		log.Println("I2P is disabled by environment variable DISABLE_I2P")
		return true
	}
	return false
}

// HIDDEN_TLS is a global variable that determines whether to use hidden TLS.
// It is set to true by default, but can be overridden by the hiddenTls function.
// If the port ends with "22", it will return false, indicating that hidden TLS should not be used.
// This is a useful workaround for SSH connections, which commonly use port 22.
var HIDDEN_TLS = true

func hiddenTls(port string) bool {
	// Check if the port is 22, which is commonly used for SSH
	if strings.HasSuffix(port, "22") {
		log.Println("Port ends with 22, setting hiddenTls to false")
		return false
	}
	// Default to true for other ports
	return HIDDEN_TLS
}

var default_CERT_DIR = "./certs"

// CERT_DIR is the directory where certificates are stored.
// It can be overridden by setting the CERT_DIR environment variable.
// if CERT_DIR is not set, it defaults to "./certs".
// if CERT_DIR is set from Go code, it will always return the value set in the code.
// if CERT_DIR is set from the environment, it will return the value from the environment unless overridden by Go code.
var CERT_DIR = default_CERT_DIR

func certDir() string {
	// Default certificate directory
	certDir := CERT_DIR
	if certDir != default_CERT_DIR {
		// if the default directory is not used, always return it
		return certDir
	}
	if dir := os.Getenv("CERT_DIR"); dir != "" {
		certDir = dir
	}
	log.Printf("Using certificate directory: %s\n", certDir)
	return certDir
}
