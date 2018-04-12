package terraform

import (
	"testing"
)

// TestDownload : Testing expression Downloader
func TestDownload(t *testing.T) {
	var (
		test     bool
		versions []string
	)

	versions = []string{"0.11.0", "0.11.0-beta1", "0.11.0-rc1", "0.10.8", "0.7.2", "0.1.0"}

	for _, version := range versions {
		test = false

		test, err = Download(version, true)
		if !test || err != nil {
			t.Fatalf("download failed (%s)\n", version)
		} else {
			t.Logf("download OK (%s)\n", version)
		}
	}
}
