package cmd

import (
	"fmt"
	"runtime"

	"github.com/perriea/tfversion/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version installed",
	Long:  `Version installed of switcher Terraform`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tfversion v%s\n\n", version.String())

		// Show if the last version
		latest, release := version.LastVersion()
		if latest && release != nil {
			switch runtime.GOOS {
			case "darwin":
				fmt.Printf("Your version is out of date ! The latest version is %s.\nYou can update with brew.", *release.TagName)
			case "linux":
				fmt.Printf("Your version is out of date ! The latest version is %s.\nYou can update with snap (Ubuntu) or download from Github (%s).", *release.TagName, *release.HTMLURL)
			default:
				fmt.Printf("You are strange man ! :D")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
