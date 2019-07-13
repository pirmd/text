package app

import (
	"bytes"
	"testing"

	"github.com/pirmd/cli/style"
	"github.com/pirmd/verify"
)

func TestManpage(t *testing.T) {
	testApp := buildTestApp()

	out := new(bytes.Buffer)

	//Ensure that manpage date will remain the same whenever the test is run
	manDate = "2019-07-12"
	PrintManpage(out, testApp, style.Man)

	verify.MatchGolden(t, out.String(), "Manpage message is incorrectly formatted")
}
