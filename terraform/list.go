package terraform

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/ryanuber/columnize"
)

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}

	return false
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
					newlist = fmt.Sprintf("\033[0;37m%s\033[1;32m%s\033[0;37m | ", newlist, list[i])
				} else {
					newlist = fmt.Sprintf("\033[0;37m%s%s | ", newlist, list[i])
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

// ListOnline : List online version
func ListOnline() {
	var (
		versions []string
		cleaned  []string
	)

	resp, err := client.Get(urlHashicorp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Verify code equal 200
	if (err == nil) && (resp.StatusCode == 200) {
		r, err := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta)[0-9]+)?")
		if err != nil {
			panic(err)
		}

		// Convert byte to string
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		newStr := buf.String()

		fmt.Printf("\033[1;34mVersions available of Terraform :\n")
		versions = r.FindAllString(newStr, -1)

		// Clean doublon
		for _, value := range versions {
			if !stringInSlice(value, cleaned) {
				cleaned = append(cleaned, value)
			}
		}

		// Show versions
		showList(cleaned, "0")
	}
}
