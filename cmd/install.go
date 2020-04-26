package cmd

import (
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Install a new version",
	Long:  `Install a new version or switch.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return r.InitFolder()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return r.Run(args, quiet)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Do not show information messages")
}
