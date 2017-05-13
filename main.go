package main

import (
	"fmt"
	"os"

	"github.com/perriea/tfversion/cloud"
	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/github"
	"github.com/perriea/tfversion/terraform/install"
	"github.com/perriea/tfversion/terraform/list"
	"github.com/perriea/tfversion/terraform/uninstall"
)

type command struct {
	desc  string
	usage string
	Func  cmdHandler
}

var commands = map[string]command{
	// list commands
	"install":   command{"Install new versions or switch.", "[0.8.8 version of terraform]", tfinstall.Run},
	"uninstall": command{"Clean cache (tmp files).", "[-a all], [-v version specific]", tfuninstall.Run},
	"list":      command{"List online or offline version of terraform.", "[-on list online], [-off list local]", tflist.Run},
	"cloud":     command{"Action cloud (Beta)", "[--aws test AWS]", tfcloud.Run},
}

var (
	// Errors
	version string
	err     error
)

func init() {
	// version app
	version = "0.0.3"
}

type cmdHandler func([]string) error

func doHelp() error {

	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}

	fmt.Printf("tfversion v%s\n\n", version)
	fmt.Printf("Usage: tfversion <command> [args]\n\n")
	fmt.Printf("Common commands:\n")

	for _, k := range keys {
		fmt.Printf("%10s: %s\n", k, commands[k].desc)
		fmt.Printf("\tUsage: %s %s\n", k, commands[k].usage)
	}

	fmt.Printf("      help: Show this help message\n\n")

	// Show if the last version
	lastrelease, release := tfgithub.Lastversion(version)
	if !lastrelease && release != nil {
		tferror.Run(2, fmt.Sprintf("Your version of tfversion is out of date !\nThe latest version is %s (%s)", *release.TagName, *release.HTMLURL))
	}

	return nil
}

func main() {

	if len(os.Args) < 2 || os.Args[1] == "help" {
		err = doHelp()
		tferror.Panic(err)
		return
	}

	if cmd, ok := commands[os.Args[1]]; ok {
		if err = cmd.Func(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Unknown command '%s'\n", os.Args[1])
		fmt.Fprintf(os.Stderr, "Usage: %s command command_arguments\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\tUse help command to list available commands\n")
		fmt.Fprintf(os.Stderr, "\tUse command help to get commands accepting options\n")
	}
}
