package terraform

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

// Download : Launch download
func Download(version string, quiet bool) (bool, error) {
	ShowMessage(fmt.Sprintf("\033[1;36mAttempting to download version (%s) ...\n", version), quiet)
	if _, err := os.Stat(filepath.Join(home, tfVersionHomePath, fmt.Sprintf("terraform-%s.zip", version))); os.IsNotExist(err) {
		match, err := regexp.MatchString("[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta)[0-9]+)?", version)
		if err != nil {
			return false, err
		}

		if match {
			// Formulation URL Terraform Website
			url := fmt.Sprintf(urlHashicorpRelease, version, version, runtime.GOOS, runtime.GOARCH)

			// Request GET URL
			resp, err = client.Get(url)
			if err != nil {
				return false, err
			}
			defer resp.Body.Close()

			// Verify code equal 200
			if (err == nil) && (resp.StatusCode == 200) {
				ShowMessage("\033[1;32mStart download ...\n", quiet)
				fileUnzip, err := os.Create(fmt.Sprintf("%s%sterraform-%s.zip", home, tfVersionHomePath, version))
				if err != nil {
					return false, err
				}
				defer fileUnzip.Close()

				// Copy reponse in file
				_, err = io.Copy(fileUnzip, resp.Body)
				if err != nil {
					return false, err
				}
				return true, nil
			}
			ShowMessage("\033[1;31m[ERROR] Failed, this version doesn't exist !", quiet)
			return false, nil
		}
		ShowMessage("\033[1;31m[ERROR] Failed, the format version is not correct ...", quiet)
		return false, nil
	}
	ShowMessage("\033[1;34mAlready in cache ...", quiet)
	return true, nil
}
