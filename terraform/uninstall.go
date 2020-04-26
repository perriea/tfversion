package terraform

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// UnInstallAll Terraform versions
func (r *Release) UnInstallAll(quiet bool) error {
	var (
		files []os.FileInfo
		count int
		err   error
	)

	if files, err = ioutil.ReadDir(filepath.Join(r.Home, PathTmp.toString())); err != nil {
		return err
	}

	for _, file := range files {
		if err = os.Remove(filepath.Join(r.Home, PathTmp.toString(), file.Name())); err != nil {
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
func (r *Release) UnInstall(quiet bool) error {
	var (
		files []os.FileInfo
		count int
		err   error
	)

	if files, err = ioutil.ReadDir(filepath.Join(r.Home, PathTmp.toString())); err != nil {
		return err
	}

	for _, file := range files {
		if file.Name() == fmt.Sprintf("terraform-%s.zip", r.Version) {
			if err = os.Remove(filepath.Join(r.Home, PathTmp.toString(), file.Name())); err != nil {
				return err
			}

			Message(fmt.Sprintf("Version %s is deleted !\n", r.Version), quiet)
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
