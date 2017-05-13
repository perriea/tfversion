package tfgithub

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/perriea/tfversion/error"
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

func ListReleases() []*github.RepositoryRelease {

	releases, _, err := client.Repositories.ListReleases(ctx, "perriea", "tfversion", nil)
	tferror.Panic(err)

	return releases
}

func Lastversion(version string) (bool, *github.RepositoryRelease) {

	// check if the internet connection is active
	errNetwork = tfnetwork.Run()
	if errNetwork {
		releases = ListReleases()

		if *releases[0].TagName == version {
			return true, releases[0]
		}
		return false, releases[0]
	}
	return false, nil
}
