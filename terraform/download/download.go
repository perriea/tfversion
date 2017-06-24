package tfdownload

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/system/network"
)

var (
	// URL and Path (file)
	doPathTF  string
	urlTF     string
	pathTF    string
	fileUnzip *os.File

	// check if format version is *.*.* and more ...
	match bool

	// HTTP request
	transport *http.Transport
	client    *http.Client
	resp      *http.Response

	// Errors
	err error
)

const host = "releases.hashicorp.com"
const tfpath = "/.tfversion/tmp/"

func init() {
	// Dont check certificate SSL + new path
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: transport}
	urlTF = "https://%s/terraform/%s/terraform_%s_%s_%s.zip"

	usr, err := user.Current()
	tferror.Panic(err)

	doPathTF = filepath.Join(usr.HomeDir, tfpath, "terraform-%s.zip")
}

// Run : Launch download
func Run(version string) bool {

	if _, err := os.Stat(fmt.Sprintf(doPathTF, version)); os.IsNotExist(err) {

		match, err := regexp.MatchString("[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta)[0-9]+)?", version)
		tferror.Panic(err)

		if match {

			errNetwork := tfnetwork.Run(host+":443", 3, false)
			if errNetwork {
				// Formulation URL Terraform Website
				fmt.Printf("Attempting to download version: %s\n", version)
				urlTF = fmt.Sprintf(urlTF, host, version, version, runtime.GOOS, runtime.GOARCH)

				// Request GET URL
				resp, err = client.Get(urlTF)
				tferror.Panic(err)
				defer resp.Body.Close()

				// Verify code equal 200
				if (err == nil) && (resp.StatusCode == 200) {

					fmt.Printf("\033[1;32mStart download ...\n")
					pathTF = fmt.Sprintf(doPathTF, version)
					fileUnzip, err = os.Create(pathTF)
					tferror.Panic(err)
					defer fileUnzip.Close()

					// Copy reponse in file
					_, err = io.Copy(fileUnzip, resp.Body)
					tferror.Panic(err)

					return true

				} else {
					fmt.Printf("\033[1;31m[ERROR] Download impossible, this version doesn't exist !\n")
					return false
				}
			} else {
				fmt.Printf("\033[1;31m[ERROR] No internet connection ...\n")
				return false
			}

		} else {
			fmt.Printf("\033[1;31m[ERROR] The version format is not correct ...\n")
			return false
		}

	} else {
		fmt.Printf("\033[1;34mAlready in cache ...\n")
		return true
	}
}
