package tffiles

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/perriea/tfversion/error"
)

var (
	path     string
	files    []os.FileInfo
	tversion []byte
	reader   *zip.ReadCloser
)

func DelFiles(path string) {

	files, err = ioutil.ReadDir(filepath.Join(path))
	tferror.Panic(err)

	tversion, err = ioutil.ReadFile(filepath.Join(path))
	tferror.Panic(err)

	for _, f := range files {
		if !f.IsDir() {
			if !(f.Name() == string(tversion)) {
				err = os.Remove(filepath.Join(path, f.Name()))
				tferror.Panic(err)
			}
		}
	}
}

func DelOneFile(path string, target string) {

	files, err = ioutil.ReadDir(filepath.Join(target))
	tferror.Panic(err)

	tversion, err = ioutil.ReadFile(filepath.Join(path, target))
	tferror.Panic(err)

	for _, f := range files {
		if !f.IsDir() {
			if f.Name() == string(tversion) || f.Name() == ".version" {
				err = os.Remove(filepath.Join(target, f.Name()))
				tferror.Panic(err)
			}
		}
	}
}

func UnZip(archive, target string) {

	reader, err = zip.OpenReader(filepath.Join(archive))
	tferror.Panic(err)

	for _, file := range reader.File {

		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		tferror.Panic(err)
		defer fileReader.Close()

		targetFile, err := os.OpenFile(filepath.Join(path), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		tferror.Panic(err)
		defer targetFile.Close()

		_, err = io.Copy(targetFile, fileReader)
		tferror.Panic(err)
	}
}
