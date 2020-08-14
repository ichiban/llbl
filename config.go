package llbl

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	resolver = "/etc/resolver"
)

// Configure configures /etc/resolver/localhost to lookup llbl.
func Configure(port int) (func() error, error) {
	localhost := filepath.Join(resolver, "localhost")

	if err := os.MkdirAll(resolver, 0644); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(localhost, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	if _, err := fmt.Fprintf(f, "nameserver %s\nport %d", ipv4loopback, port); err != nil {
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return func() error {
		return os.Remove(localhost)
	}, nil
}
