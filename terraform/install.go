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
func unZipFile(archive string, target string) error {

	reader, err := zip.OpenReader(filepath.Join(archive))
	if err != nil {
		panic(err)
	}

	err = UnInstallAll(filepath.Join(home, tfVersionHomeBin))
	if err == nil {
		for _, file := range reader.File {

			path := filepath.Join(target, file.Name)
			if file.FileInfo().IsDir() {
				os.MkdirAll(path, file.Mode())
				continue
			}

			fileReader, err := file.Open()
			if err != nil {
				panic(err)
			}
			defer fileReader.Close()

			targetFile, err := os.OpenFile(filepath.Join(home, tfVersionHomeBin, "terraform"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				panic(err)
			}
			defer targetFile.Close()

			_, err = io.Copy(targetFile, fileReader)
			if err != nil {
				panic(err)
			}
		}
	}

	return err
}

// Install : Install terraform
func Install(version string, quiet bool) error {

	// Lauch Terraform download
	check := Download(version, quiet)

	// Check if download is ok and install
	if check {
		// Unzip archive
		Quiet("\033[1;33mInstall the binary file ...", quiet)
		err = unZipFile(filepath.Join(home, tfVersionHomePath, fmt.Sprintf("/terraform-%s.zip", version)), filepath.Join(home, tfVersionHomeBin))
		if err != nil {
			panic(err)
		}

		// Save version in file
		err = ioutil.WriteFile(filepath.Join(home, tfVersionHomePath, ".version"), []byte(version), 0600)
		if err != nil {
			panic(err)
		}
		Quiet(fmt.Sprintf("\033[1;32mv%s installed ♥", version), quiet)
	}

	return err
}

// Init : Create folders (init)
func InitFolder() {
	var tfpaths = []string{tfVersionHomePath, tfVersionHomeBin}

	for _, tfpath := range tfpaths {
		err = os.MkdirAll(filepath.Join(home, tfpath), os.FileMode(0755))
		if err != nil {
			panic(err)
		}
	}
}
