package main

import (
    "flag"

    "github.com/perriea/tfversion/network"
    "github.com/perriea/tfversion/terraform/download"
    "github.com/perriea/tfversion/terraform/install"
    "github.com/perriea/tfversion/terraform/list"
    "github.com/perriea/tfversion/error"
)

var (
    // Errors
    err_msg           string
    err_msg_network   string
    err_network       bool
    check             bool
    err               error
    // Flag func version param
    version           string
    // Flag launch func List
    list              bool
)

func init()  {
    check = false
    err_network = false

    // Init show error message
    err_msg = "[ERROR] Too many argument or nothing !"
    err_msg_network = "[ERROR] No internet connection ..."

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
            tferror.Run(3, err_msg_network)
        }

    } else if list == false && version != "0" {

        if err_network {
            // Lauch Terraform download
            check = tfdownload.Run(version)
            // Check if download is done and install
            if check {
                tfinstall.Run(version)
            }

        } else {
            // No network
            tferror.Run(3, err_msg_network)
        }

    } else {
        // Error args
        tferror.Run(2, err_msg)
    }
}
