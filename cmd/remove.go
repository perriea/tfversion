package cmd

import (
	"github.com/perriea/tfversion/terraform"
	"github.com/spf13/cobra"
)

var (
	vTerraform string
	all        bool
	err        error
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove local version of Terraform",
	Long:  `Remove local version of Terraform`,
	Run: func(cmd *cobra.Command, args []string) {
		if (vTerraform != "" && vTerraform != "noversion") && !all {
			if err = terraform.UnInstall(vTerraform, false); err != nil {
				panic(err)
			}
		} else if all {
			if err = terraform.UnInstall("all", false); err != nil {
				panic(err)
			}
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVar(&vTerraform, "version", "noversion", "Remove one version of Terraform")
	removeCmd.Flags().BoolVar(&all, "all", false, "Remove all version of Terraform")
}
