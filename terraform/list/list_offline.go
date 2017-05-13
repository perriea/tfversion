package tflist

import (
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

		tferror.Run(0, "[INFO] All local version:")

		files, err := ioutil.ReadDir(filepath.Join(usr.HomeDir, "/terraform/tmp/"))
		tferror.Panic(err)

		for _, f := range files {
			tfversion = r.FindString(f.Name())
			if tfversion != "" {
				if tfversion == string(tversion) {
					tferror.Run(1, tfversion)
				} else {
					tferror.Run(-1, tfversion)
				}
				count++
			}
		}

		if count == 0 {
			tferror.Run(0, "[INFO] No local versions !")
		}
	} else {
		tferror.Run(2, "[WARN] No installed version yet")
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
		tferror.Run(0, "[INFO] Nothing deleted !")
	} else {
		tferror.Run(1, "All files are deleted !")
	}
}
