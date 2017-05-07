package tfuninstall

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/perriea/tfversion/error"
)

var (
	path    string
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
	path = usr.HomeDir + "/terraform/tmp/"

	clean = flag.NewFlagSet("uninstall", flag.ExitOnError)
	clean.BoolVar(&all, "a", false, "Delete all versions locale.")
	clean.StringVar(&version, "v", "0", "Delete one version locale.")
}

func Run(params []string) error {

	clean.Parse(params)

	if all && version != "0" {
		return fmt.Errorf("-on and -off are mutually exclusive")
	}

	if len(params) != 1 {
		return fmt.Errorf("Only one argument is accepted.")
	}

	files, err := ioutil.ReadDir(path)
	tferror.Panic(err)
	for _, f := range files {
		err = os.Remove(path + f.Name())
		tferror.Panic(err)
		count++
	}

	if count == 0 {
		tferror.Run(2, "Nothing deleted !")
	} else {
		tferror.Run(1, "All files are deleted !")
	}

	return nil
}
