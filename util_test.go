package text

import (
	"testing"
)

func TestVisualLen(t *testing.T) {
	testCases := []struct {
		in  string
		out int
	}{
		{"abc", 3},
		{"敬具", 2},
		{"a\x00bc", 3},
		{"\x1b[35ma\x00bc\x1b[0m", 3},
	}

	for _, tc := range testCases {
		got := visualLen(tc.in)
		if got != tc.out {
			t.Errorf("Length of '%s' failed: got %d, wanted %d", string(tc.in), got, tc.out)
		}
	}
}

func TestVisualTruncate(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out string
	}{
		{"Coucou", 9, "Coucou"},
		{"This is a long sentence", 9, "This is a"},
		{"This \x1b[34mis\x1b[0m a long sentence in color", 9, "This \x1b[34mis\x1b[0m a\x1b[0m"},
		{"This \x1b[34mis a long sentence in\x1b[0m color", 9, "This \x1b[34mis a\x1b[0m"},
	}

	for _, tc := range testCases {
		got := visualTruncate(tc.in, tc.sz)
		if got != tc.out {
			t.Errorf("visual Truncate failed for '%s'.\nWanted: %s\nGot   : %s\n", tc.in, tc.out, got)
		}
	}
}

func TestVisualPad(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out string
	}{
		{"Coucou", 9, "Coucou   "},
		{"\x1b[34mCoucou\x1b[0m", 9, "\x1b[34mCoucou\x1b[0m   "},
		{"Coucou Coucou", 9, "Coucou Coucou"},
	}

	for _, tc := range testCases {
		got := visualPad(tc.in, tc.sz, ' ')
		if got != tc.out {
			t.Errorf("visual Padding failed for '%s'.\nWanted: %s\nGot   : %s\n", tc.in, tc.out, got)
		}
	}
}
