package terraform

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// UninstallOne : Delete one version
func UninstallOne(version string) error {
	count := 0

	files, err := ioutil.ReadDir(filepath.Join(home, tfVersionHomePath))
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.Name() == fmt.Sprintf("terraform-%s.zip", version) {
			err = os.Remove(filepath.Join(home, tfVersionHomePath, f.Name()))
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

// UninstallAll : Delete all files in folder
func UninstallAll(path string) error {

	files, err := ioutil.ReadDir(filepath.Join(home, tfVersionHomePath))
	if err != nil {
		return err
	}

	for _, f := range files {
		err = os.Remove(filepath.Join(home, tfVersionHomePath, f.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}
