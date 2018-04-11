package cmd

import (
	"github.com/perriea/tfversion/terraform"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all available versions",
	Long:  `List of all available versions`,
	Run: func(cmd *cobra.Command, args []string) {
		terraform.ListOnline()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
