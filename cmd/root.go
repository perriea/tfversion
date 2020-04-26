package cmd

import (
	"fmt"
	"os"

	"github.com/perriea/tfversion/terraform"
	"github.com/perriea/tfversion/version"
	"github.com/spf13/cobra"
)

var (
	r             *terraform.Release = terraform.NewRelease()
	home, cfgFile string
	quiet, all    bool
	err           error
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tfversion",
	Short: fmt.Sprintf("tfversion v%s - Switcher Terraform", version.String()),
	Long:  fmt.Sprintf("tfversion v%s - Switcher Terraform", version.String()),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err = rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
