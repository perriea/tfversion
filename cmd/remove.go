package cmd

import (
	"github.com/perriea/tfversion/errors"
	"github.com/perriea/tfversion/terraform"
	"github.com/spf13/cobra"
)

var (
	Version string
	err     error
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove local version of Terraform",
	Long:  `Remove local version of Terraform`,
	Run: func(cmd *cobra.Command, args []string) {
		err = terraform.UninstallOne(args[0])
		errors.Panic(err)

		// err = terraform.UninstallAll(filepath.Join(usr.HomeDir, "/.tfversion/bin/"))
		// errors.Panic(err)
		// fmt.Printf("All files are deleted !\n")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(&Version, "version", "v", "", "Remove one version of Terraform")
	// removeCmd.Flags().BoolP("all", "a", false, "Remove all version of Terraform")
}
