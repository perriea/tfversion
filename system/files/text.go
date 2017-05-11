package tffiles

import (
	"io/ioutil"
	"os/user"

	"github.com/perriea/tfversion/error"
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

	err = ioutil.WriteFile(usr.HomeDir+"/terraform/tmp/.version", fileByte, 0600)
	tferror.Panic(err)
}
