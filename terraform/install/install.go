package tfinstall

import (
	"flag"
	"fmt"
	"os/user"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/system/files"
	"github.com/perriea/tfversion/system/folders"
	"github.com/perriea/tfversion/system/network"
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

func Version(version string) {

	// Unzip zip archive
	tferror.Run(-1, "Unzip file ...")
	tffiles.UnZip(fmt.Sprintf(pathZip, version), pathBin)

	tferror.Run(-1, "Install the binary file ...")
	tffiles.CreateText(version)

	tferror.Run(1, fmt.Sprintf("Installed %s, Thanks ! ♥\n", version))
}

func Run(params []string) error {

	var dl *flag.FlagSet

	dl = flag.NewFlagSet("install", flag.ExitOnError)
	dl.Parse(params)
	params = dl.Args()

	errNetwork = tfnetwork.Run()

	if len(params) != 1 {
		return fmt.Errorf("exactly one parameter needed for move (from path and to path)")
	}

	if errNetwork {
		// Lauch Terraform download
		tffolder.CreateFolder(pathTmp, 0755)
		check = tfdownload.Run(params[0])

		// Check if download is done and install
		if check {
			tffolder.CreateFolder(pathBin, 0755)
			Version(params[0])
		}

	} else {
		// No network
		tferror.Run(2, "[ERROR] No internet connection ...")
	}
	return err
}
