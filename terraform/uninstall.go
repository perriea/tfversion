package terraform

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// UnInstall Terraform versions
func (release Release) UnInstall(quiet bool) error {
	var (
		files []os.FileInfo
		count int
	)

	files, err := ioutil.ReadDir(filepath.Join(release.Home, folders["tmp"]))
	if err != nil {
		return err
	}

	for _, file := range files {
		if release.Version != "all" {
			if file.Name() == fmt.Sprintf("terraform-%s.zip", release.Version) {
				if err = os.Remove(filepath.Join(release.Home, folders["tmp"], file.Name())); err != nil {
					return err
				}

				Message(fmt.Sprintf("\033[1;32mVersion %s is deleted !\n", release.Version), quiet)
				return nil
			}
		} else if release.Version == "all" {
			if err = os.Remove(filepath.Join(release.Home, folders["tmp"], file.Name())); err != nil {
				return err
			}

			count++
		}
	}

	if count == 0 {
		Message("\033[1;34mNo version has been removed", quiet)
	} else {
		Message("\033[1;34mAll versions have been removed", quiet)
	}

	return nil
}
