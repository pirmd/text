package txtwriter

import (
	"strings"
	"testing"
)

func TestPrefix(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "Hello world!",
			want: "    - Hello\n      world!",
		},
		{
			in:   "Hello ladies and gentlemen!",
			want: "    - Hello\n      ladies\n      and\n      gentle\n      men!",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(12)
		tstwriter.Indent()
		tstwriter.SetPrefix("- ", "  ")

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%s'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}

func TestPrefixMultiwrite(t *testing.T) {
	testCases := []struct {
		in   []string
		want string
	}{
		{
			in:   []string{"Hello world!\n", "Hello ladies and gentlemen!"},
			want: "    - Hello\n      world!\n    - Hello\n      ladies\n      and\n      gentle\n      men!",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(12)
		tstwriter.Indent()
		tstwriter.SetPrefix("- ", "  ")

		for _, tst := range tc.in {
			tstwriter.Write([]byte(tst))
			tstwriter.ResetPrefix()
		}
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%v'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}
