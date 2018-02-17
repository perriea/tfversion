package terraform

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// UnInstallOne : Delete one version
func UnInstallOne(version string) error {
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

// UnInstallAll : Delete all files in folder
func UnInstallAll(path string) error {

	files, err := ioutil.ReadDir(filepath.Join(path))
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
