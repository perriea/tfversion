package tfinstall

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/system/files"
	"github.com/perriea/tfversion/terraform/download"
)

var (
	pathBin    string
	pathZip    string
	pathTmp    string
	check      bool
	errNetwork bool
	usr        *user.User
	err        error
)

func init() {
	usr, err = user.Current()
	tferror.Panic(err)

	pathBin = filepath.Join(usr.HomeDir, "/terraform/bin/")
	pathTmp = filepath.Join(usr.HomeDir, "/terraform/tmp/")
	pathZip = pathTmp + "/terraform-%s.zip"
}

func Run(version string) error {

	// Lauch Terraform download
	check = tfdownload.Run(version)

	// Check if download is done and install
	if check {
		// Unzip zip archive
		fmt.Println("Unzip file ...")
		tffiles.UnZip(fmt.Sprintf(pathZip, version), pathBin)

		fmt.Println("Install the binary file ...")
		tffiles.CreateText(version)

		tferror.Run(1, fmt.Sprintf("Installed %s, Thanks ! ♥", version))
	}

	return err
}
