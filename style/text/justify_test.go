package text

import (
	"testing"
)

func TestTruncate(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out string
	}{
		{"Coucou", 10, "Coucou"},
		{"This is a long sentence", 10, "This is a…"},
		{"This \x1b[34mis\x1b[0m a long sentence in color", 10, "This \x1b[34mis\x1b[0m a\x1b[0m…"},
		{"This \x1b[34mis a long sentence in\x1b[0m color", 10, "This \x1b[34mis a\x1b[0m…"},
	}

	for _, tc := range testCases {
		got := Truncate(tc.in, tc.sz)
		if got != tc.out {
			t.Errorf("Truncate failed for '%s'.\nWanted: %s\nGot   : %s\n", tc.in, tc.out, got)
		}
	}
}

func TestExactSize(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out string
	}{
		{"Coucou", 10, "Coucou    "},
		{"This is a long sentence", 10, "This is a…"},
		{"This \x1b[34mis\x1b[0m a long sentence in color", 10, "This \x1b[34mis\x1b[0m a\x1b[0m…"},
		{"This \x1b[34mis a long sentence in\x1b[0m color", 10, "This \x1b[34mis a\x1b[0m…"},
	}

	for _, tc := range testCases {
		got := ExactSize(tc.in, tc.sz)
		if got != tc.out {
			t.Errorf("Truncate failed for '%s'.\nWanted: %s\nGot   : %s\n", tc.in, tc.out, got)
		}
	}
}

func TestJustify(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out string
	}{
		{"Coucou", 10, "Coucou    "},
		{"Coucou", 6, "Coucou"},
		{"Coucou ", 6, "Coucou"},
		{"This is a long sentence", 10, "This is a \nlong      \nsentence  "},
		{"This \x1b[34mis\x1b[0m a long sentence", 10, "This \x1b[34mis\x1b[0m a \nlong      \nsentence  "},
		{"This \x1b[34mis a long sentence\x1b[0m", 10, "This \x1b[34mis a \nlong      \nsentence\x1b[0m  "},
	}

	for _, tc := range testCases {
		got := Justify(tc.in, tc.sz)
		if got != tc.out {
			t.Errorf("Justify failed for '%s'.\nWanted: '%s'\nGot   : '%s'\n", tc.in, tc.out, got)
		}
	}
}
