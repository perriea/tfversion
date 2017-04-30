package tfinstall

import (
  "os/user"
  "fmt"

  "github.com/perriea/tfversion/system/zip"
  "github.com/perriea/tfversion/error"
)

var (
    path_bin  string
    path_zip  string
    usr       *user.User

    err       error
)

func init()  {
  
    usr, err = user.Current()
    tferror.Panic(err)

    path_zip = usr.HomeDir + "/terraform/tmp/terraform-%s.zip"
    path_bin = usr.HomeDir + "/terraform/bin/"
}

func Run(version string)  {
    // Unzip zip archive
    fmt.Println("Unzip file ...")
    tfzip.Run(fmt.Sprintf(path_zip, version), path_bin)
    fmt.Println("Install the binary file ...")

    tferror.Run(1, fmt.Sprintf("Installed %s, Thanks ! ♥\n", version))
}
