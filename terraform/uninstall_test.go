package terraform

import (
	"fmt"
	"testing"
)

// TestUnInstallOne : testing installation
func TestUnInstallOne(t *testing.T) {
	var (
		versions []string
	)

	versions = []string{"0.11.0", "0.11.0-beta1", "0.11.0-rc1", "0.10.8", "0.7.2", "0.1.0"}

	for _, version := range versions {

		err = UnInstallOne(version)
		if err != nil {
			t.Fatalf("uninstall failed (%s)\n", version)
		} else {
			fmt.Printf("uninstall OK (%s)\n", version)
		}
	}
}

// TestUnInstallOne : testing installation
func TestUnInstallAll(t *testing.T) {
	if err = UnInstallAll(); err != nil {
		t.Fatalf("uninstall failed (%s)\n", version)
	} else {
		fmt.Printf("uninstall OK (%s)\n", version)
	}
}
