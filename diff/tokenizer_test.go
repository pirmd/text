package diff

import (
	"reflect"
	"testing"
)

func TestByLines(t *testing.T) {
	testCases := []struct {
		in   string
		want []string
	}{
		{"abc", []string{"abc"}},
		{"abc\ncde", []string{"abc\n", "cde"}},
		{"abc\n\ncde", []string{"abc\n", "\n", "cde"}},
		{"abc\ncde\n", []string{"abc\n", "cde\n", ""}},
		{"import (\n\t\"strings\"\n\t\"path\"\n)", []string{"import (\n", "\t\"strings\"\n", "\t\"path\"\n", ")"}},
		{"import (\n\t\"os\"\n\t\"strings\"\n\t\"path/filepath\"\n)", []string{"import (\n", "\t\"os\"\n", "\t\"strings\"\n", "\t\"path/filepath\"\n", ")"}},
	}

	for _, tc := range testCases {
		got := ByLines(tc.in)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Failed to split %#v bylines.\nGot : %#v\nWant: %#v\n", tc.in, got, tc.want)
		}
	}
}

func TestByWords(t *testing.T) {
	testCases := []struct {
		in   string
		want []string
	}{
		{"abc", []string{"abc"}},
		{"abc cde", []string{"abc", " ", "cde"}},
		{"abc, cde", []string{"abc", ",", " ", "cde"}},
		{"abc\ncde", []string{"abc", "\n", "cde"}},
		{"\t\"path\"\n)", []string{"\t", "\"", "path", "\"", "\n", ")"}},
		{"\t\"path/filepath\"\n)", []string{"\t", "\"", "path", "/", "filepath", "\"", "\n", ")"}},
	}

	for _, tc := range testCases {
		got := ByWords(tc.in)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Failed to split %#v by words.\nGot : %#v\nWant: %#v\n", tc.in, got, tc.want)
		}
	}
}
