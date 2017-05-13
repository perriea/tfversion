package tfinstall

import (
	"flag"
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/system/files"
	"github.com/perriea/tfversion/terraform/download"
	"github.com/perriea/tfversion/terraform/init"
)

var (
	pathBin    string
	pathZip    string
	pathTmp    string
	check      bool
	usr        *user.User
	errNetwork bool
	err        error
)

func init() {
	usr, err = user.Current()
	tferror.Panic(err)

	pathBin = filepath.Join(usr.HomeDir, "/terraform/bin/")
	pathTmp = filepath.Join(usr.HomeDir, "/terraform/tmp/")
	pathZip = pathTmp + "/terraform-%s.zip"
}

func Run(params []string) error {

	var dl *flag.FlagSet

	dl = flag.NewFlagSet("install", flag.ExitOnError)
	dl.Parse(params)
	params = dl.Args()

	if len(params) != 1 {
		return fmt.Errorf("One parameter is accepted ...")
	}

	// Lauch Terraform download
	tfinit.CreateTree()
	check = tfdownload.Run(params[0])

	// Check if download is done and install
	if check {
		// Unzip zip archive
		fmt.Println("Unzip file ...")
		tffiles.UnZip(fmt.Sprintf(pathZip, params[0]), pathBin)

		fmt.Println("Install the binary file ...")
		tffiles.CreateText(params[0])

		tferror.Run(1, fmt.Sprintf("Installed %s, Thanks ! ♥", params[0]))
	}

	return err
}
