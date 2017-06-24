package tffiles

import (
	"archive/zip"
	"io"
	"os"
	"os/user"
	"path/filepath"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/terraform/uninstall"
)

var (
	path     string
	files    []os.FileInfo
	tversion []byte
	reader   *zip.ReadCloser
)

func init() {
	usr, err = user.Current()
	tferror.Panic(err)
}

// UnZipFile : UnZip one file
func UnZipFile(archive, target string) error {

	reader, err = zip.OpenReader(filepath.Join(archive))
	tferror.Panic(err)

	err = tfuninstall.All(filepath.Join(usr.HomeDir, "/.tfversion/bin/"))
	if err == nil {
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

	return err
}
