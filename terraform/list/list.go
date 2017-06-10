package tflist

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/ryanuber/columnize"
)

var (
	online  bool
	offline bool
	good    *color.Color
	err     error
)

func init() {
	good = color.New(color.FgGreen, color.Bold)
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
