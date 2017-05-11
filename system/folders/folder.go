package tffolder

import (
	"os"

	"github.com/perriea/tfversion/error"
)

var (
	err error
)

func CreateFolder(name string, chmod int) {
	err = os.MkdirAll(name, os.FileMode(chmod))
	tferror.Panic(err)
}
