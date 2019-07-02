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
	testCases := []struct {
		inP, inT string
		out      string
	}{
		{
			"    ",
			"Test",
			"    Test",
		},
		{
			"  > ",
			"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped because it hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached.",
			"  > This very long and detailed sentence\n  > is here to demonstrate that list can\n  > be formatted and wrapped because it \n  > hopefully be so long that it will \n  > not fulfill the maximum number of \n  > authorized char per lines is \n  > reached.",
		},
		{
			"  \x1b[1m>\x1b[22m ",
			"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped because it hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached.",
			"  \x1b[1m>\x1b[22m This very long and detailed sentence\n  \x1b[1m>\x1b[22m is here to demonstrate that list can\n  \x1b[1m>\x1b[22m be formatted and wrapped because it \n  \x1b[1m>\x1b[22m hopefully be so long that it will \n  \x1b[1m>\x1b[22m not fulfill the maximum number of \n  \x1b[1m>\x1b[22m authorized char per lines is \n  \x1b[1m>\x1b[22m reached.",
		},
	}

	for _, tc := range testCases {
		got := Tab(tc.inT, tc.inP, 40)
		if got != tc.out {
			t.Errorf("Indenting or/and wrapping failed.\nWanted:\n%s\nGot   :\n%s\n", tc.out, got)
		}
	}
}
