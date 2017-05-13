package tfuninstall

import (
	"flag"
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
	clean   *flag.FlagSet
	err     error
)

func init() {
	usr, err = user.Current()
	tferror.Panic(err)

	count = 0

	clean = flag.NewFlagSet("uninstall", flag.ExitOnError)
	clean.BoolVar(&all, "all", false, "Delete all versions locale.")
}

func Run(params []string) error {

	clean.Parse(params)

	if len(params) != 1 {
		return fmt.Errorf("Only one argument is accepted.")
	}

	files, err := ioutil.ReadDir(filepath.Join(usr.HomeDir, "/terraform/tmp/"))
	tferror.Panic(err)
	for _, f := range files {
		err = os.Remove(filepath.Join(usr.HomeDir, "/terraform/tmp/", f.Name()))
		tferror.Panic(err)
		count++
	}

	if count == 0 {
		tferror.Run(0, "[INFO] Nothing deleted !")
	} else {
		tferror.Run(1, "All files are deleted !")
	}

	return nil
}
