package main

import (
	"fmt"
	"os"

	"github.com/perriea/tfversion/cloud"
	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/terraform/install"
	"github.com/perriea/tfversion/terraform/list"
	"github.com/perriea/tfversion/terraform/uninstall"
)

type command struct {
	desc  string
	usage string
	Func  cmdHandler
}

var (
	// Errors
	commands map[string]command
	version  string
	err      error
)

type cmdHandler func([]string) error

func init() {
	// list commands
	commands = map[string]command{
		"install":   command{"Install new versions or switch.", "[0.8.8 version of terraform]", tfinstall.Run},
		"uninstall": command{"Clean cache (tmp files).", "[-a all], [-v version specific]", tfuninstall.Run},
		"list":      command{"List online or offline version of terraform.", "[-on list online], [-off list local]", tflist.Run},
		"cloud":     command{"Action cloud (Beta)", "[--aws test AWS]", tfcloud.Run},
	}
}

func main() {
	// Check if the command is helper
	if len(os.Args) < 2 || os.Args[1] == "help" {
		err = DoHelp()
		tferror.Panic(err)
		return
	}

	// Look command
	if cmd, ok := commands[os.Args[1]]; ok {
		if err = cmd.Func(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	} else {
		CmdUnknown()
	}
}
