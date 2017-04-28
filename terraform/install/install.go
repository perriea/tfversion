package tfinstall

import (
  "archive/zip"
  "os"
  "io"
  "path/filepath"
  "os/exec"
  "fmt"

  "github.com/fatih/color"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	check(err)

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		check(err)
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		check(err)
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}

    err = os.Remove(archive)
    check(err)
	}

	return nil
}

func Run()  {

  var (
    err error
    cmd *exec.Cmd
  )

  good := color.New(color.FgGreen, color.Bold)

  fmt.Printf("Unzip file ...\n")
  err = unzip("/tmp/terraform-" + os.Args[1] + ".zip", "/tmp/")
  check(err)

  fmt.Printf("Install the binary file ...\n")
  cmd = exec.Command("alias", "terraform=/tmp/terraform")
  err = cmd.Run()
  check(err)

  good.Printf("Installed %s, Thanks ! ♥", os.Args[1])
}
