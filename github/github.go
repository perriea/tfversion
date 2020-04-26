package github

import (
	"context"

	"github.com/google/go-github/github"
	pversion "github.com/hashicorp/go-version"
)

// LastVersion : Check last version of package
func LastVersion(version string) (bool, *github.RepositoryRelease) {
	var (
		semVer                   = pversion.Must(pversion.NewVersion(version))
		ctx      context.Context = context.Background()
		client   *github.Client  = github.NewClient(nil)
		releases []*github.RepositoryRelease
		err      error
	)

	if releases, _, err = client.Repositories.ListReleases(ctx, "perriea", "tfversion", nil); err != nil {
		panic(err)
	}

	if len(releases) > 0 {
		lastRelease, _ := pversion.NewVersion(*releases[0].TagName)
		if semVer.LessThan(lastRelease) {
			return true, releases[0]
		}
	}

	return false, nil
}
