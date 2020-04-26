package main

import (
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [version]",
	Short: "Remove local version of Terraform",
	Long:  `Remove local version of Terraform`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check argument number
		if len(args) > 0 && !all {
			// affect value version
			r.Version = args[0]
			if r.Regex() {
				return r.UnInstall(quiet)
			}
		} else if len(args) == 0 && all {
			return r.UnInstallAll(quiet)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolVarP(&all, "all", "a", false, "Remove all version of Terraform")
	removeCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Do not show information messages")
}
