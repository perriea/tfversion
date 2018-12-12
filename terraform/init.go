package terraform

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var (
	folders map[string]string
	err     error
)

// Release struct : information switch release
type Release struct {
	Home         string
	Version      string
	Repository   string
	HTTPclient   *http.Client
	HTTPResponse *http.Response
}

// InitFolder : Create folders (init)
func (release Release) InitFolder() error {
	for _, folder := range folders {
		err := os.MkdirAll(filepath.Join(release.Home, folder), os.FileMode(0755))
		if err != nil {
			return err
		}
	}

	return nil
}

// Message : Quiet mode
func Message(message string, quiet bool) {
	if !quiet {
		fmt.Println(message)
	}
}

func init() {
	folders = map[string]string{
		"bin": "/.tfversion/bin/",
		"tmp": "/.tfversion/tmp/",
	}
}
