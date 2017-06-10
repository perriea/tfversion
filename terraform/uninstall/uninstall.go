package tfuninstall

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/perriea/tfversion/error"
)

var (
	count   int
	all     bool
	version string
	usr     *user.User
	err     error
)

func init() {
	usr, err = user.Current()
	tferror.Panic(err)
}

func Uniq(version string) error {

	count = 0
	files, err := ioutil.ReadDir(filepath.Join(usr.HomeDir, "/terraform/tmp/"))

	tferror.Panic(err)
	for _, f := range files {
		if f.Name() == fmt.Sprintf("terraform-%s.zip", version) {
			err = os.Remove(filepath.Join(usr.HomeDir, "/terraform/tmp/", f.Name()))
			tferror.Panic(err)
			count++
		}
	}

	if count == 0 {
		fmt.Printf("\033[1;34m[INFO] Nothing deleted !\n")
	} else {
		fmt.Printf("\033[1;32mVersion is deleted !\n")
	}

	return nil
}

func All() error {

	files, err := ioutil.ReadDir(filepath.Join(usr.HomeDir, "/terraform/tmp/"))
	tferror.Panic(err)
	for _, f := range files {
		err = os.Remove(filepath.Join(usr.HomeDir, "/terraform/tmp/", f.Name()))
		tferror.Panic(err)
		count++
	}

	if count == 0 {
		fmt.Printf("\033[1;34m[INFO] Nothing deleted !\n")
	} else {
		fmt.Printf("\033[1;32mAll files are deleted !\n")
	}

	return nil
}
