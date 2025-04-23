package mirror

import (
	"log"
	"net"
	"strings"

	"github.com/go-i2p/go-meta-listener"
	"github.com/go-i2p/onramp"

	wileedot "github.com/opd-ai/wileedot"
)

type Mirror struct {
	*meta.MetaListener
	Onions  map[string]*onramp.Onion
	Garlics map[string]*onramp.Garlic
}

var _ net.Listener = &Mirror{}

func (m *Mirror) Close() error {
	if err := m.MetaListener.Close(); err != nil {
		log.Println("Error closing MetaListener:", err)
	}
	for _, onion := range m.Onions {
		if err := onion.Close(); err != nil {
			log.Println("Error closing Onion:", err)
		}
	}
	for _, garlic := range m.Garlics {
		if err := garlic.Close(); err != nil {
			log.Println("Error closing Garlic:", err)
		}
	}
	return nil
}

func NewMirror(name string) (*Mirror, error) {
	inner := meta.NewMetaListener()
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, " ", "")
	if name == "" {
		name = "mirror"
	}
	onion, err := onramp.NewOnion("metalistener-" + name)
	if err != nil {
		return nil, err
	}
	garlic, err := onramp.NewGarlic("metalistener-"+name, "127.0.0.1:7656", onramp.OPT_WIDE)
	if err != nil {
		return nil, err
	}
	_, port, err := net.SplitHostPort(name)
	if err != nil {
		port = "3000"
	}
	onions := make(map[string]*onramp.Onion)
	garlics := make(map[string]*onramp.Garlic)
	onions[port] = onion
	garlics[port] = garlic
	ml := &Mirror{
		MetaListener: inner,
		Onions:       onions,
		Garlics:      garlics,
	}
	return ml, nil
}

func (ml Mirror) Listen(name, addr, certdir string, hiddenTls bool) (net.Listener, error) {
	log.Println("Starting Mirror Listener")
	log.Printf("Actual args: name: '%s' addr: '%s' certDir: '%s' hiddenTls: '%t'\n", name, addr, certdir, hiddenTls)
	// get the port:
	_, port, err := net.SplitHostPort(name)
	if err != nil {
		// check if host is an IP address
		if net.ParseIP(name) == nil {
			// host = "127.0.0.1"
		}
		port = "3000"
	}
	localAddr := net.JoinHostPort("127.0.0.1", port)
	// Listen on plain HTTP
	tcpListener, err := net.Listen("tcp", localAddr)
	if err != nil {
		return nil, err
	}
	if err := ml.AddListener("http", tcpListener); err != nil {
		return nil, err
	}
	log.Printf("HTTP Local listener added http://%s\n", tcpListener.Addr())
	log.Println("Checking for existing onion and garlic listeners")
	// Check if onion and garlic listeners already exist
	if ml.Onions[port] == nil {
		// make a new onion listener
		// and add it to the map
		log.Println("Creating new onion listener")
		onion, err := onramp.NewOnion("metalistener-" + name + port)
		if err != nil {
			return nil, err
		}
		log.Panicln("Onion listener created")
		ml.Onions[port] = onion
	}
	if ml.Garlics[port] == nil {
		// make a new garlic listener
		// and add it to the map
		log.Println("Creating new garlic listener")
		garlic, err := onramp.NewGarlic("metalistener-"+name+port, "127.0.0.1:7656", onramp.OPT_WIDE)
		if err != nil {
			return nil, err
		}
		log.Panicln("Garlic listener created")
		ml.Garlics[port] = garlic
	}
	if hiddenTls {
		// make sure an onion and a garlic listener exist at ml.Onions[port] and ml.Garlics[port]
		// and listen on them, check existence first
		onionListener, err := ml.Onions[port].ListenTLS()
		if err != nil {
			return nil, err
		}
		if err := ml.AddListener("onion", onionListener); err != nil {
			return nil, err
		}
		log.Printf("OnionTLS listener added https://%s\n", onionListener.Addr())
		garlicListener, err := ml.Garlics[port].ListenTLS()
		if err != nil {
			return nil, err
		}
		if err := ml.AddListener("garlic", garlicListener); err != nil {
			return nil, err
		}
		log.Printf("GarlicTLS listener added https://%s\n", garlicListener.Addr())
	} else {
		onionListener, err := ml.Onions[port].Listen()
		if err != nil {
			return nil, err
		}
		if err := ml.AddListener("onion", onionListener); err != nil {
			return nil, err
		}
		log.Printf("Onion listener added http://%s\n", onionListener.Addr())
		garlicListener, err := ml.Garlics[port].Listen()
		if err != nil {
			return nil, err
		}
		if err := ml.AddListener("garlic", garlicListener); err != nil {
			return nil, err
		}
		log.Printf("Garlic listener added http://%s\n", garlicListener.Addr())
	}
	if addr != "" {
		cfg := wileedot.Config{
			Domain:         name,
			AllowedDomains: []string{name},
			CertDir:        certdir,
			Email:          addr,
		}
		tlsListener, err := wileedot.New(cfg)
		if err != nil {
			return nil, err
		}
		if err := ml.AddListener("tls", tlsListener); err != nil {
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
func Listen(name, addr, certdir string, hiddenTls bool) (net.Listener, error) {
	ml, err := NewMirror(name)
	if err != nil {
		return nil, err
	}
	return ml.Listen(name, addr, certdir, hiddenTls)
}
