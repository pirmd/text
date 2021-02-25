package ansi

import (
	"reflect"
	"testing"
)

func TestParseSGR(t *testing.T) {
	testCases := []struct {
		in   string
		want Sequence
	}{
		{in: cCSI + "31m", want: Sequence{cRed}},
		{in: cCSI + "1;31m", want: Sequence{cBold, cRed}},
		{in: cCSI + "m", want: Sequence{cReset}},
		{in: cCSI + "1;m", want: Sequence{cBold, cReset}},
		{in: cCSI + "1;;m", want: Sequence{cBold, cReset, cReset}},
		{in: cCSI + "48;5;1;32m", want: Sequence{"48;5;1", cGreen}},
		{in: cCSI + "38;2;1;2;3;32m", want: Sequence{"38;2;1;2;3", cGreen}},
		{in: cCSI + "A", want: nil},
	}

	for _, tc := range testCases {
		got := ParseSGR(tc.in)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Fail to initiate proper SGR sequence from %#v.\nWant: %#v\nGot : %#v", tc.in, tc.want, got)
		}
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		in   []Code
		want Sequence
	}{
		{
			[]Code{cRed, cBold},
			Sequence{cRed, cBold},
		},

		{
			[]Code{cRed, cGreen},
			Sequence{cGreen},
		},

		{
			[]Code{cRed, cBold, cReset},
			Sequence{cReset},
		},

		{
			[]Code{cRed, cBold, cBlue, cNormal},
			Sequence{cBlue},
		},
	}

	for _, tc := range testCases {
		got := Sequence{}
		for _, c := range tc.in {
			got.add(c)
		}

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Failed to append %#v.\nWant: %v\nGot : %v\n", tc.in, tc.want, got)
		}
	}
}

func TestCombine(t *testing.T) {
	testCases := []struct {
		in   []string
		want Sequence
	}{
		{
			[]string{RedOn, BoldOn},
			Sequence{cRed, cBold},
		},

		{
			[]string{RedOn, GreenOn},
			Sequence{cGreen},
		},

		{
			[]string{RedOn, BoldOn, Reset},
			Sequence{cReset},
		},

		{
			[]string{RedOn, BoldOn, BlueOn, Normal},
			Sequence{cBlue},
		},

		{
			[]string{RedOn, BoldOn, "\x1b[A", BlueOn, Normal},
			Sequence{cBlue},
		},
	}

	for _, tc := range testCases {
		got := Sequence{}
		for _, esc := range tc.in {
			got.Combine(esc)
		}

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Failed to combine %#v.\nWant: %#v\nGot : %#v\n", tc.in, tc.want, got)
		}
	}
}
