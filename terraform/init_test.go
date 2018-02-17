package terraform

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// TestInit : Testing creation folder
func TestInit(t *testing.T) {
	var (
		paths []string
		info  os.FileInfo
	)

	paths = []string{filepath.Join(home, tfVersionHomePath), filepath.Join(home, tfVersionHomeBin)}

	for _, path := range paths {
		Init()

		info, err = os.Stat(path)
		if err != nil && info.IsDir() {
			t.Fatalf("folder created (%s)\n", path)
		} else {
			fmt.Printf("folder OK (%s)\n", path)
		}
	}
}
