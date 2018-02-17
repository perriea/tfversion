package terraform

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/perriea/tfversion/errors"
)

// Download : Launch download
func Download(version string) bool {

	fmt.Printf("Attempting to download version: %s\n", version)
	if _, err := os.Stat(filepath.Join(home, tfVersionHomePath, fmt.Sprintf("terraform-%s.zip", version))); os.IsNotExist(err) {
		match, err := regexp.MatchString("[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta)[0-9]+)?", version)
		errors.Panic(err)

		if match {
			// Formulation URL Terraform Website
			url := fmt.Sprintf(urlHashicorpRelease, version, version, runtime.GOOS, runtime.GOARCH)

			// Request GET URL
			resp, err = client.Get(url)
			errors.Panic(err)
			defer resp.Body.Close()

			// Verify code equal 200
			if (err == nil) && (resp.StatusCode == 200) {
				fmt.Printf("\033[1;32mStart download ...\n")
				pathTF := fmt.Sprintf("%s%sterraform-%s.zip", home, tfVersionHomePath, version)

				fileUnzip, err := os.Create(pathTF)
				errors.Panic(err)
				defer fileUnzip.Close()

				// Copy reponse in file
				_, err = io.Copy(fileUnzip, resp.Body)
				errors.Panic(err)
				return true
			}
			fmt.Printf("\033[1;31m[ERROR] Failed, this version doesn't exist !\n")
			return false
		}
		fmt.Printf("\033[1;31m[ERROR] Failed, the format version is not correct ...\n")
		return false
	}
	fmt.Printf("\033[1;34mAlready in cache ...\n")
	return true
}
