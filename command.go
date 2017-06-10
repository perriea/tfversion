package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
	"github.com/perriea/tfversion/github"
	"github.com/perriea/tfversion/terraform/cloud/aws"
	"github.com/perriea/tfversion/terraform/init"
	"github.com/perriea/tfversion/terraform/install"
	"github.com/perriea/tfversion/terraform/list"
	"github.com/perriea/tfversion/terraform/uninstall"
)

var help = cli.HelpCommand("display help informations")
var tfversion = "0.1.2"

// root command
type rootT struct {
	cli.Helper
	Version bool `cli:"v,version" usage:"show version and check update"`
}

var root = &cli.Command{
	Desc: fmt.Sprintf("tfversion v%s \n\n\033[1mUsage:\033[0m\n\n  tfversion <command> [option]", tfversion),
	Argv: func() interface{} { return new(rootT) },
	OnRootBefore: func(ctx *cli.Context) error {
		tfinit.CreateTree()
		return nil
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)

		// AutoHelper
		if argv.Help || len(os.Args) == 1 {
			ctx.WriteUsage()
			return nil
		}

		if argv.Version {
			fmt.Printf("tfversion v%s\n\n", tfversion)

			// Show if the last version
			lastrelease, release := tfgithub.Lastversion(tfversion)
			if !lastrelease && release != nil {
				fmt.Printf("Your version is out of date ! The latest version is %s. You can update by downloading from Github (%s).", *release.TagName, *release.HTMLURL)
			}
			return nil
		}
		return nil
	},
}

// install command
type installT struct {
	cli.Helper
	Version string `cli:"*version" usage:"install or switch version"`
}

var install = &cli.Command{
	Name: "install",
	Desc: "install new versions or switch.",
	Argv: func() interface{} { return new(installT) },
	OnBefore: func(ctx *cli.Context) error {
		tfinit.CreateTree()
		return nil
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*installT)

		tfinstall.Run(argv.Version)
		return nil
	},
}

// uninstall command
type uninstallT struct {
	cli.Helper
	All     bool   `cli:"a,all" usage:"delete all version (tmp)"`
	Version string `cli:"!v,version" usage:"Delete version (tmp)"`
}

var uninstall = &cli.Command{
	Name: "uninstall",
	Desc: "uninstall local version of Terraform",
	Argv: func() interface{} { return new(uninstallT) },
	OnBefore: func(ctx *cli.Context) error {
		tfinit.CreateTree()
		return nil
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*uninstallT)

		if argv.All {
			tfuninstall.All()
		}

		if argv.Version != "" {
			tfuninstall.Uniq(argv.Version)
		}

		return nil
	},
}

// list command
type listT struct {
	cli.Helper
	On  bool `cli:"!on,online" usage:"list online version"`
	Off bool `cli:"off,offline" usage:"list offline version"`
}

var list = &cli.Command{
	Name: "list",
	Desc: "list online or offline version of terraform",
	Argv: func() interface{} { return new(listT) },
	OnBefore: func(ctx *cli.Context) error {
		tfinit.CreateTree()
		return nil
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*listT)

		if argv.On {
			tflist.ListOn()
		} else {
			tflist.ListOff()
		}
		return nil
	},
}

// test command
type testT struct {
	cli.Helper
	Aws bool `cli:"*aws,amazon" usage:"test connection to AWS"`
	//Gcp bool `cli:"*gcp,google" usage:"test connection to GCP"`
}

var test = &cli.Command{
	Name: "test",
	Desc: "test provider cloud (AWS)",
	Argv: func() interface{} { return new(testT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*testT)

		if argv.Aws {
			tfaws.TestConnect()
		}
		return nil
	},
}
