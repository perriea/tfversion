package tffolder

import (
	"os"
	"path/filepath"

	"github.com/perriea/tfversion/error"
)

var (
	err error
)

func CreateFolder(name string, chmod int) {
	err = os.MkdirAll(filepath.Join(name), os.FileMode(chmod))
	tferror.Panic(err)
}
