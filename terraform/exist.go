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
func (r *Release) Regex() bool {
	var (
		rmatch string = "[0-9]+\\.[0-9]+\\.[0-9]+(-(rc|beta|alpha)[0-9]+)?"
		match  bool   = false
		err    error
	)

	if match, err = regexp.MatchString(rmatch, r.Version); err != nil {
		return false
	}

	return match
}

// localExist version zipped offline
func (r *Release) localExist() error {
	var (
		version string = fmt.Sprintf("terraform-%s.zip", r.Version)
		err     error
	)

	if _, err = os.Stat(filepath.Join(r.Home, PathTmp.toString(), version)); !os.IsNotExist(err) {
		fmt.Println("Already in cache ...")
		return err
	}

	return nil
}

// remoteExist version zipped online
func (r *Release) remoteExist() error {
	var (
		url  string = fmt.Sprintf(PathTerraform.toString(), r.Version, r.Version, runtime.GOOS, runtime.GOARCH)
		resp *http.Response
		err  error
	)

	if resp, err = r.HTTPclient.Get(url); err != nil {
		return err
	}
	defer resp.Body.Close()

	// Verify code equal 200
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return nil
}
