package terraform

import (
	"crypto/tls"
	"net/http"

	homedir "github.com/mitchellh/go-homedir"
)

const urlHashicorp = "https://releases.hashicorp.com/terraform/"
const urlHashicorpRelease = urlHashicorp + "%s/terraform_%s_%s_%s.zip"
const tfVersionHomePath = "/.tfversion/tmp/"
const tfVersionHomeBin = "/.tfversion/bin/"

var (
	// homedir
	home string

	// HTTP request
	transport *http.Transport
	client    *http.Client
	resp      *http.Response

	version string

	// Errors
	err error
)

func init() {
	// Dont check certificate SSL + new path
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: transport}

	home, err = homedir.Dir()
	if err != nil {
		panic(err)
	}
}
