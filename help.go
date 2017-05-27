package main

import (
	"fmt"
	"os"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/github"
)

func doHelp() error {

	var version string

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

func doUnkw() {

	fmt.Fprintf(os.Stderr, "Unknown command '%s'\n", os.Args[1])
	fmt.Fprintf(os.Stderr, "Usage: %s command command_arguments\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tUse help command to list available commands\n")
	fmt.Fprintf(os.Stderr, "\tUse command help to get commands accepting options\n")
}
