package tffiles

import (
	"testing"
)

func TestCreateVersioning(t *testing.T) {
	err = CreateVersioning("0.9.6")

	if err != nil {
		t.Fatalf("Error creation file : %s", err)
	}
}
