package txtwriter

import (
	"strings"
	"testing"
)

func TestWriteWithoutWrap(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "Hello world!",
			want: "Hello world!",
		},
		{
			in:   "Hello world!\n",
			want: "Hello world!\n",
		},
		{
			in:   "Hello ladies and gentlemen!",
			want: "Hello ladies and gentlemen!",
		},
		{
			in:   "Supercalifragilisticexpialidocious\n\nChim Chim Cher-ee",
			want: "Supercalifragilisticexpialidocious\n\nChim Chim Cher-ee",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(0)

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%s'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}
