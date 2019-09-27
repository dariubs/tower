package libtower

import "net"

// DNSLookup func
func DNSLookup(addr string) (*net.IPAddr, error) {
	dst, err := net.ResolveIPAddr("ip4", "google.com")
	if err != nil {
		return new(net.IPAddr), err
	}
	return dst, nil
}