package tflist

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
	"github.com/ryanuber/columnize"
)

var (
	online  bool
	offline bool
	clist   *flag.FlagSet
	good    *color.Color
	err     error
)

func init() {
	good = color.New(color.FgGreen, color.Bold)
	clist = flag.NewFlagSet("list", flag.ExitOnError)
	clist.BoolVar(&online, "on", false, "View all versions available.")
	clist.BoolVar(&offline, "off", false, "View all version already downloaded.")
}

func showList(list []string, tfversion string) {

	var (
		max int
		i   int
		k   int
	)

	i = 0
	max = 10
	reslist := []string{}

	for i < len(list) {
		newlist := ""
		for k <= max {
			if (k != max) && (len(list)-i) > 0 {
				if list[i] == tfversion {
					newlist = newlist + good.Sprintf(list[i]) + " | "
				} else {
					newlist = newlist + list[i] + " | "
				}
			} else {
				if (len(list) - i) >= 0 {
					reslist = append(reslist, newlist)
				}
			}
			k++
			i++
		}
		k = 0
	}

	result := columnize.SimpleFormat(reslist)
	fmt.Println(result)
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
