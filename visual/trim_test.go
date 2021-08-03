package visual

import (
	"testing"
)

func TestTrimSpace(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"Coucou", "Coucou"},
		{"Coucou ", "Coucou"},
		{"  Coucou \n", "Coucou"},
		{"This \x1b[34mis\x1b[0m a long sentence", "This \x1b[34mis\x1b[0m a long sentence"},
		{"\x1b[34m  Coucou c'est\x1b[0m nous!  \n", "\x1b[34mCoucou c'est\x1b[0m nous!"},
		{" \x1b[34m Coucou c'est nous!\x1b[0m  \n", "\x1b[34mCoucou c'est nous!\x1b[0m"},
		{"\x1b[34m  Coucou c'est nous!\n  \x1b[0m", "\x1b[34mCoucou c'est nous!\x1b[0m"},
	}

	for _, tc := range testCases {
		got := string(TrimSpace([]byte(tc.in)))
		if got != tc.out {
			t.Errorf("Trim spaces failed for %#v.\nWant:\n'%#v'\nGot :\n'%#v'\n", tc.in, tc.out, got)
		}
	}
}

func TestTrimLeadingSpace(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"Coucou", "Coucou"},
		{"Coucou ", "Coucou "},
		{"  Coucou \n", "Coucou \n"},
		{"This \x1b[34mis\x1b[0m a long sentence", "This \x1b[34mis\x1b[0m a long sentence"},
		{"\x1b[34m  Coucou c'est\x1b[0m nous!  \n", "\x1b[34mCoucou c'est\x1b[0m nous!  \n"},
		{" \x1b[34m Coucou c'est nous!\x1b[0m  \n", "\x1b[34mCoucou c'est nous!\x1b[0m  \n"},
		{"\n\x1b[34m  Coucou c'est nous!\n  \x1b[0m", "\x1b[34mCoucou c'est nous!\n  \x1b[0m"},
	}

	for _, tc := range testCases {
		got := string(TrimLeadingSpace([]byte(tc.in)))
		if got != tc.out {
			t.Errorf("Trim spaces failed for %#v.\nWant:\n'%#v'\nGot :\n'%#v'\n", tc.in, tc.out, got)
		}
	}
}

func TestTrimTrailingSpace(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"Coucou", "Coucou"},
		{"Coucou ", "Coucou"},
		{"  Coucou \n", "  Coucou"},
		{"This \x1b[34mis\x1b[0m a long sentence", "This \x1b[34mis\x1b[0m a long sentence"},
		{"\x1b[34m  Coucou c'est\x1b[0m nous!  \n", "\x1b[34m  Coucou c'est\x1b[0m nous!"},
		{" \x1b[34m Coucou c'est nous!\x1b[0m  \n", " \x1b[34m Coucou c'est nous!\x1b[0m"},
		{"\x1b[34m  Coucou c'est nous!\n  \x1b[0m", "\x1b[34m  Coucou c'est nous!\x1b[0m"},
	}

	for _, tc := range testCases {
		got := string(TrimTrailingSpace([]byte(tc.in)))
		if got != tc.out {
			t.Errorf("Trim spaces failed for %#v.\nWant:\n'%#v'\nGot :\n'%#v'\n", tc.in, tc.out, got)
		}
	}
}

func TestTrimSuffix(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"Coucou", "Coucou"},
		{"Coucou \n", "Coucou "},
		{"\x1b[34mCoucou\x1b[0m\n", "\x1b[34mCoucou\x1b[0m"},
		{"\x1b[34mCoucou\n\x1b[0m", "\x1b[34mCoucou\x1b[0m"},
		{"Coucou\nc'est\nnous\n", "Coucou\nc'est\nnous"},
		{"Coucou\nc'est\x1b[34m\nnous\n\x1b[0m", "Coucou\nc'est\x1b[34m\nnous\x1b[0m"},
	}

	for _, tc := range testCases {
		got := string(TrimSuffix([]byte(tc.in), '\n'))
		if got != tc.out {
			t.Errorf("Trim (EOL) failed for %#v.\nWant:\n'%#v'\nGot :\n'%#v'\n", tc.in, tc.out, got)
		}
	}
}
