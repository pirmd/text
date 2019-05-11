package style

import (
	"testing"
	"github.com/pirmd/verify"
)

func TestEscaping(t *testing.T) {
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
		verify.EqualString(t, got, tc.out, "escape mandoc")
	}
}
