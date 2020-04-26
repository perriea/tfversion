package terraform

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	"github.com/mitchellh/go-homedir"
)

func (r *Release) Run(args []string, quiet bool) error {
	var (
		nArg int = len(args)
		err  error
	)

	// Check argument number
	if nArg > 0 {
		// affect value version
		r.Version = args[0]

		// Check this value
		if r.Regex() {
			// Check if r is stocked in localy & remotely
			if err = r.localExist(); err != nil {
				return r.install(quiet)
			} else if err = r.remoteExist(); err == nil {
				if err = r.download(quiet); err != nil {
					return err
				}

				if err := r.install(quiet); err != nil {
					return err
				}

				return nil
			}

			return nil
		}

		return errors.New("format version invalid")
	}

	return errors.New("Argument(s) missing")
}

// NewRelease Client
func NewRelease() *Release {
	var (
		home string
		err  error
	)

	if home, err = homedir.Dir(); err != nil {
		panic(err)
	}

	return &Release{
		Home: home,
		HTTPclient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: time.Duration(60 * time.Second),
		},
	}
}
