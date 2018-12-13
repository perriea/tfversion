package terraform

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// UnInstallAll Terraform versions
func (release Release) UnInstallAll(quiet bool) error {
	var (
		files []os.FileInfo
		count int
	)

	files, err := ioutil.ReadDir(filepath.Join(release.Home, folders["tmp"]))
	if err != nil {
		return err
	}

	for _, file := range files {
		if err = os.Remove(filepath.Join(release.Home, folders["tmp"], file.Name())); err != nil {
			return err
		}

		count++
	}

	if count == 0 {
		Message("No version has been removed", quiet)
	} else {
		Message("All versions have been removed", quiet)
	}

	return nil
}

// UnInstall Terraform version
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
		if file.Name() == fmt.Sprintf("terraform-%s.zip", release.Version) {
			if err = os.Remove(filepath.Join(release.Home, folders["tmp"], file.Name())); err != nil {
				return err
			}

			Message(fmt.Sprintf("Version %s is deleted !\n", release.Version), quiet)
			return nil
		}

		count++
	}

	if count == 0 {
		Message("No version has been removed", quiet)
	} else {
		Message("All versions have been removed", quiet)
	}

	return nil
}
