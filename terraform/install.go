package terraform

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// Download : Launch download
func (r *Release) download(quiet bool) error {
	var (
		url       string = fmt.Sprintf(PathTerraform.toString(), r.Version, r.Version, runtime.GOOS, runtime.GOARCH)
		path      string = fmt.Sprintf("%s%sterraform-%s.zip", r.Home, PathTmp.toString(), r.Version)
		resp      *http.Response
		fileUnzip *os.File
		err       error
	)

	// Request GET URL
	if resp, err = r.HTTPclient.Get(url); err != nil {
		return err
	}
	defer resp.Body.Close()

	// Verify code equal 200
	if resp.StatusCode == http.StatusOK {
		Message("Downloading ...", quiet)

		if fileUnzip, err = os.Create(path); err != nil {
			return err
		}
		defer fileUnzip.Close()

		// Copy reponse in file
		_, err = io.Copy(fileUnzip, resp.Body)

		return err
	}

	return errors.New("failed, this version doesn't exist")
}

// UnZipFile : UnZip one file
func (r *Release) unZip(archive string, target string) error {
	var (
		path       string
		fileReader io.ReadCloser
		reader     *zip.ReadCloser
		targetFile *os.File
		err        error
	)

	if reader, err = zip.OpenReader(filepath.Join(archive)); err != nil {
		return err
	}

	for _, file := range reader.File {
		path = filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		if fileReader, err = file.Open(); err != nil {
			return err
		}
		defer fileReader.Close()

		if targetFile, err = os.OpenFile(filepath.Join(r.Home, PathBin.toString(), "terraform"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode()); err != nil {
			return err
		}
		defer targetFile.Close()

		_, err = io.Copy(targetFile, fileReader)

		return err
	}

	return err
}

// Install Terraform versions
func (r *Release) install(quiet bool) error {
	var (
		err error
	)

	// UnZip archive
	Message("Installing ...", quiet)
	if err = r.unZip(filepath.Join(r.Home, PathTmp.toString(), fmt.Sprintf("/terraform-%s.zip", r.Version)), filepath.Join(r.Home, PathBin.toString())); err != nil {
		return err
	}

	// Save version in file
	if err = ioutil.WriteFile(filepath.Join(r.Home, PathTmp.toString(), ".version"), []byte(r.Version), 0600); err != nil {
		return err
	}

	Message(fmt.Sprintf("v%s installed ♥", r.Version), quiet)

	return nil
}
