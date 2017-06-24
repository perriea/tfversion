package tffolder

import (
	"fmt"
	"testing"
)

var (
	nameFolders = []string{"test", "tfversion"}
)

// TestMakeFolder : create folder
func TestMakeFolder(t *testing.T) {

	for _, nameFolder := range nameFolders {
		err = MakeFolder(nameFolder, 755)
		if err != nil {
			t.Fatalf("Error creation folder")
		} else {
			fmt.Printf("Folder %s created\n", nameFolder)
		}
	}
}
