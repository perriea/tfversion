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
)

var help = cli.HelpCommand("Display help informations")
var tfversion = "0.1.2"

// root command
type rootT struct {
	Version bool `cli:"v,version" usage:"Show version and check update"`
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
		if argv.Version {
			fmt.Fprintf(os.Stderr, "tfversion v%s\n\n", tfversion)

			// Show if the last version
			lastrelease, release := tfgithub.Lastversion(tfversion)
			if !lastrelease && release != nil {
				fmt.Printf("Your version is out of date ! The latest version is %s. You can update by downloading from Github (%s).", *release.TagName, *release.HTMLURL)
			}
		}
		return nil
	},
}

// child command
type installT struct {
	cli.Helper
	Version string `cli:"*version" usage:"Install or switch version"`
}

var install = &cli.Command{
	Name: "install",
	Desc: "Install new versions or switch.",
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

// child command
type uninstallT struct {
	cli.Helper
	Version string `cli:"*v,version" usage:"uninstall version"`
}

var uninstall = &cli.Command{
	Name: "uninstall",
	Desc: "Uninstall local version of Terraform",
	Argv: func() interface{} { return new(uninstallT) },
	OnBefore: func(ctx *cli.Context) error {
		tfinit.CreateTree()
		return nil
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*uninstallT)
		ctx.String("Hello, child command, I am %s\n", argv.Version)
		return nil
	},
}

// child command
type listT struct {
	cli.Helper
	On  bool `cli:"!on,online" usage:"list online version"`
	Off bool `cli:"off,offline" usage:"list offline version"`
}

var list = &cli.Command{
	Name: "list",
	Desc: "List online or offline version of terraform",
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

// child command
type testT struct {
	cli.Helper
	Aws bool `cli:"*aws,amazon" usage:"test connection to AWS"`
	//Gcp bool `cli:"*gcp,google" usage:"test connection to GCP"`
}

var test = &cli.Command{
	Name: "test",
	Desc: "Test provider cloud (AWS)",
	Argv: func() interface{} { return new(testT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*testT)
		if argv.Aws {
			tfaws.TestConnect()
		}
		return nil
	},
}
