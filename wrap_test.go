package text

import (
	"strings"
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
		{"Supercalifragilisticexpialidocious\nChim Chim Cher-ee", 10, "Supercalif\nragilistic\nexpialidoc\nious\nChim Chim \nCher-ee"},
		{"description: This edition contains Alice's Adventures in Wonderland. Tweedledum and Tweedledee, the Mad Hatter, the Cheshire Cat, the Red Queen and the White Rabbit all make their appearances, and are now familiar figures in writing, conversation and idiom.\nauthor: Lewis Caroll", 58, "description: This edition contains Alice's Adventures in \nWonderland. Tweedledum and Tweedledee, the Mad Hatter, the\nCheshire Cat, the Red Queen and the White Rabbit all make \ntheir appearances, and are now familiar figures in \nwriting, conversation and idiom.\nauthor: Lewis Caroll"},
		{
			"All details can be found in [![GoDoc](https://godoc.org/github.com/pirmd/style?status.svg)](https://godoc.org/github.com/pirmd/style)",
			80,
			"All details can be found in \n[![GoDoc](https://godoc.org/github.com/pirmd/style?status.svg)](https://godoc.or\ng/github.com/pirmd/style)",
		},
	}

	for _, tc := range testCases {
		got := Wrap(tc.in, tc.sz, true)
		if got != tc.out {
			t.Errorf("Wrap failed for %#v.\nWanted:\n%#v\nGot   :\n%#v\n", tc.in, tc.out, got)
		}
	}
}

func TestWrapAndIndent(t *testing.T) {
	testCases := []struct {
		inTxt string
		inP   string
		out   string
	}{
		{
			"Test",
			"    ",
			"    Test",
		},
		{
			"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped because it hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached.",
			"  > ",
			"  > This very long and detailed sentence\n  > is here to demonstrate that list can\n  > be formatted and wrapped because it \n  > hopefully be so long that it will \n  > not fulfill the maximum number of \n  > authorized char per lines is \n  > reached.",
		},
		{
			"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped because it hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached.",
			"  \x1b[31m>\x1b[0m ",
			"  \x1b[31m>\x1b[0m This very long and detailed sentence\n  \x1b[31m>\x1b[0m is here to demonstrate that list can\n  \x1b[31m>\x1b[0m be formatted and wrapped because it \n  \x1b[31m>\x1b[0m hopefully be so long that it will \n  \x1b[31m>\x1b[0m not fulfill the maximum number of \n  \x1b[31m>\x1b[0m authorized char per lines is \n  \x1b[31m>\x1b[0m reached.",
		},
	}

	for _, tc := range testCases {
		got := Tab(tc.inTxt, "", tc.inP, 40, true)
		if got != tc.out {
			t.Errorf("Indenting or/and wrapping failed.\nWanted:\n%#v\nGot   :\n%#v\n", tc.out, got)
		}
	}
}

func TestTabWithTag(t *testing.T) {
	testCases := []struct {
		inTxt    string
		inT, inP string
		out      string
	}{
		{
			"Test",
			"  - ",
			"    ",
			"  - Test",
		},
		{
			"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped because it hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached.",
			"  \x1b[31m-\x1b[0m ",
			"  \x1b[31m>\x1b[0m ",
			"  \x1b[31m-\x1b[0m This very long and detailed sentence\n  \x1b[31m>\x1b[0m is here to demonstrate that list can\n  \x1b[31m>\x1b[0m be formatted and wrapped because it \n  \x1b[31m>\x1b[0m hopefully be so long that it will \n  \x1b[31m>\x1b[0m not fulfill the maximum number of \n  \x1b[31m>\x1b[0m authorized char per lines is \n  \x1b[31m>\x1b[0m reached.",
		},
		{
			"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped because it hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached.",
			"- ",
			"  > ",
			"  - This very long and detailed sentence\n  > is here to demonstrate that list can\n  > be formatted and wrapped because it \n  > hopefully be so long that it will \n  > not fulfill the maximum number of \n  > authorized char per lines is \n  > reached.",
		},
		{
			"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped because it hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached.",
			"\x1b[31m-\x1b[0m ",
			"  \x1b[31m>\x1b[0m ",
			"  \x1b[31m-\x1b[0m This very long and detailed sentence\n  \x1b[31m>\x1b[0m is here to demonstrate that list can\n  \x1b[31m>\x1b[0m be formatted and wrapped because it \n  \x1b[31m>\x1b[0m hopefully be so long that it will \n  \x1b[31m>\x1b[0m not fulfill the maximum number of \n  \x1b[31m>\x1b[0m authorized char per lines is \n  \x1b[31m>\x1b[0m reached.",
		},
		{
			"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped because it hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached.",
			"- ",
			"",
			"- This very long and detailed sentence \n  is here to demonstrate that list can \n  be formatted and wrapped because it \n  hopefully be so long that it will not \n  fulfill the maximum number of \n  authorized char per lines is reached.",
		},
		{
			"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped because it hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached.",
			"  - ",
			"> ",
			"  - This very long and detailed sentence\n>   is here to demonstrate that list can\n>   be formatted and wrapped because it \n>   hopefully be so long that it will \n>   not fulfill the maximum number of \n>   authorized char per lines is \n>   reached.",
		},
	}

	for _, tc := range testCases {
		got := Tab(tc.inTxt, tc.inT, tc.inP, 40, true)
		if got != tc.out {
			t.Errorf("Inserting tag (tag='%s', indent='%s') failed.\nWanted:\n%#v\nGot   :\n%#v\n", tc.inT, tc.inP, tc.out, got)
		}
	}
}

func TestIndent(t *testing.T) {
	testCases := []struct {
		inTxt    string
		inT, inP string
		out      string
	}{
		{
			"Test",
			"  - ",
			"    ",
			"  - Test",
		},
		{
			"This two-lines paragraph will be indented with a bullet on the first line.\nSecond line is only here for test with no real message or interesting content.",
			"  \x1b[31m-\x1b[0m ",
			"  \x1b[31m>\x1b[0m ",
			"  \x1b[31m-\x1b[0m This two-lines paragraph will be indented with a bullet on the first line.\n  \x1b[31m>\x1b[0m Second line is only here for test with no real message or interesting content.",
		},
		{
			"This two-lines paragraph will be indented with a bullet on the first line.\nSecond line is only here for test with no real message or interesting content.",
			"- ",
			"  > ",
			"  - This two-lines paragraph will be indented with a bullet on the first line.\n  > Second line is only here for test with no real message or interesting content.",
		},
		{
			"This two-lines paragraph will be indented with a bullet on the first line.\nSecond line is only here for test with no real message or interesting content.",
			"\x1b[31m-\x1b[0m ",
			"  \x1b[31m>\x1b[0m ",
			"  \x1b[31m-\x1b[0m This two-lines paragraph will be indented with a bullet on the first line.\n  \x1b[31m>\x1b[0m Second line is only here for test with no real message or interesting content.",
		},
		{
			"This two-lines paragraph will be indented with a bullet on the first line.\nSecond line is only here for test with no real message or interesting content.",
			"- ",
			"",
			"- This two-lines paragraph will be indented with a bullet on the first line.\n  Second line is only here for test with no real message or interesting content.",
		},
		{
			"This two-lines paragraph will be indented with a bullet on the first line.\nSecond line is only here for test with no real message or interesting content.",
			"  - ",
			"> ",
			"  - This two-lines paragraph will be indented with a bullet on the first line.\n>   Second line is only here for test with no real message or interesting content.",
		},
	}

	for _, tc := range testCases {
		got := Indent(tc.inTxt, tc.inT, tc.inP)
		if got != tc.out {
			t.Errorf("Inserting tag (tag='%s', indent='%s') failed.\nWanted:\n%#v\nGot   :\n%#v\n", tc.inT, tc.inP, tc.out, got)
		}
	}
}

func BenchmarkWrap(b *testing.B) {
	in := strings.Repeat("\x1b[31mbonjour\x1b[m ", 20)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Wrap(in, len(in)/20, true)
	}
}
