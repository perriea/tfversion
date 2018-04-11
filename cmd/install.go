package cmd

import (
	"fmt"

	"github.com/perriea/tfversion/terraform"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install new version",
	Long:  `Install new version or switch.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		terraform.InitFolder()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			terraform.Install(args[0], true)
		} else {
			fmt.Printf("\033[1;31mNone version specified\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
