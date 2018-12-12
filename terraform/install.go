package terraform

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// UnZipFile : UnZip one file
func (release Release) unZip(archive string, target string) error {

	reader, err := zip.OpenReader(filepath.Join(archive))
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(filepath.Join(release.Home, folders["bin"], "terraform"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		_, err = io.Copy(targetFile, fileReader)
		if err != nil {
			return err
		}
	}

	return err
}

// Install Terraform versions
func (release Release) Install(quiet bool) error {
	// UnZip archive
	ShowMessage("\033[1;33mInstall the binary file ...", quiet)
	if err := release.unZip(filepath.Join(release.Home, folders["tmp"], fmt.Sprintf("/terraform-%s.zip", release.Version)), filepath.Join(release.Home, folders["bin"])); err != nil {
		return err
	}

	// Save version in file
	if err := ioutil.WriteFile(filepath.Join(release.Home, folders["tmp"], ".version"), []byte(release.Version), 0600); err != nil {
		return err
	}

	ShowMessage(fmt.Sprintf("\033[1;32mv%s installed ♥", release.Version), quiet)

	return nil
}

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

				ShowMessage(fmt.Sprintf("\033[1;32mVersion %s is deleted !\n", release.Version), quiet)
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
		ShowMessage("\033[1;34mNo version has been removed", quiet)
	} else {
		ShowMessage("\033[1;34mAll versions have been removed", quiet)
	}

	return nil
}
