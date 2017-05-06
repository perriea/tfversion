package tffiles

import (
	"io/ioutil"
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

func CreateText(version string) {

	fileByte := []byte(version)

	tffolder.CreateFolder(usr.HomeDir+"/terraform/tmp/", 0755)
	err = ioutil.WriteFile(usr.HomeDir+"/terraform/tmp/version.txt", fileByte, 0600)
	tferror.Panic(err)
}
