package tfuninstall

import (
	"fmt"
	"path/filepath"
	"testing"
)

// TestOneVersion : Delete binary
func TestOneVersion(t *testing.T) {
	err = OneVersion("0.9.6")
	if err != nil {
		t.Fatalf("Error suppression files !")
	}
}

// TestAll : Delete all cache
func TestAll(t *testing.T) {
	err = All(filepath.Join(usr.HomeDir, "/.tfversion/tmp/"))
	if err != nil {
		t.Fatalf("Error suppression files !")
	} else {
		fmt.Print("All files deleted")
	}
}
