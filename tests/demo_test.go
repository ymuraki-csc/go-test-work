package tests

import (
	demopkg "example.com/demo-pkg"
	gocmp "github.com/google/go-cmp/cmp"
	"testing"
)

func TestDemoPkg(t *testing.T) {
	expected := &demopkg.Demo{Str: "expected"}
	actual := demopkg.New("actual")
	if gocmp.Diff(expected, actual) != "" {
		t.Errorf(gocmp.Diff(expected, actual))
	}
}
