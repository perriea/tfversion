package terraform

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateFolder on disk
func (r *Release) CreateFolder(folder string) error {
	var err error

	if err = os.MkdirAll(filepath.Join(r.Home, folder), os.FileMode(0755)); err != nil {
		return err
	}

	return nil
}

// InitFolder : Create folders (init)
func (r *Release) InitFolder() error {
	var err error

	if err = r.CreateFolder(PathBin.toString()); err != nil {
		return err
	}

	if err = r.CreateFolder(PathTmp.toString()); err != nil {
		return err
	}

	return nil
}

// Message : Quiet mode
func Message(message string, quiet bool) {
	if !quiet {
		fmt.Println(message)
	}
}
