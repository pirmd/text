package style

import (
	"github.com/pirmd/verify"
	"testing"
)

func TestMandocEscaping(t *testing.T) {
	st := NewMan()

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
		got := st.Escape(tc.in)
		verify.EqualString(t, got, tc.out, "escape man (second pass): '%s'", tc.in)

		gotgot := st.Escape(got)
		verify.EqualString(t, gotgot, tc.out, "escape man (second pass): '%s'", got)
	}
}
