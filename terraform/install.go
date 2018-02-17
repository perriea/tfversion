package terraform

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/perriea/tfversion/errors"
)

// UnZipFile : UnZip one file
func UnZipFile(archive string, target string) error {

	reader, err := zip.OpenReader(filepath.Join(archive))
	errors.Panic(err)

	err = UninstallAll(filepath.Join(home, tfVersionHomeBin))
	if err == nil {
		for _, file := range reader.File {

			path := filepath.Join(target, file.Name)
			if file.FileInfo().IsDir() {
				os.MkdirAll(path, file.Mode())
				continue
			}

			fileReader, err := file.Open()
			errors.Panic(err)
			defer fileReader.Close()

			targetFile, err := os.OpenFile(filepath.Join(home, tfVersionHomeBin, "terraform"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			errors.Panic(err)
			defer targetFile.Close()

			_, err = io.Copy(targetFile, fileReader)
			errors.Panic(err)
		}
	}

	return err
}

// CreateVersioning :
func CreateVersioning(version string) error {
	fileByte := []byte(version)
	err = ioutil.WriteFile(filepath.Join(home, tfVersionHomePath, ".version"), fileByte, 0600)

	return err
}

// Install : Install terraform
func Install(version string) error {

	// Lauch Terraform download
	check := Download(version)

	// Check if download is ok and install
	if check {
		// Unzip archive
		fmt.Printf("\033[0;37mUnzip file ...\n")
		err = UnZipFile(filepath.Join(home, tfVersionHomePath, fmt.Sprintf("/terraform-%s.zip", version)), filepath.Join(home, "/.tfversion/bin/"))
		errors.Panic(err)

		fmt.Println("\033[0;37mInstall the binary file ...")
		err = CreateVersioning(version)
		errors.Panic(err)
		fmt.Printf("\033[1;32mInstalled %s, Thanks ! ♥\n", version)
	}

	return err
}
