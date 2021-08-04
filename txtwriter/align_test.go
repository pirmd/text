package txtwriter

import (
	"strings"
	"testing"
)

func TestAlignLeftPadWithSpaces(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "Hello world!",
			want: "Hello world!",
		},
		{
			in:   "Hello!",
			want: "Hello!      ",
		},
		{
			in:   "Hello ladies and gentlemen!",
			want: "Hello ladies\nand         \ngentlemen!  ",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(12).AlignLeft().PadLeft()

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%s'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}

func TestAlignRight(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "Hello world!",
			want: "Hello world!",
		},
		{
			in:   "Hello!",
			want: "      Hello!",
		},
		{
			in:   "Hello ladies and gentlemen!",
			want: "Hello ladies\n         and\n  gentlemen!",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(12).AlignRight()

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%s'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}

func TestAlignRightPadWithSpaces(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "Hello world!",
			want: "Hello world!",
		},
		{
			in:   "Hello!",
			want: "      Hello!",
		},
		{
			in:   "Hello ladies and gentlemen!",
			want: "Hello ladies\n         and\n  gentlemen!",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(12).AlignRight().PadLeft()

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%s'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}

func TestAlignCenter(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "Hello world!",
			want: "Hello world!",
		},
		{
			in:   "Hello!",
			want: "   Hello!",
		},
		{
			in:   "Hello ladies and gentlemen!",
			want: "Hello ladies\n    and\n gentlemen!",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(12).AlignCenter()

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%s'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}

func TestAlignCenterPadWithSpaces(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{
			in:   "Hello world!",
			want: "Hello world!",
		},
		{
			in:   "Hello!",
			want: "   Hello!   ",
		},
		{
			in:   "Hello ladies and gentlemen!",
			want: "Hello ladies\n    and     \n gentlemen! ",
		},
	}

	for _, tc := range testCases {
		got := new(strings.Builder)

		tstwriter := New(got).SetMaxWidth(12).AlignCenter().PadLeft()

		tstwriter.Write([]byte(tc.in))
		tstwriter.Flush()

		if got.String() != tc.want {
			t.Errorf("Fail to write '%s'.\nWant:\n%#v\n\nGot :\n%#v", tc.in, tc.want, got.String())
		}
	}
}
