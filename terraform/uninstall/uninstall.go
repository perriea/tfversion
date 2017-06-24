package tfuninstall

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/perriea/tfversion/error"
)

var (
	count     int
	all       bool
	version   string
	tfPathTmp string
	tfPathBin string
	usr       *user.User
	err       error
)

func init() {
	usr, err = user.Current()
	tferror.Panic(err)

	tfPathTmp = filepath.Join(usr.HomeDir, "/.tfversion/tmp/")
	tfPathBin = filepath.Join(usr.HomeDir, "/.tfversion/bin/")
}

// OneVersion : Delete one version
func OneVersion(version string) error {
	count = 0

	files, err := ioutil.ReadDir(tfPathTmp)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.Name() == fmt.Sprintf("terraform-%s.zip", version) {
			err = os.Remove(filepath.Join(tfPathTmp, f.Name()))
			if err != nil {
				return err
			}
			count++
		}
	}

	if count == 0 {
		fmt.Printf("\033[1;34m[INFO] Nothing deleted !\n")
	} else {
		fmt.Printf("\033[1;32mVersion %s is deleted !\n", version)
	}

	return nil
}

// All : Delete all files in folder
func All(path string) error {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		err = os.Remove(filepath.Join(path, f.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}
