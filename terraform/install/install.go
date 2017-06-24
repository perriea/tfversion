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
	// different path
	pathBin string
	pathZip string
	pathTmp string

	// check error
	check      bool
	errNetwork bool

	// user parameters
	usr *user.User
	err error
)

func init() {
	usr, err = user.Current()
	tferror.Panic(err)

	pathBin = filepath.Join(usr.HomeDir, "/.tfversion/bin/")
	pathTmp = filepath.Join(usr.HomeDir, "/.tfversion/tmp/")
	pathZip = filepath.Join(pathTmp + "/terraform-%s.zip")
}

// Run : Install terraform
func Run(version string) error {

	// Lauch Terraform download
	check = tfdownload.Run(version)

	// Check if download is done and install
	if check {
		// Unzip zip archive
		fmt.Printf("\033[0;37mUnzip file ...\n")
		tffiles.UnZip(fmt.Sprintf(pathZip, version), pathBin)

		fmt.Println("\033[0;37mInstall the binary file ...")
		tffiles.CreateText(version)

		fmt.Printf("\033[1;32mInstalled %s, Thanks ! ♥\n", version)
	}

	return err
}
