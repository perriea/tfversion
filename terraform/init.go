package terraform

import (
	"os"
	"path/filepath"

	"github.com/perriea/tfversion/errors"
)

// Init : Create folders (init)
func Init() {
	// var tfpaths = []string{"/.tfversion", "/.tfversion/tmp", "/.tfversion/bin"}
	var tfpaths = []string{tfVersionHomePath, tfVersionHomeBin}

	for _, tfpath := range tfpaths {
		err = os.MkdirAll(filepath.Join(home, tfpath), os.FileMode(0755))
		errors.Panic(err)
	}
}
