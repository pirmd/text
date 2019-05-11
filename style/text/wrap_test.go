package text

import (
	"testing"
)

func TestWrap(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out string
	}{
		{"Coucou", 10, "Coucou"},
		{"Coucou", 6, "Coucou"},
		{"Coucou ", 6, "Coucou"},
		{"Coucou\n", 8, "Coucou\n"},
		{"This is a long sentence", 10, "This is a \nlong \nsentence"},
		{"This \x1b[34mis\x1b[0m a long sentence", 10, "This \x1b[34mis\x1b[0m a \nlong \nsentence"},
		{"This \x1b[34mis a long sentence\x1b[0m", 10, "This \x1b[34mis a \nlong \nsentence\x1b[0m"},
	}

	for _, tc := range testCases {
		got := Wrap(tc.in, tc.sz)
		if got != tc.out {
			t.Errorf("Wrap failed for '%s'.\nWanted:\n'%s'\nGot   :\n'%s'\n", tc.in, tc.out, got)
		}
	}
}

func TestWrapAndIndent(t *testing.T) {
	testCase := "This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped because it hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached."
	wanted := `    This very long and detailed sentence
    is here to demonstrate that list can
    be formatted and wrapped because it 
    hopefully be so long that it will 
    not fulfill the maximum number of 
    authorized char per lines is 
    reached.`

	got := Tab(testCase, []byte("    "), 40)
	if got != wanted {
		t.Errorf("Indenting or/and wrapping failed.\nWanted:\n%s\nGot   :\n%s\n", wanted, got)
	}
}
