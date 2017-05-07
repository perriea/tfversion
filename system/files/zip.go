package tffiles

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/perriea/tfversion/error"
)

func DelFiles(target string) {

	files, err := ioutil.ReadDir(target)
	tferror.Panic(err)

	for _, f := range files {
		if !f.IsDir() {
			err = os.Remove(target + f.Name())
			tferror.Panic(err)
		}
	}
}

func UnZip(archive, target string) {

	reader, err := zip.OpenReader(archive)
	tferror.Panic(err)

	err = os.MkdirAll(target, 0755)
	tferror.Panic(err)

	DelFiles(target)

	for _, file := range reader.File {

		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		tferror.Panic(err)
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		tferror.Panic(err)
		defer targetFile.Close()

		_, err = io.Copy(targetFile, fileReader)
		tferror.Panic(err)
	}
}
