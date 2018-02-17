package cmd

import (
	"github.com/perriea/tfversion/terraform"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of terraform versions",
	Long:  `List of terraform versions`,
	Run: func(cmd *cobra.Command, args []string) {
		terraform.ListOnline()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
