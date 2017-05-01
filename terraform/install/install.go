package tfinstall

import (
  "os/user"
  "fmt"

  "github.com/perriea/tfversion/error"
  "github.com/perriea/tfversion/system/files"
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
    fmt.Printf("Unzip file ...\n")
    tffiles.UnZip(fmt.Sprintf(path_zip, version), path_bin)
    fmt.Printf("Install the binary file ...\n")

    tferror.Run(1, fmt.Sprintf("Installed %s, Thanks ! ♥\n", version))
}
