package tffolder

import (
	"os"
	"path/filepath"
)

var (
	err error
)

// MakeFolder : create folder
func MakeFolder(name string, chmod int) error {
	err = os.MkdirAll(filepath.Join(name), os.FileMode(chmod))
	return err
}
