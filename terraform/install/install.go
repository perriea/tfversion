package tfinstall

import (
	"fmt"
	"os/user"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/system/files"
)

var (
	pathBin string
	pathZip string
	usr     *user.User
	err     error
)

func init() {
	usr, err = user.Current()
	tferror.Panic(err)

	pathZip = usr.HomeDir + "/terraform/tmp/terraform-%s.zip"
	pathBin = usr.HomeDir + "/terraform/bin/"
}

func Run(version string) {

	// Unzip zip archive
	tferror.Run(-1, "Unzip file ...")
	tffiles.UnZip(fmt.Sprintf(pathZip, version), pathBin)

	tferror.Run(-1, "Install the binary file ...")
	tffiles.CreateText(version)

	tferror.Run(1, fmt.Sprintf("Installed %s, Thanks ! ♥\n", version))
}
