package tfinstall

import (
	"fmt"
	"os/user"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/system/files"
)

var (
	pathBin  string
	pathZip string
	usr      *user.User
	err      error
)

func init() {
	usr, err = user.Current()
	tferror.Panic(err)

	pathZip = usr.HomeDir + "/terraform/tmp/terraform-%s.zip"
	pathBin = usr.HomeDir + "/terraform/bin/"
}

func Run(version string) {

	// Unzip zip archive
	fmt.Printf("Unzip file ...\n")
	tffiles.UnZip(fmt.Sprintf(pathZip, version), pathBin)
	fmt.Printf("Install the binary file ...\n")

	tferror.Run(1, fmt.Sprintf("Installed %s, Thanks ! ♥\n", version))
}
