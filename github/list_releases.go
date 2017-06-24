package tfgithub

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/perriea/tfversion/system/network"
)

var (
	releases   []*github.RepositoryRelease
	ctx        context.Context
	client     *github.Client
	errNetwork bool
	err        error
)

func init() {
	ctx = context.Background()
	client = github.NewClient(nil)
	errNetwork = false
}

// ListReleases : get all version published
func ListReleases(username string, repository string) ([]*github.RepositoryRelease, error) {
	releases, _, err := client.Repositories.ListReleases(ctx, username, repository, nil)
	return releases, err
}

// LastVersion : Check last version of package
func LastVersion(version string) (bool, *github.RepositoryRelease) {

	// check if the internet connection is active
	errNetwork = tfnetwork.Run("github.com:80", 3, false)
	if errNetwork {
		releases, _ = ListReleases("perriea", "tfversion")

		if *releases[0].TagName == version {
			return true, releases[0]
		}
		return false, releases[0]
	}
	return false, nil
}
