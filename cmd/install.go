package cmd

import (
	"github.com/perriea/tfversion/terraform"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Install a new version",
	Long:  `Install a new version or switch.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if err = terraform.Install(args[0], quiet); err != nil {
				panic(err)
			}
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Do not show information messages")
}
