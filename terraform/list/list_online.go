package tflist

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"regexp"

	"github.com/perriea/tfversion/error"
	"github.com/perriea/tfversion/system/network"
)

var (
	url_tf     string
	cleaned    []string
	available  []string
	tfversions []string
	errNetwork bool
	transport  *http.Transport
	client     *http.Client
)

func init() {
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: transport}

	url_tf = "https://releases.hashicorp.com/terraform/"
}

func stringInSlice(str string, list []string) bool {

	for _, v := range list {
		if v == str {
			return true
		}
	}

	return false
}

func ListOn() {
	errNetwork = tfnetwork.Run("releases.hashicorp.com:80", 3, false)

	if errNetwork {
		resp, err := client.Get(url_tf)
		tferror.Panic(err)
		defer resp.Body.Close()

		// Verify code equal 200
		if (err == nil) && (resp.StatusCode == 200) {

			r, err := regexp.Compile("[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta)[0-9]+)?")
			tferror.Panic(err)

			// Convert byte to string
			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			newStr := buf.String()

			fmt.Printf("\033[1;34m[INFO] Versions availables of terraform :\n")
			tfversions = r.FindAllString(newStr, -1)

			// Clean doublon
			for _, value := range tfversions {
				if !stringInSlice(value, cleaned) {
					cleaned = append(cleaned, value)
				}
			}

			// Show versions
			showList(cleaned, "0")
		}
	} else {
		fmt.Printf("\033[1;31m[ERROR] No internet connection ...\n")
	}
}
