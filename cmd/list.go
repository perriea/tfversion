package cmd

import (
	"github.com/perriea/tfversion/terraform"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of available versions",
	Long:  `List of available versions`,
	Run: func(cmd *cobra.Command, args []string) {
		if err = terraform.ListOnline(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
