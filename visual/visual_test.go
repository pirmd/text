package visual

import (
	"reflect"
	"strings"
	"testing"
)

func TestLen(t *testing.T) {
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
		got := Len(tc.in)
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

func TestPad(t *testing.T) {
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
		got := Pad(tc.in, tc.sz, ' ')
		if got != tc.out {
			t.Errorf("visual Padding failed for '%s' (max %d).\nWanted: %s\nGot   : %s\n", tc.in, tc.sz, tc.out, got)
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
			t.Errorf("visual repeating failed for '%s' (up to %d).\nWanted: %s\nGot   : %s\n", tc.in, tc.sz, tc.out, got)
		}
	}
}

func TestInterruptFormattingatEOL(t *testing.T) {
	testCases := []struct {
		in  []string
		out []string
	}{
		{[]string{"This \x1b[34mis\x1b[0m a ", "long ", "sentence"}, []string{"This \x1b[34mis\x1b[0m a ", "long ", "sentence"}},
		{[]string{"This \x1b[34mis a ", "long ", "sentence\x1b[0m"}, []string{"This \x1b[34mis a \x1b[0m", "\x1b[34mlong \x1b[0m", "\x1b[34msentence\x1b[0m"}},
		{
			[]string{"\x1b[34mX This is a long sentence", "in color.\x1b[39m\x1b[9mAnd an error\x1b[29m"},
			[]string{"\x1b[34mX This is a long sentence\x1b[0m", "\x1b[34min color.\x1b[39m\x1b[9mAnd an error\x1b[29m"},
		},
	}

	for _, tc := range testCases {
		got := make([]string, len(tc.in))
		copy(got, tc.in)
		InterruptFormattingAtEOL(got)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("Wrap failed for %#v.\nWanted:\n%#v\nGot   :\n%#v\n", tc.in, tc.out, got)
		}
	}
}

func TestWrap(t *testing.T) {
	testCases := []struct {
		in  string
		sz  int
		out []string
	}{
		{"Coucou", 10, []string{"Coucou"}},
		{"Coucou", 6, []string{"Coucou"}},
		{"Coucou ", 6, []string{"Coucou"}},
		{"Coucou\n", 8, []string{"Coucou", ""}},
		{"This is a long sentence", 10, []string{"This is a ", "long ", "sentence"}},
		{"This \x1b[34mis\x1b[0m a long sentence", 10, []string{"This \x1b[34mis\x1b[0m a ", "long ", "sentence"}},
		{"This \x1b[34mis a long sentence\x1b[0m", 10, []string{"This \x1b[34mis a ", "long ", "sentence\x1b[0m"}},
		{"Supercalifragilisticexpialidocious\nChim Chim Cher-ee", 10, []string{"Supercalif", "ragilistic", "expialidoc", "ious", "Chim Chim ", "Cher-ee"}},
		{
			"description: This edition contains Alice's Adventures in Wonderland. Tweedledum and Tweedledee, the Mad Hatter, the Cheshire Cat, the Red Queen and the White Rabbit all make their appearances, and are now familiar figures in writing, conversation and idiom.\nauthor: Lewis Caroll",
			58,
			[]string{"description: This edition contains Alice's Adventures in ", "Wonderland. Tweedledum and Tweedledee, the Mad Hatter, the", "Cheshire Cat, the Red Queen and the White Rabbit all make ", "their appearances, and are now familiar figures in ", "writing, conversation and idiom.", "author: Lewis Caroll"},
		},
		{
			"All details can be found in [![GoDoc](https://godoc.org/github.com/pirmd/style?status.svg)](https://godoc.org/github.com/pirmd/style)",
			80,
			[]string{"All details can be found in ", "[![GoDoc](https://godoc.org/github.com/pirmd/style?status.svg)](https://godoc.or", "g/github.com/pirmd/style)"},
		},
	}

	for _, tc := range testCases {
		got := Wrap(tc.in, tc.sz)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("Wrap failed for %#v.\nWanted:\n%#v\nGot   :\n%#v\n", tc.in, tc.out, got)
		}
	}
}

func BenchmarkLen(b *testing.B) {
	in := strings.Repeat("\x1b[31mbonjour\x1b[m", 20)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Len(in)
	}
}

func BenchmarkTruncate(b *testing.B) {
	in := strings.Repeat("\x1b[31mbonjour\x1b[m", 20)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Truncate(in, len(in)/2)
	}
}
