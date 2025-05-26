package meta

import "net"

// MetaAddr implements the net.Addr interface for a meta listener.
type MetaAddr struct {
	addresses []net.Addr
}

// Network returns the name of the network.
func (ma *MetaAddr) Network() string {
	return "meta"
}

// String returns a string representation of all managed addresses.
func (ma *MetaAddr) String() string {
	if len(ma.addresses) == 0 {
		return "meta(empty)"
	}

	result := "meta("
	for i, addr := range ma.addresses {
		if i > 0 {
			result += ", "
		}
		result += addr.String()
	}
	result += ")"

	return result
}
