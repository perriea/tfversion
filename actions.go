package main

import (
	"fmt"
	"runtime"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/github"
)

var (
	version string
)

func init() {
	version = "0.0.2"
}

func ShowVersion() {

	fmt.Printf("Family OS: %s\n", runtime.GOOS)
	fmt.Printf("Arch processor: %s\n\n", runtime.GOARCH)

	fmt.Printf("Version: %s\n", version)

	// Show if the last version
	lastrelease, release := tfgithub.Lastversion(version)
	if !lastrelease {
		tferror.Run(2, fmt.Sprintf("Now, a new version is available : %s (%s)", *release.TagName, *release.HTMLURL))
	}
}
