package cmd

import (
	"github.com/perriea/tfversion/terraform"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [version]",
	Short: "Remove local version of Terraform",
	Long:  `Remove local version of Terraform`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		release = terraform.Release{
			Home: home,
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check argument number
		if len(args) > 0 && !all {
			// affect value version
			release.Version = args[0]
			if release.Regex() {
				return release.UnInstall(quiet)
			}
		} else if len(args) == 0 && all {
			return release.UnInstallAll(quiet)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolVarP(&all, "all", "a", false, "Remove all version of Terraform")
	removeCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Do not show information messages")
}
