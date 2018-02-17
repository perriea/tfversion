package cmd

import (
	"fmt"

	"github.com/perriea/tfversion/errors"
	"github.com/perriea/tfversion/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version installed of switcher Terraform",
	Long:  `Version installed of switcher Terraform`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tfversion v%s\n\n", version.String())

		// Show if the last version
		test, release := version.LastVersion()
		errors.Debug(0, "Version", fmt.Sprintf("test: %t", test))
		if !test && release != nil {
			fmt.Printf("Your version is out of date ! The latest version is %s.\nYou can update by downloading from Github (%s).", *release.TagName, *release.HTMLURL)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
