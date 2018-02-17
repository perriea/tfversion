package cmd

import (
	"github.com/perriea/tfversion/terraform"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install new versions or switch",
	Long:  `Install new versions or switch.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		terraform.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		terraform.Install(args[0])
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
