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
	path = usr.HomeDir + "/terraform/tmp/"
}

func ListOff() {

	r, err := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta)[0-9]+)?")
	tferror.Panic(err)

	if _, err := os.Stat(filepath.Join(usr.HomeDir, "/terraform/tmp/.version")); !os.IsNotExist(err) {

		tversion, err = ioutil.ReadFile(filepath.Join(usr.HomeDir, "/terraform/tmp/.version"))
		tferror.Panic(err)

		files, err := ioutil.ReadDir(filepath.Join(usr.HomeDir, "/terraform/tmp/"))
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
			fmt.Printf("\033[1;34m[INFO] No local versions !\n")
		} else {
			fmt.Printf("\033[1;34m[INFO] All local version:\n")
			showList(dirs, string(tversion))
		}
	} else {
		fmt.Printf("\033[1;33m[WARN] No installed version yet\n")
	}
}

func Cleanup() {

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
}
