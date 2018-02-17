package version

import (
	"fmt"
	"testing"
)

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
