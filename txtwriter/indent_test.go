package txtwriter

import (
	"strings"
	"testing"
)

func TestIndent(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "Hello!",
			want: "    Hello!",
		},
		{
			in:   "Hello ladies and gentlemen!",
			want: "    Hello ladies\n    and\n    gentlemen!",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(16)
		tstwriter.Indent()

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%s'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}

func TestBlockIndent(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "Hello!",
			want: "    Hello!",
		},
		{
			in:   "Hello ladies and gentlemen!",
			want: "    Hello\n    ladies\n    and\n    gentleme\n    n!",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(16)
		tstwriter.BlockIndent()
		tstwriter.Indent()

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%s'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}
