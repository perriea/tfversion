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

func CreateTree() {
	tffolder.CreateFolder(filepath.Join(usr.HomeDir, "/terraform"), 0755)
	tffolder.CreateFolder(filepath.Join(usr.HomeDir, "/terraform/tmp"), 0755)
	tffolder.CreateFolder(filepath.Join(usr.HomeDir, "/terraform/bin"), 0755)
}
