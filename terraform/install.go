package terraform

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// Download : Launch download
func (release Release) Download(quiet bool) error {
	// Formulation URL Terraform Website
	url := fmt.Sprintf(release.Repository, release.Version, release.Version, runtime.GOOS, runtime.GOARCH)

	// Request GET URL
	resp, err := release.HTTPclient.Get("https://" + url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Verify code equal 200
	if (err == nil) && (resp.StatusCode == 200) {
		Message("\033[1;32mStart download ...\n", quiet)
		fileUnzip, err := os.Create(fmt.Sprintf("%s%sterraform-%s.zip", release.Home, folders["tmp"], release.Version))
		if err != nil {
			return err
		}
		defer fileUnzip.Close()

		// Copy reponse in file
		_, err = io.Copy(fileUnzip, resp.Body)
		if err != nil {
			return err
		}
		return nil
	}
	Message("\033[1;31m[ERROR] Failed, this version doesn't exist !", quiet)
	return nil
}

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
	Message("Install the binary file ...", quiet)
	if err := release.unZip(filepath.Join(release.Home, folders["tmp"], fmt.Sprintf("/terraform-%s.zip", release.Version)), filepath.Join(release.Home, folders["bin"])); err != nil {
		return err
	}

	// Save version in file
	if err := ioutil.WriteFile(filepath.Join(release.Home, folders["tmp"], ".version"), []byte(release.Version), 0600); err != nil {
		return err
	}

	Message(fmt.Sprintf("v%s installed ♥", release.Version), quiet)

	return nil
}
