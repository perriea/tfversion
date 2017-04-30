package tfinstall

import (
  "archive/zip"
  "os"
  "io"
  "path/filepath"
  "os/exec"
  "fmt"

  "github.com/perriea/tfversion/error"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Unzip(archive, target string) error {

  	reader, err := zip.OpenReader(archive)
  	tferror.Panic(err)

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
    		tferror.Panic(err)
    		defer fileReader.Close()

    		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
    		tferror.Panic(err)
    		defer targetFile.Close()

    		if _, err := io.Copy(targetFile, fileReader); err != nil {
    			   return err
    		}

        err = os.Remove(archive)
        tferror.Panic(err)
  	}

  	return nil
}

func Run(version string)  {

    var (
      err error
      cmd *exec.Cmd
    )

    // Unzip zip archive
    fmt.Printf("Unzip file ...\n")
    err = Unzip("/tmp/terraform-" + version + ".zip", "/tmp/")
    tferror.Panic(err)

    // Create alias command
    fmt.Printf("Install the binary file ...\n")
    cmd = exec.Command("alias", "terraform=/tmp/terraform")
    err = cmd.Run()
    tferror.Panic(err)

    tferror.Run(1, fmt.Sprintf("Installed %s, Thanks ! ♥\n", version))
}
