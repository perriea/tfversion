package tfnetwork

import "testing"

// TestRun : require internet connection
func TestRun(t *testing.T) {

	var (
		Result bool
		host   string
		hosts  []string
	)

	hosts = []string{
		"github.com:80",
		"github.com:443",
		"releases.hashicorp.com:80",
		"releases.hashicorp.com:443",
	}

	for _, host = range hosts {
		Result = Run(host, 3, true)
		if !Result {
			t.Fatalf("Error request (%s)", host)
		}
	}
}
