package terraform

import (
	"fmt"
	"testing"
)

// TestInstall : testing installation
func TestInstall(t *testing.T) {
	var (
		versions []string
	)

	versions = []string{"0.11.0", "0.11.0-beta1", "0.11.0-rc1", "0.10.8", "0.7.2", "0.1.0"}

	for _, version := range versions {
		err = Install(version, true)
		if err != nil {
			t.Fatalf("installation failed (%s)\n", version)
		} else {
			fmt.Printf("installation OK (%s)\n", version)
		}
	}
}

// TestUnInstall : testing installation
func TestUnInstall(t *testing.T) {
	var (
		versions []string
	)

	versions = []string{"0.11.0", "0.11.0-beta1", "0.11.0-rc1", "0.10.8", "0.7.2", "0.1.0", "!", "all"}

	for _, version := range versions {
		err = UnInstall(version, true)
		if err != nil {
			t.Fatalf("uninstall failed (%s)\n", version)
		} else {
			fmt.Printf("uninstall OK (%s)\n", version)
		}
	}
}
