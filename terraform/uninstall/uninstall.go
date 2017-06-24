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

// OneBinary : delete binary
func OneBinary() error {
	files, err := ioutil.ReadDir(tfPathBin)
	if err != nil {
		return err
	}

	for _, f := range files {
		err = os.Remove(filepath.Join(tfPathBin, f.Name()))
		if err != nil {
			return err
		}
	}

	return nil
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

// AllVersion : Delete all cache
func AllVersion() error {

	files, err := ioutil.ReadDir(tfPathTmp)
	if err != nil {
		return err
	}

	for _, f := range files {
		err = os.Remove(filepath.Join(tfPathTmp, f.Name()))
		if err != nil {
			return err
		}
		count++
	}

	if count == 0 {
		fmt.Printf("Nothing deleted !\n")
	} else {
		fmt.Printf("All files are deleted !\n")
	}

	return nil
}
