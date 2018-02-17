package terraform

import (
	"fmt"
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

		test = Download(version)
		if !test {
			t.Fatalf("download failed (%s)\n", version)
		} else {
			fmt.Printf("download OK (%s)\n", version)
		}
	}
}
