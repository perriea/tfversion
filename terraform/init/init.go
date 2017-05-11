package tfinit

import (
	"os/user"

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
	tffolder.CreateFolder(usr.HomeDir+"/terraform", 0755)
	tffolder.CreateFolder(usr.HomeDir+"/terraform/tmp", 0755)
	tffolder.CreateFolder(usr.HomeDir+"/terraform/bin", 0755)
}
