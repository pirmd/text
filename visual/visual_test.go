package visual

import (
	"reflect"
	"testing"
)

func TestStringwidth(t *testing.T) {
	testCases := []struct {
		in  string
		out int
	}{
		{"abc", 3},
		{"敬具", 4},
		{"a\x00bc", 3},
		{"\x1b[35ma\x00bc\x1b[0m", 3},
	}

	for _, tc := range testCases {
		got := Stringwidth(tc.in)
		if got != tc.out {
			t.Errorf("Length of '%s' failed: got %d, wanted %d", string(tc.in), got, tc.out)
		}
	}
}

func TestTruncate(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out string
	}{
		{"Coucou", 9, "Coucou"},
		{"This is a long sentence", 9, "This is a"},
		{"This \x1b[34mis\x1b[0m a long sentence in color", 9, "This \x1b[34mis\x1b[0m a"},
		{"This \x1b[34mis a long sentence in\x1b[0m color", 9, "This \x1b[34mis a\x1b[0m"},
	}

	for _, tc := range testCases {
		got := Truncate(tc.in, tc.sz)
		if got != tc.out {
			t.Errorf("visual Truncate failed for %#v.\nWanted: %#v\nGot   : %#v\n", tc.in, tc.out, got)
		}
	}
}

func TestRepeat(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out string
	}{
		{"Cou", 6, "CouCou"},
		{"\x1b[34mCou\x1b[0m", 6, "\x1b[34mCou\x1b[0m\x1b[34mCou\x1b[0m"},
		{"Cou", 4, "CouC"},
		{"\x1b[34mCou\x1b[0m", 4, "\x1b[34mCou\x1b[0m\x1b[34mC\x1b[0m"},
	}

	for _, tc := range testCases {
		got := Repeat(tc.in, tc.sz)
		if got != tc.out {
			t.Errorf("visual repeating failed for '%s' (up to %d).\nWant: %s\nGot : %s\n", tc.in, tc.sz, tc.out, got)
		}
	}
}

func TestCut(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out []string
	}{
		{"Coucou", 10, []string{"Coucou"}},
		{"Coucou", 6, []string{"Coucou"}},
		{"Coucou\n", 6, []string{"Coucou\n"}},
		{"Coucou ", 6, []string{"Coucou", " "}},
		{"Coucou\n", 8, []string{"Coucou\n"}},
		{"\x1b[34mCoucou\n\x1b[0m", 6, []string{"\x1b[34mCoucou\n\x1b[0m"}},
		{"\x1b[34mCoucou\x1b[0m\n", 6, []string{"\x1b[34mCoucou\x1b[0m\n"}},
		{"This is a long sentence", 10, []string{"This is a ", "long ", "sentence"}},
		{"This is a long sentence", 0, []string{"This is a long sentence"}},
		{"This \x1b[34mis\x1b[0m a long sentence", 10, []string{"This \x1b[34mis\x1b[0m a ", "long ", "sentence"}},
		{"This \x1b[34mis a long sentence\x1b[0m", 10, []string{"This \x1b[34mis a ", "long ", "sentence\x1b[0m"}},
		{"This \x1b[34mis a long sentence\n\x1b[0m", 10, []string{"This \x1b[34mis a ", "long ", "sentence\n\x1b[0m"}},
		{"Supercalifragilisticexpialidocious\nChim Chim Cher-ee", 10, []string{"Supercalif", "ragilistic", "expialidoc", "ious\n", "Chim Chim ", "Cher-ee"}},
		{"Supercalifragilisticexpialidocious\nChim Chim Cher-ee", -1, []string{"Supercalifragilisticexpialidocious\n", "Chim Chim Cher-ee"}},
		{
			"description: This edition contains Alice's Adventures in Wonderland. Tweedledum and Tweedledee, the Mad Hatter, the Cheshire Cat, the Red Queen and the White Rabbit all make their appearances, and are now familiar figures in writing, conversation and idiom.\nauthor: Lewis Caroll",
			58,
			[]string{"description: This edition contains Alice's Adventures in ", "Wonderland. Tweedledum and Tweedledee, the Mad Hatter, the", " Cheshire Cat, the Red Queen and the White Rabbit all make", " their appearances, and are now familiar figures in ", "writing, conversation and idiom.\n", "author: Lewis Caroll"},
		},
		{
			"All details can be found in [![GoDoc](https://godoc.org/github.com/pirmd/style?status.svg)](https://godoc.org/github.com/pirmd/style)",
			80,
			[]string{"All details can be found in [![GoDoc]", "(https://godoc.org/github.com/pirmd/style?status.svg)]", "(https://godoc.org/github.com/pirmd/style)"},
		},
		{
			"\x1b[34mWhatever goes upon two legs is an enemy.\n\x1b[39m",
			80,
			[]string{"\x1b[34mWhatever goes upon two legs is an enemy.\n\x1b[39m"},
		},
	}

	for _, tc := range testCases {
		got := Cut(tc.in, tc.sz)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("Wrap failed for %#v.\nWanted:\n%#v\nGot   :\n%#v\n", tc.in, tc.out, got)
		}
	}
}

func TestLazyCut(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out []string
	}{
		{"Coucou", 10, []string{"Coucou"}},
		{"Coucou", 6, []string{"Coucou"}},
		{"Coucou\n", 6, []string{"Coucou\n"}},
		{"Coucou ", 6, []string{"Coucou", " "}},
		{"Coucou\n", 8, []string{"Coucou\n"}},
		{"\x1b[34mCoucou\n\x1b[0m", 6, []string{"\x1b[34mCoucou\n\x1b[0m"}},
		{"\x1b[34mCoucou\x1b[0m\n", 6, []string{"\x1b[34mCoucou\x1b[0m\n"}},
		{"This is a long sentence", 10, []string{"This is a ", "long ", "sentence"}},
		{"This \x1b[34mis\x1b[0m a long sentence", 10, []string{"This \x1b[34mis\x1b[0m a ", "long ", "sentence"}},
		{"This \x1b[34mis a long sentence\x1b[0m", 10, []string{"This \x1b[34mis a ", "long ", "sentence\x1b[0m"}},
		{"Supercalifragilisticexpialidocious\nChim Chim Cher-ee", 10, []string{"Supercalifragilisticexpialidocious\n", "Chim Chim ", "Cher-ee"}},
		{
			"description: This edition contains Alice's Adventures in Wonderland. Tweedledum and Tweedledee, the Mad Hatter, the Cheshire Cat, the Red Queen and the White Rabbit all make their appearances, and are now familiar figures in writing, conversation and idiom.\nauthor: Lewis Caroll",
			58,
			[]string{"description: This edition contains Alice's Adventures in ", "Wonderland. Tweedledum and Tweedledee, the Mad Hatter, the", " Cheshire Cat, the Red Queen and the White Rabbit all make", " their appearances, and are now familiar figures in ", "writing, conversation and idiom.\n", "author: Lewis Caroll"},
		},
		{
			"All details can be found in [![GoDoc](https://godoc.org/github.com/pirmd/style?status.svg)](https://godoc.org/github.com/pirmd/style)",
			80,
			[]string{"All details can be found in [![GoDoc]", "(https://godoc.org/github.com/pirmd/style?status.svg)]", "(https://godoc.org/github.com/pirmd/style)"},
		},
		{
			"\x1b[34mWhatever goes upon two legs is an enemy.\n\x1b[39m",
			40,
			[]string{"\x1b[34mWhatever goes upon two legs is an enemy.\n\x1b[39m"},
		},
		{
			"No animal shall sleep in a bed\x1b[9m \x1b[29m\x1b[9mwithout\x1b[29m\x1b[9m \x1b[29m\x1b[9msheets\x1b[29m.\n",
			40,
			[]string{"No animal shall sleep in a bed\x1b[9m \x1b[29m\x1b[9mwithout\x1b[29m\x1b[9m \x1b[29m\x1b[9m", "sheets\x1b[29m.\n"},
		},
	}

	for _, tc := range testCases {
		got := LazyCut(tc.in, tc.sz)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("Wrap failed for %#v.\nWanted:\n%#v\nGot   :\n%#v\n", tc.in, tc.out, got)
		}
	}
}
