package tfgithub

import (
	"fmt"
	"testing"
)

// ListReleases : get all version published
func TestListReleases(t *testing.T) {
	releases, err = ListReleases("perriea", "tfversion")
	if err != nil {
		t.Fatalf("Error request : %s", err)
	} else {
		fmt.Printf("Versions available : %s\n", *releases[0].TagName)
	}
}

// LastVersion : Check last version of package
func TestLastVersion(t *testing.T) {
	var (
		test bool
	)

	test, _ = LastVersion("0.1.3")

	if releases == nil {
		t.Fatalf("Error release is null")
	} else {
		fmt.Printf("Last version avalaible ? : %t\n\n", test)
	}
}
