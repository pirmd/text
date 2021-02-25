package table

import (
	"reflect"
	"testing"
)

func TestCol2Rows(t *testing.T) {
	testCases := []struct {
		in  [][]string
		out [][]string
	}{
		{
			[][]string{{"val1.1", "val1.2", "val1.3"}, {"val2.1", "val2.2", "val2.3"}},
			[][]string{{"val1.1", "val2.1"}, {"val1.2", "val2.2"}, {"val1.3", "val2.3"}},
		},

		{
			[][]string{{"val1.1", "val1.2"}, {"val2.1", "val2.2", "val2.3"}},
			[][]string{{"val1.1", "val2.1"}, {"val1.2", "val2.2"}, {"", "val2.3"}},
		},
	}

	for _, tc := range testCases {
		got := col2rows(tc.in)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("table failed for '%#v'.\nWant:\n%#v\nGot :\n%#v\n", tc.in, tc.out, got)
		}
	}
}

func TestTableAddRows(t *testing.T) {
	testCases := []struct {
		in  [][]string
		out string
	}{
		{
			[][]string{
				{"val1.1", "val1.2", "val1.3"}, {"val2.1", "val2.2", "val2.3"},
			},
			"val1.1|val1.2|val1.3\nval2.1|val2.2|val2.3",
		},
		{
			[][]string{
				{"val1.1", "val1.2", "val1.3"}, {"val2.1", "val2.2.1\nval2.2.2", "val2.3"},
			},
			"val1.1|val1.2  |val1.3\nval2.1|val2.2.1|val2.3\n      |val2.2.2|      ",
		},
		{
			[][]string{
				{"val1.1", "val1.2", "val1.3"}, {"val2.1", "val2.2\nval2.2.1", "val2.3"},
			},
			"val1.1|val1.2  |val1.3\nval2.1|val2.2  |val2.3\n      |val2.2.1|      ",
		},
		{
			[][]string{
				{"val1.1", "val1.2", "val1.3"}, {"val2.1", "val2.2.1 val2.2.2", "val2.3"},
			},
			"val1.1|val1.2    |val1.3\nval2.1|val2.2.1  |val2.3\n      |val2.2.2  |      ",
		},
		{
			[][]string{
				{"", "val1.2", "val1.3"}, {"val2.1", "val2.2.1 val2.2.2", "val2.3"},
			},
			"      |val1.2    |val1.3\nval2.1|val2.2.1  |val2.3\n      |val2.2.2  |      ",
		},
		{
			[][]string{
				{"\x1b[31mval1", "\x1b[34mval2\nval2.1\x1b[0m\nval2.2", "val3\nval3.1"},
			},
			"\x1b[31mval1\x1b[0m|\x1b[34mval2\x1b[0m  |val3  \n    |\x1b[34mval2.1\x1b[0m|val3.1\n    |val2.2|      ",
		},
	}

	for _, tc := range testCases {
		got := New().SetGrid(&Grid{Columns: "|"}).SetMaxWidth(24).AddRows(tc.in...).String()
		if got != tc.out {
			t.Errorf("table failed for '%#v'.\nWanted:\n%s\nGot   :\n%s\n", tc.in, tc.out, got)
		}
	}
}

func TestTableAddCol(t *testing.T) {
	testCases := []struct {
		in  [][]string
		out string
	}{
		{
			[][]string{
				{"val1.1", "val2.1"}, {"val1.2", "val2.2"}, {"val1.3", "val2.3"},
			},
			"val1.1|val1.2|val1.3\nval2.1|val2.2|val2.3",
		},
		{
			[][]string{
				{"val1.1", "val2.1"}, {"val1.2", "val2.2.1\nval2.2.2"}, {"val1.3", "val2.3"},
			},
			"val1.1|val1.2  |val1.3\nval2.1|val2.2.1|val2.3\n      |val2.2.2|      ",
		},
		{
			[][]string{
				{"val1.1", "val2.1"}, {"val1.2", "val2.2.1 val2.2.2"}, {"val1.3", "val2.3"},
			},
			"val1.1|val1.2    |val1.3\nval2.1|val2.2.1  |val2.3\n      |val2.2.2  |      ",
		},
	}

	for _, tc := range testCases {
		got := New().SetGrid(&Grid{Columns: "|"}).SetMaxWidth(24).AddCol(tc.in...).String()
		if got != tc.out {
			t.Errorf("table failed for '%s'.\nWanted:\n%s\nGot   :\n%s\n", tc.in, tc.out, got)
		}
	}
}

func TestTableAddTabbedText(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{
			"val1.1\tval1.2\tval1.3\t\nval2.1\tval2.2\tval2.3\t\n",
			"val1.1|val1.2|val1.3\nval2.1|val2.2|val2.3",
		},
		{
			"val1.1\tval1.2\tval1.3\t\nval2.1\tval2.2.1\nval2.2.2\tval2.3\t\n",
			"val1.1|val1.2  |val1.3\nval2.1|val2.2.1|val2.3\n      |val2.2.2|      ",
		},
		{
			"val1.1\tval1.2\tval1.3\t\nval2.1\tval2.2\nval2.2.1\tval2.3\t\n",
			"val1.1|val1.2  |val1.3\nval2.1|val2.2  |val2.3\n      |val2.2.1|      ",
		},
		{
			"val1.1\tval1.2\tval1.3\t\nval2.1\tval2.2.1 val2.2.2\tval2.3\t\n",
			"val1.1|val1.2    |val1.3\nval2.1|val2.2.1  |val2.3\n      |val2.2.2  |      ",
		},
		{
			"\tval1.2\tval1.3\t\nval2.1\tval2.2.1 val2.2.2\tval2.3\t\n",
			"      |val1.2    |val1.3\nval2.1|val2.2.1  |val2.3\n      |val2.2.2  |      ",
		},
		{
			"\x1b[31mval1\t\x1b[34mval2\nval2.1\x1b[0m\nval2.2\tval3\nval3.1\t\n",
			"\x1b[31mval1\x1b[0m|\x1b[34mval2\x1b[0m  |val3  \n    |\x1b[34mval2.1\x1b[0m|val3.1\n    |val2.2|      ",
		},
	}

	for _, tc := range testCases {
		got := New().SetGrid(&Grid{Columns: "|"}).SetMaxWidth(24).AddTabbedText(tc.in).String()
		if got != tc.out {
			t.Errorf("table failed for '%#v'.\nWanted:\n%s\nGot   :\n%s\n", tc.in, tc.out, got)
		}
	}
}

func TestTableWithGrid(t *testing.T) {
	testTable := [][]string{{"val1.1", "val1.2", "val1.3"}, {"val2.1", "val2.2", "val2.3"}}
	testTable2 := [][]string{{"val1.1", "val1.2", "val1.3"}, {"val2.1", "val2.2", "val2.3"}, {"val3.1", "val3.2", "val3.3"}}

	testCases := []struct {
		inTheader []string
		inTbody   [][]string
		inTfooter []string
		inSep     *Grid
		out       string
	}{
		{
			inTbody: testTable,
			inSep:   &Grid{Columns: "|"},
			out:     "val1.1|val1.2|val1.3\nval2.1|val2.2|val2.3",
		},
		{
			inTbody: testTable,
			inSep:   &Grid{Columns: "|", BodyRows: "-"},
			out:     "val1.1|val1.2|val1.3\n------|------|------\nval2.1|val2.2|val2.3",
		},
		{
			inTbody: testTable,
			inSep:   &Grid{Columns: "|", BodyRows: "-**-"},
			out:     "val1.1|val1.2|val1.3\n-**--*|-**--*|-**--*\nval2.1|val2.2|val2.3",
		},
		{
			inTheader: []string{"A", "B", "C"},
			inTbody:   testTable2,
			inSep:     &Grid{Columns: "|"},
			out:       "A     |B     |C     \nval1.1|val1.2|val1.3\nval2.1|val2.2|val2.3\nval3.1|val3.2|val3.3",
		},
		{
			inTheader: []string{"A", "B", "C"},
			inTbody:   testTable2,
			inSep:     &Grid{Columns: "|", Header: "="},
			out:       "A     |B     |C     \n======|======|======\nval1.1|val1.2|val1.3\nval2.1|val2.2|val2.3\nval3.1|val3.2|val3.3",
		},
		{
			inTheader: []string{"A", "B", "C"},
			inTbody:   testTable2,
			inSep:     &Grid{Columns: "|", BodyRows: "-"},
			out:       "A     |B     |C     \n------|------|------\nval1.1|val1.2|val1.3\n------|------|------\nval2.1|val2.2|val2.3\n------|------|------\nval3.1|val3.2|val3.3",
		},
		{
			inTheader: []string{"A", "B", "C"},
			inTbody:   testTable2,
			inTfooter: []string{"", "S", "42"},
			inSep:     &Grid{Columns: "|", BodyRows: "-"},
			out:       "A     |B     |C     \n------|------|------\nval1.1|val1.2|val1.3\n------|------|------\nval2.1|val2.2|val2.3\n------|------|------\nval3.1|val3.2|val3.3\n------|------|------\n      |S     |42    ",
		},
		{
			inTheader: []string{"A", "B", "C"},
			inTbody:   testTable2,
			inTfooter: []string{"", "S", "42"},
			inSep:     &Grid{Columns: "|", Header: "*", BodyRows: "-", Footer: "="},
			out:       "A     |B     |C     \n******|******|******\nval1.1|val1.2|val1.3\n------|------|------\nval2.1|val2.2|val2.3\n------|------|------\nval3.1|val3.2|val3.3\n======|======|======\n      |S     |42    ",
		},
	}

	for _, tc := range testCases {
		got := New().SetGrid(tc.inSep).SetMaxWidth(24).AddRows(tc.inTbody...).SetHeader(tc.inTheader...).SetFooter(tc.inTfooter...).String()
		if got != tc.out {
			t.Errorf("table failed for sep='%#v'.\nWanted:\n%s\nGot   :\n%s\n", tc.inSep, tc.out, got)
		}
	}
}

func TestInterruptFormattingAtEOL(t *testing.T) {
	testCases := []struct {
		in  []string
		out []string
	}{
		{
			[]string{"This \x1b[34mis\x1b[0m a ", "long ", "sentence"},
			[]string{"This \x1b[34mis\x1b[0m a ", "long ", "sentence"},
		},
		{
			[]string{"This \x1b[34mis a ", "long ", "sentence\x1b[0m"},
			[]string{"This \x1b[34mis a \x1b[0m", "\x1b[34mlong \x1b[0m", "\x1b[34msentence\x1b[0m"},
		},
		{
			[]string{"\x1b[34m This is a long sentence", "in color.\x1b[39m \x1b[9mAnd an error\x1b[29m"},
			[]string{"\x1b[34m This is a long sentence\x1b[0m", "\x1b[34min color.\x1b[39m \x1b[9mAnd an error\x1b[29m"},
		},
	}

	for _, tc := range testCases {
		got := make([]string, len(tc.in))
		copy(got, tc.in)
		interruptFormattingAtEOL(got)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("Wrap failed for %#v.\nWanted:\n%#v\nGot   :\n%#v\n", tc.in, tc.out, got)
		}
	}
}
