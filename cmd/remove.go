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
		if vTerraform != "" {
			if err = terraform.UnInstallOne(vTerraform); err != nil {
				panic(err)
			}
		} else if all {
			if err = terraform.UnInstallAll(); err != nil {
				panic(err)
			}
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(&vTerraform, "version", "v", "", "Remove one version of Terraform")
	removeCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "Remove all version of Terraform")
}
