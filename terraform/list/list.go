package tflist

import (
	"flag"
	"fmt"
)

var (
	online  bool
	offline bool
	clist   *flag.FlagSet
	err     error
)

func init() {
	clist = flag.NewFlagSet("list", flag.ExitOnError)
	clist.BoolVar(&online, "on", false, "View all versions available.")
	clist.BoolVar(&offline, "off", false, "View all version already downloaded.")
}

func Run(params []string) error {

	clist.Parse(params)

	if online && offline {
		return fmt.Errorf("-on and -off are mutually exclusive")
	}

	if len(params) != 1 {
		return fmt.Errorf("One parameter is accepted ...")
	}

	if online {
		ListOn()
	} else {
		ListOff()
	}

	return nil
}
