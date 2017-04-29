package main

import (
    "flag"

    "github.com/perriea/tfversion/terraform/download"
    "github.com/perriea/tfversion/terraform/install"
    "github.com/perriea/tfversion/terraform/list"
    "github.com/fatih/color"
)

var (
    version string
    list    bool
    check   bool
    fatal   *color.Color
)

func init()  {
    check = false
    fatal = color.New(color.FgRed, color.Bold)

    flag.BoolVar(&list, "list", false, "List version of terraform")
    flag.StringVar(&version, "version", "0", "Version of terraform to install or switch")
    flag.Parse()
}

func main()  {

    if list == true && version == "0" {
        tflist.Run()
    } else if list == false && version != "0" {
        check = tfdownload.Run(version)
        if check {
          tfinstall.Run(version)
        }
    } else {
        fatal.Println("Error, too many argument !")
    }
}
