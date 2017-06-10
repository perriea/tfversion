package tflist

import (
	"fmt"

	"github.com/ryanuber/columnize"
)

var (
	online  bool
	offline bool
	err     error
)

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
					newlist = "\033[0;37m" + newlist + fmt.Sprintf("\033[1;32m"+list[i]+"\033[0;37m") + " | "
				} else {
					newlist = "\033[0;37m" + newlist + list[i] + " | "
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
