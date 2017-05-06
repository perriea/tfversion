package main

import (
	"flag"

	"github.com/perriea/tfversion/cloud/aws"
	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/system/folders"
	"github.com/perriea/tfversion/system/network"
	"github.com/perriea/tfversion/terraform/download"
	"github.com/perriea/tfversion/terraform/install"
	"github.com/perriea/tfversion/terraform/list"
)

var (
	// Errors
	errNetwork bool
	check      bool
	err        error
	// Flag func version param
	pathBin string
	pathTmp string
	install string
	// Flag launch func List
	listOnline  bool
	listOffline bool
	cleanup     bool
	awsconfig   bool
)

func init() {
	// Simply error
	check = false
	errNetwork = false
	// Paths
	pathBin = "/terraform/bin/"
	pathTmp = "/terraform/tmp/"
	// Flags CLI
	flag.BoolVar(&listOnline, "liston", false, "List online version of terraform")
	flag.BoolVar(&listOffline, "listoff", false, "List local version of terraform")
	flag.BoolVar(&cleanup, "cleanup", false, "Clean cache (tmp files)")
	flag.StringVar(&install, "install", "0", "Version of terraform to install or switch")
	flag.BoolVar(&awsconfig, "awsconfig", false, "Test AWSConfig")
	flag.Parse()
}

func main() {
	// Check if internet is available (releases.hashicorp.com)
	errNetwork = tfnetwork.Run()
	if listOnline == true {

		if errNetwork {
			// Show version terraform
			tflist.ListOn()
		} else {
			// No network
			tferror.Run(2, "[ERROR] No internet connection ...")
		}

	} else if listOffline == true {
		// List all versions local
		tflist.ListOff()

	} else if install != "0" {

		if errNetwork {
			// Lauch Terraform download
			tffolder.CreateFolder(pathTmp, 0755)
			check = tfdownload.Run(install)

			// Check if download is done and install
			if check {
				tffolder.CreateFolder(pathBin, 0755)
				tfinstall.Run(install)
			}

		} else {
			// No network
			tferror.Run(2, "[WARN] No internet connection ...")
		}

	} else if cleanup {
		// Delete all cache
		tflist.Cleanup()

	} else if awsconfig {
		aws.TestConnect()
	} else {
		// Show version
		ShowVersion()
	}
}
