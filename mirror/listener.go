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
	Onion  *onramp.Onion
	Garlic *onramp.Garlic
}

var _ net.Listener = &Mirror{}

func (m *Mirror) Close() error {
	if err := m.MetaListener.Close(); err != nil {
		log.Println("Error closing MetaListener:", err)
	}
	if err := m.Onion.Close(); err != nil {
		log.Println("Error closing Onion:", err)
	}
	if err := m.Garlic.Close(); err != nil {
		log.Println("Error closing Garlic:", err)
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
	garlic, err := onramp.NewGarlic("metalistener-"+name, "127.0.0.1:7657", onramp.OPT_WIDE)
	if err != nil {
		return nil, err
	}
	ml := &Mirror{
		MetaListener: inner,
		Onion:        onion,
		Garlic:       garlic,
	}
	return ml, nil
}

func (ml *Mirror) Listen(name, addr, certdir string, hiddenTls bool) (net.Listener, error) {
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
	} else {
		// Listen on plain HTTP
		tlsListener, err := net.Listen("tcp", "localhost:3000")
		if err != nil {
			return nil, err
		}
		if err := ml.AddListener("http", tlsListener); err != nil {
			return nil, err
		}
	}
	if hiddenTls {
		onionListener, err := ml.Onion.ListenTLS()
		if err != nil {
			return nil, err
		}
		if err := ml.AddListener("onion", onionListener); err != nil {
			return nil, err
		}
		garlicListener, err := ml.Garlic.ListenTLS()
		if err != nil {
			return nil, err
		}
		if err := ml.AddListener("garlic", garlicListener); err != nil {
			return nil, err
		}
	} else {
		onionListener, err := ml.Onion.Listen()
		if err != nil {
			return nil, err
		}
		if err := ml.AddListener("onion", onionListener); err != nil {
			return nil, err
		}
		garlicListener, err := ml.Garlic.Listen()
		if err != nil {
			return nil, err
		}
		if err := ml.AddListener("garlic", garlicListener); err != nil {
			return nil, err
		}
	}
	return ml, nil
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
