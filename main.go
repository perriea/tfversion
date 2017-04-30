package main

import (
    "flag"
    "errors"

    "github.com/fatih/color"
    "github.com/perriea/tfversion/terraform/download"
    "github.com/perriea/tfversion/terraform/install"
    "github.com/perriea/tfversion/terraform/list"
)

var (
    // Errors
    err_msg string
    check   bool
    err     error
    fatal   *color.Color
    // Flag func version param
    version string
    // Flag launch func List
    list    bool
)

func init()  {
    check = false

    // Init show error message
    fatal = color.New(color.FgRed, color.Bold)
    err_msg = "[ERROR] Too many argument or nothing !"

    // Flags CLI
    flag.BoolVar(&list, "list", false, "List version of terraform")
    flag.StringVar(&version, "version", "0", "Version of terraform to install or switch")
    flag.Parse()
}

func main()  {

    if list == true && version == "0" {
        // Show version terraform
        tflist.Run()
    } else if list == false && version != "0" {
        // Lauch Terraform download
        check = tfdownload.Run(version)
        // Check if download is done and install
        if check {
          tfinstall.Run(version)
        }
    } else {
        err = errors.New(err_msg)
        if err != nil {
          fatal.Println(err)
        }
    }
}
