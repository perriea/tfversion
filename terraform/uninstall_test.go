package terraform

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

// TestUnInstall : testing installation
func TestUnInstall(t *testing.T) {
	var (
		home     string
		versions []string
		err      error
	)

	versions = []string{"0.11.0", "0.11.0-beta1", "0.11.0-rc1", "0.10.8", "0.7.2", "0.1.0"}

	home, err = homedir.Dir()
	if err != nil {
		panic(err)
	}

	for _, version := range versions {

		release := Release{
			Home:       home,
			Version:    version,
			HTTPclient: &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
			Repository: "releases.hashicorp.com/terraform/%s/terraform_%s_%s_%s.zip",
		}

		err = release.UnInstall(true)
		if err != nil {
			t.Fatalf("uninstall failed (%s)\n", version)
		} else {
			fmt.Printf("uninstall OK (%s)\n", version)
		}
	}
}
