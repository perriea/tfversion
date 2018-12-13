package cmd

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	"github.com/perriea/tfversion/terraform"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Install a new version",
	Long:  `Install a new version or switch.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		release = terraform.Release{
			Home:       home,
			HTTPclient: &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}, Timeout: time.Duration(5 * time.Second)},
			Repository: "releases.hashicorp.com/terraform/%s/terraform_%s_%s_%s.zip",
		}

		return release.InitFolder()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check argument number
		if len(args) > 0 {
			// affect value version
			release.Version = args[0]

			// Check this value
			if release.Regex() {
				// Check if release is stocked in localy & remotely
				if release.LocalExist() {
					return release.Install(quiet)
				} else if release.RemoteExist() {
					if err := release.Download(quiet); err != nil {
						return err
					}

					if err := release.Install(quiet); err != nil {
						return err
					}

					return nil
				}

				return nil
			}
		}

		return errors.New("Argument(s) missing")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Do not show information messages")
}
