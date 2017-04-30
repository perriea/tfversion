package main

import (
    "flag"

    "github.com/perriea/tfversion/system/network"
    "github.com/perriea/tfversion/terraform/download"
    "github.com/perriea/tfversion/terraform/install"
    "github.com/perriea/tfversion/terraform/list"
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
    list              bool
)

func init()  {
    check = false
    err_network = false

    // Paths
    path_bin = "/terraform/bin/"
    tmp_bin = "/terraform/tmp/"

    // Flags CLI
    flag.BoolVar(&list, "list", false, "List version of terraform")
    flag.StringVar(&version, "version", "0", "Version of terraform to install or switch")
    flag.Parse()
}

func main()  {

    err_network = tfnetwork.Run()
    if list == true && version == "0" {

        if err_network {
            // Show version terraform
            tflist.Run()
        } else {
            // No network
            tferror.Run(2, "[ERROR] No internet connection ...")
        }

    } else if list == false && version != "0" {

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

    } else {
        // Error args
        tferror.Run(2, "[ERROR] Too many argument or nothing !")
    }
}
