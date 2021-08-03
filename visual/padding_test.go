package visual

import (
	"testing"
)

func TestPadRight(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out string
	}{
		{"Coucou", 9, "Coucou   "},
		{"\x1b[34mCoucou\x1b[0m", 9, "\x1b[34mCoucou\x1b[0m   "},
		{"Coucou", 6, "Coucou"},
		{"\x1b[34mCoucou\x1b[0m", 6, "\x1b[34mCoucou\x1b[0m"},
		{"Coucou, c'est nous", 9, "Coucou, c'est nous"},
		{"", 3, "   "},
	}

	for _, tc := range testCases {
		got := string(PadRight([]byte(tc.in), tc.sz))
		if got != tc.out {
			t.Errorf("visual Padding failed for '%s' (max %d).\nWant: %s\nGot : %s\n", tc.in, tc.sz, tc.out, got)
		}
	}
}

func TestPadLeft(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out string
	}{
		{"Coucou", 9, "   Coucou"},
		{"\x1b[34mCoucou\x1b[0m", 9, "   \x1b[34mCoucou\x1b[0m"},
		{"This \x1b[34mis\x1b[0m a long sentence", 30, "       This \x1b[34mis\x1b[0m a long sentence"},
		{"Coucou", 6, "Coucou"},
		{"Coucou, c'est nous", 9, "Coucou, c'est nous"},
		{"", 3, "   "},
	}

	for _, tc := range testCases {
		got := string(PadLeft([]byte(tc.in), tc.sz))
		if got != tc.out {
			t.Errorf("visual Padding failed for '%s' (max %d).\nWant: %#v\nGot : %#v\n", tc.in, tc.sz, tc.out, got)
		}
	}
}

func TestPadCenter(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out string
	}{
		{"Coucou", 10, "  Coucou  "},
		{"Coucou ", 10, "  Coucou  "},
		{"  Coucou ", 10, "  Coucou  "},
		{"  Coucou ", 9, " Coucou  "},
		{"\x1b[34mCoucou\x1b[0m", 9, " \x1b[34mCoucou\x1b[0m  "},
		{"Coucou", 6, "Coucou"},
		{"Coucou, c'est nous", 9, "Coucou, c'est nous"},
		{"", 3, "   "},
	}

	for _, tc := range testCases {
		got := string(PadCenter([]byte(tc.in), tc.sz))
		if got != tc.out {
			t.Errorf("visual Padding failed for '%s' (max %d).\nWant: '%s'\nGot : '%s'\n", tc.in, tc.sz, tc.out, got)
		}
	}
}
