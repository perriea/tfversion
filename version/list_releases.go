package version

import (
	"context"

	"github.com/google/go-github/github"
	version "github.com/hashicorp/go-version"
)

// LastVersion : Check last version of package
func LastVersion() (bool, *github.RepositoryRelease) {
	var (
		client   *github.Client
		ctx      context.Context
		releases []*github.RepositoryRelease
	)

	client = github.NewClient(nil)
	ctx = context.Background()

	releases, _, err := client.Repositories.ListReleases(ctx, "perriea", "tfversion", nil)
	if err != nil {
		panic(err)
	}

	if len(releases) > 0 {
		lastRelease, _ := version.NewVersion(*releases[0].TagName)
		if SemVer.LessThan(lastRelease) {
			return true, releases[0]
		}
	}

	return false, nil
}
