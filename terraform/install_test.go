package terraform

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// TestInstall : testing installation
func TestInstall(t *testing.T) {
	var (
		versions []string
	)

	versions = []string{"0.11.0", "0.11.0-beta1", "0.11.0-rc1", "0.10.8", "0.7.2", "0.1.0"}

	for _, version := range versions {

		err = Install(version, false)
		if err != nil {
			t.Fatalf("installation failed (%s)\n", version)
		} else {
			fmt.Printf("installation OK (%s)\n", version)
		}
	}
}

// TestInitFolder : Testing creation folder
func TestInitFolder(t *testing.T) {
	var (
		paths []string
		info  os.FileInfo
	)

	paths = []string{filepath.Join(home, tfVersionHomePath), filepath.Join(home, tfVersionHomeBin)}

	for _, path := range paths {
		InitFolder()

		info, err = os.Stat(path)
		if err != nil && info.IsDir() {
			t.Fatalf("folder created (%s)\n", path)
		} else {
			fmt.Printf("folder OK (%s)\n", path)
		}
	}
}
