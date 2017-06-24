package tfuninstall

import (
	"fmt"
	"testing"
)

// TestOneBinary : Delete binary
func TestOneBinary(t *testing.T) {
	err = OneBinary()
	if err != nil {
		t.Fatalf("Error suppression files !")
	}
}

// TestOneBinary : Delete binary
func TestOneVersion(t *testing.T) {
	err = OneVersion("0.9.6")
	if err != nil {
		t.Fatalf("Error suppression files !")
	}
}

// TestAllVersion : Delete all cache
func TestAllVersion(t *testing.T) {
	err = AllVersion()
	if err != nil {
		t.Fatalf("Error suppression files !")
	} else {
		fmt.Print("All files deleted")
	}
}
