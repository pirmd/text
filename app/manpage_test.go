package app

import (
	"bytes"
	"testing"

	"github.com/pirmd/verify"
	"github.com/pirmd/cli/style"
)

func TestManpage(t *testing.T) {
	testApp := buildTestApp()

	out := new(bytes.Buffer)
	PrintManpage(out, testApp, style.Mandoc)
	verify.MatchGolden(t, out.String(), "Manpage message is incorrectly formatted")
}
