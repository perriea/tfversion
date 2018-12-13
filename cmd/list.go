package cmd

import (
	"crypto/tls"
	"net/http"

	"github.com/perriea/tfversion/terraform"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of available versions",
	Long:  `List of available versions`,
	PreRun: func(cmd *cobra.Command, args []string) {
		release = terraform.Release{
			Home:       home,
			HTTPclient: &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
			Repository: "releases.hashicorp.com/terraform/",
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err = release.ListOnline(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
