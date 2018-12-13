package terraform

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestDownload(t *testing.T) {
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

		err = release.Download(true)
		if err != nil {
			t.Fatalf("download failed (%s)\n", version)
		} else {
			t.Logf("download OK (%s)\n", version)
		}
	}
}

// TestInstall : testing installation
func TestInstall(t *testing.T) {
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

		err = release.Install(true)
		if err != nil {
			t.Fatalf("installation failed (%s)\n", version)
		} else {
			fmt.Printf("installation OK (%s)\n", version)
		}
	}
}
