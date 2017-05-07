package tfinstall

import (
	"flag"
	"fmt"
	"os/user"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/system/files"
	"github.com/perriea/tfversion/system/folders"
	"github.com/perriea/tfversion/terraform/download"
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

	pathBin = usr.HomeDir + "/terraform/bin/"
	pathTmp = usr.HomeDir + "/terraform/tmp/"
	pathZip = pathTmp + "terraform-%s.zip"
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
	tffolder.CreateFolder(pathTmp, 0755)
	check = tfdownload.Run(params[0])

	// Check if download is done and install
	if check {
		tffolder.CreateFolder(pathBin, 0755)
		// Unzip zip archive
		fmt.Println("Unzip file ...")
		tffiles.UnZip(fmt.Sprintf(pathZip, params[0]), pathBin)

		fmt.Println("Install the binary file ...")
		tffiles.CreateText(params[0])

		tferror.Run(1, fmt.Sprintf("Installed %s, Thanks ! ♥", params[0]))
	}

	return err
}
