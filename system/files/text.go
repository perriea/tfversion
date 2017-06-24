package tffiles

import (
	"io/ioutil"
	"os/user"
	"path/filepath"

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

// CreateText : Creating a file with the active version
func CreateText(version string) {
	fileByte := []byte(version)
	err = ioutil.WriteFile(filepath.Join(usr.HomeDir, "/.tfversion/tmp/.version"), fileByte, 0600)
	tferror.Panic(err)
}
