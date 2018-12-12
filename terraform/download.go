package terraform

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

// Download : Launch download
func (release Release) Download(quiet bool) error {
	// Formulation URL Terraform Website
	url := fmt.Sprintf(release.Repository, release.Version, release.Version, runtime.GOOS, runtime.GOARCH)

	// Request GET URL
	resp, err := release.HTTPclient.Get("https://" + url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Verify code equal 200
	if (err == nil) && (resp.StatusCode == 200) {
		ShowMessage("\033[1;32mStart download ...\n", quiet)
		fileUnzip, err := os.Create(fmt.Sprintf("%s%sterraform-%s.zip", release.Home, folders["tmp"], release.Version))
		if err != nil {
			return err
		}
		defer fileUnzip.Close()

		// Copy reponse in file
		_, err = io.Copy(fileUnzip, resp.Body)
		if err != nil {
			return err
		}
		return nil
	}
	ShowMessage("\033[1;31m[ERROR] Failed, this version doesn't exist !", quiet)
	return nil
}
