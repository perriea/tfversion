package terraform

import (
	"fmt"
	"path/filepath"
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
	var (
		paths []string
	)

	paths = []string{filepath.Join(home, tfVersionHomePath), filepath.Join(home, tfVersionHomeBin)}

	for _, path := range paths {

		err = UnInstallAll(path)
		if err != nil {
			t.Fatalf("uninstall failed (%s)\n", version)
		} else {
			fmt.Printf("uninstall OK (%s)\n", version)
		}
	}
}
