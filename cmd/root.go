package cmd

import (
	"fmt"
	"os"

	"github.com/perriea/tfversion/version"
	"github.com/spf13/cobra"
)

var (
	quiet   bool
	all     bool
	cfgFile string
	err     error
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tfversion",
	Short: "tfversion v" + version.String() + " - Switcher Terraform",
	Long:  "tfversion v" + version.String() + " - Switcher Terraform",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err = rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
