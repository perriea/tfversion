package terraform

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

// Regex version submited
func (release Release) Regex() bool {
	match, err := regexp.MatchString("[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta|alpha)[0-9]+)?", release.Version)
	if err != nil {
		return false
	}

	return match
}

// LocalExist version zipped offline
func (release Release) LocalExist() bool {
	if _, err := os.Stat(filepath.Join(release.Home, folders["tmp"], fmt.Sprintf("terraform-%s.zip", release.Version))); !os.IsNotExist(err) {
		fmt.Println("Already in cache ...")
		return true
	}

	return false
}

// RemoteExist version zipped online
func (release Release) RemoteExist() bool {
	url := fmt.Sprintf(release.Repository, release.Version, release.Version, runtime.GOOS, runtime.GOARCH)
	resp, err := release.HTTPclient.Get("https://" + url)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	// Verify code equal 200
	if (err == nil) && (resp.StatusCode == http.StatusOK) {
		return true
	}

	return false
}
