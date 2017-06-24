package tflist

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"regexp"

	"github.com/perriea/tfversion/error"
)

var (
	path      string
	tfversion string
	tversion  []byte
	count     int
	usr       *user.User
)

func init() {
	usr, err = user.Current()
	tferror.Panic(err)

	count = 0
	path = usr.HomeDir + "/.tfversion/tmp/"
}

// ListOff : List version cache
func ListOff() {

	r, err := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta)[0-9]+)?")
	tferror.Panic(err)

	if _, err := os.Stat(filepath.Join(path, ".version")); !os.IsNotExist(err) {

		tversion, err = ioutil.ReadFile(filepath.Join(path, ".version"))
		tferror.Panic(err)

		files, err := ioutil.ReadDir(filepath.Join(path))
		tferror.Panic(err)

		dirs := make([]string, len(files)) //fi == your []os.FileInfo
		for i, f := range files {
			tfversion = r.FindString(f.Name())
			if tfversion != "" {
				dirs[i] = tfversion
				count++
			}
		}

		if count == 0 {
			fmt.Printf("\033[1;34mNo local versions !\n")
		} else {
			fmt.Printf("All local version:\n")
			showList(dirs, string(tversion))
		}
	} else {
		fmt.Printf("\033[1;33mNo installed version yet\n")
	}
}

// Cleanup : Delete files
func Cleanup() {

	files, err := ioutil.ReadDir(filepath.Join(path))
	tferror.Panic(err)
	for _, f := range files {
		err = os.Remove(filepath.Join(path, f.Name()))
		tferror.Panic(err)
		count++
	}

	if count == 0 {
		fmt.Printf("\033[1;34m[INFO] Nothing deleted !\n")
	} else {
		fmt.Printf("\033[1;32mAll files are deleted !\n")
	}
}
