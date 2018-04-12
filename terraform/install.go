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
func unZipArchive(archive string, target string) error {

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

		targetFile, err := os.OpenFile(filepath.Join(home, tfVersionHomeBin, "terraform"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
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
func Install(version string, quiet bool) error {

	// Lauch Terraform download
	check, err := Download(version, quiet)

	// Check if download is ok and install
	if check && err == nil {
		// UnZip archive
		ShowMessage("\033[1;33mInstall the binary file ...", quiet)
		if err = unZipArchive(filepath.Join(home, tfVersionHomePath, fmt.Sprintf("/terraform-%s.zip", version)), filepath.Join(home, tfVersionHomeBin)); err != nil {
			return err
		}

		// Save version in file
		if err = ioutil.WriteFile(filepath.Join(home, tfVersionHomePath, ".version"), []byte(version), 0600); err != nil {
			return err
		}

		ShowMessage(fmt.Sprintf("\033[1;32mv%s installed ♥", version), quiet)
	}

	return err
}

// UnInstall Terraform versions
func UnInstall(version string, quiet bool) error {
	var (
		files []os.FileInfo
		count int
	)

	files, err = ioutil.ReadDir(filepath.Join(home, tfVersionHomePath))
	if err != nil {
		return err
	}

	for _, file := range files {
		if version != "all" {
			if file.Name() == fmt.Sprintf("terraform-%s.zip", version) {
				if err = os.Remove(filepath.Join(home, tfVersionHomePath, file.Name())); err != nil {
					return err
				}

				ShowMessage(fmt.Sprintf("\033[1;32mVersion %s is deleted !\n", version), quiet)
				return nil
			}
		} else if version == "all" {
			if err = os.Remove(filepath.Join(home, tfVersionHomePath, file.Name())); err != nil {
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
