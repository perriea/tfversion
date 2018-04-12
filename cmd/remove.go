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
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && !all {
			if err = terraform.UnInstall(args[0], quiet); err != nil {
				panic(err)
			}
		} else if all {
			if err = terraform.UnInstall("all", quiet); err != nil {
				panic(err)
			}
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolVarP(&all, "all", "a", false, "Remove all version of Terraform")
	removeCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Do not show information messages")
}
