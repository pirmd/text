package txtwriter

import (
	"strings"
	"testing"
)

func TestWrap(t *testing.T) {
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
			want: "Hello ladies\nand\ngentlemen!",
		},
		{
			in:   "Hello ladies  and gentlemen!",
			want: "Hello ladies\nand\ngentlemen!",
		},
		{
			in:   "Supercalifragilisticexpialidocious",
			want: "Supercalifra\ngilisticexpi\nalidocious",
		},
		{
			in:   "Supercalifragilisticexpialidocious\n\nChim Chim Cher-ee",
			want: "Supercalifra\ngilisticexpi\nalidocious\n\nChim Chim\nCher-ee",
		},
		{
			in:   "My favorite words are: Supercalifragilisticexpialidocious\n\nChim Chim Cher-ee",
			want: "My favorite\nwords are:\nSupercalifra\ngilisticexpi\nalidocious\n\nChim Chim\nCher-ee",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(12)

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%#v'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}

func TestWrapWithANSI(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "\x1b[34mHello world!\x1b[0m",
			want: "\x1b[34mHello world!\x1b[0m",
		},
		{
			in:   "Hello \x1b[1mladies\x1b[0m and gentlemen!",
			want: "Hello \x1b[1mladies\x1b[0m\nand\ngentlemen!",
		},
		{
			in:   "This \x1b[34mis\x1b[0m a long sentence",
			want: "This \x1b[34mis\x1b[0m a\nlong\nsentence",
		},
		{
			in:   "This \x1b[34mis a long sentence\x1b[0m",
			want: "This \x1b[34mis a\nlong\nsentence\x1b[0m",
		},
		{
			in:   "\x1b[34mWhatever goes upon two legs is an enemy.\n\x1b[39m",
			want: "\x1b[34mWhatever\ngoes upon\ntwo legs is\nan enemy.\x1b[39m\n",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(12)

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%#v'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}

func TestWrapSplitLongWords(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "All details can be found in [![GoDoc](https://godoc.org/github.com/pirmd/style?status.svg)](https://godoc.org/github.com/pirmd/style)",
			want: "All details can be found in [![GoDoc]\n(https://godoc.org/github.com/pirmd/style?status.svg)]\n(https://godoc.org/github.com/pirmd/style)",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(80)

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%#v'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}

func TestLazyWrap(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "Hello world!",
			want: "Hello world!",
		},
		{
			in:   "Hello ladies and gentlemen!",
			want: "Hello ladies\nand\ngentlemen!",
		},
		{
			in:   "Supercalifragilisticexpialidocious\n\nChim Chim Cher-ee",
			want: "Supercalifragilisticexpialidocious\n\nChim Chim\nCher-ee",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(12).LazyWrap()

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%#v'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}

func TestMultipleWrite(t *testing.T) {
	testCases := []struct {
		in   []string
		want string
	}{
		{
			in:   []string{"Hello", " world", "!"},
			want: "Hello world!",
		},
		{
			in:   []string{"Hello ladi", "es and", " gentlemen!"},
			want: "Hello ladies\nand\ngentlemen!",
		},
		{
			in:   []string{"Supercalifragilisticexpialidocious\n", "\n", "Chim Chim Cher-ee"},
			want: "Supercalifra\ngilisticexpi\nalidocious\n\nChim Chim\nCher-ee",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(12)

		for _, in := range tc.in {
			tstwriter.Write([]byte(in))
		}
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%#v'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
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
		{"\n\x1b[34m  Coucou c'est nous!\n  \x1b[0m", "\n\x1b[34mCoucou c'est nous!\n  \x1b[0m"},
	}

	for _, tc := range testCases {
		got := string(trimLeadingSpace([]byte(tc.in)))
		if got != tc.out {
			t.Errorf("Trim spaces failed for %#v.\nWant:\n'%#v'\nGot :\n'%#v'\n", tc.in, tc.out, got)
		}
	}
}
