package style

import (
	"github.com/pirmd/verify"
	"testing"
)

func TestMandocEscaping(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"Test -- test is fun", "Test \\-\\- test is fun"},
		{"Test \\-\\- test is fun", "Test \\-\\- test is fun"},
		{"Test -- test is \\fBfun\\fP", "Test \\-\\- test is \\fBfun\\fP"},
		{"Test a backslash \\\\", "Test a backslash \\\\"},
	}

	for _, tc := range testCases {
		got := escapeMandoc(tc.in)
		verify.EqualString(t, got, tc.out, "escape mandoc: '%s'", tc.in)

		gotgot := escapeMandoc(got)
		verify.EqualString(t, gotgot, tc.out, "escape mandoc: '%s'", got)
	}
}

func TestMdocEscaping(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{".Bl -bullet\n.It\nTesting is ~fun~\n.El", ".Bl -bullet\n.It\nTesting is \\~fun\\~\n.El"},
	}

	for _, tc := range testCases {
		got := escapeMdoc(tc.in)
		verify.EqualString(t, got, tc.out, "escape mdoc: '%s'", tc.in)

		gotgot := escapeMdoc(got)
		verify.EqualString(t, gotgot, tc.out, "escape mdoc: '%s'", got)
	}
}
