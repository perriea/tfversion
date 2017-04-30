package main

import (
    "flag"

    "github.com/perriea/tfversion/system/network"
    "github.com/perriea/tfversion/terraform/download"
    "github.com/perriea/tfversion/terraform/install"
    "github.com/perriea/tfversion/terraform/list/online"
    "github.com/perriea/tfversion/terraform/list/offline"
    "github.com/perriea/tfversion/system/folder"
    "github.com/perriea/tfversion/error"
)

var (
    // Errors
    err_network       bool
    check             bool
    err               error
    // Flag func version param
    path_bin          string
    tmp_bin           string
    version           string
    // Flag launch func List
    list_online       bool
    list_offline      bool
    cleanup           bool
)

func init()  {
    check = false
    err_network = false

    // Paths
    path_bin = "/terraform/bin/"
    tmp_bin = "/terraform/tmp/"

    // Flags CLI
    flag.BoolVar(&list_online, "list-online", false, "List online version of terraform")
    flag.BoolVar(&list_offline, "list-offline", false, "List local version of terraform")
    flag.BoolVar(&cleanup, "cleanup", false, "Clean cache (tmp files)")
    flag.StringVar(&version, "install", "0", "Version of terraform to install or switch")
    flag.Parse()
}

func main()  {

    err_network = tfnetwork.Run()
    if list_online == true {

        if err_network {
            // Show version terraform
            tflist_online.Run()
        } else {
            // No network
            tferror.Run(2, "[ERROR] No internet connection ...")
        }

    } else if list_offline == true {
        // List all versions local
        tflist_offline.Run()
    } else if version != "0" {

        if err_network {
            // Lauch Terraform download
            tffolder.Run(tmp_bin, 0755)
            check = tfdownload.Run(version)

            // Check if download is done and install
            if check {
                tffolder.Run(path_bin, 0755)
                tfinstall.Run(version)
            }

        } else {
            // No network
            tferror.Run(2, "[WARN] No internet connection ...")
        }

    } else if cleanup {
        tflist_offline.Cleanup()
    } else {
        // Error args
        tferror.Run(2, "[ERROR] Too many argument or nothing !")
    }
}
