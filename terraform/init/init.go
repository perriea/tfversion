package tfinit

import (
	"os/user"
	"path/filepath"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/system/folders"
)

var (
	usr *user.User
	err error
)

func init() {
	usr, err = user.Current()
	tferror.Panic(err)
}

// CreateTree : Create folders (init)
func CreateTree() {
	var tfpaths = []string{"/.terraform", "/.terraform/tmp", "/.terraform/bin"}

	for _, tfpath := range tfpaths {
		err = tffolder.MakeFolder(filepath.Join(usr.HomeDir, tfpath), 0755)
		tferror.Panic(err)
	}
}
