package text

import (
	"testing"
)

func TestColumnize(t *testing.T) {
	testCases := []struct {
		in  []string
		out string
	}{
		{[]string{"Col1", "Col2", "Col3"}, "Col1|Col2|Col3"},
		{[]string{"Col1", "Col2\nCol2.1\nCol2.2", "Col3\nCol3.1"}, "Col1|Col2  |Col3  \n    |Col2.1|Col3.1\n    |Col2.2|      "},
	}

	for _, tc := range testCases {
		got := columnize("|", tc.in...)
		if got != tc.out {
			t.Errorf("Columnize failed for '%s'.\nWanted:\n%s\nGot   :\n%s\n", tc.in, tc.out, got)
		}
	}
}

func TestTable(t *testing.T) {
	testCases := []struct {
		in  [][]string
		out string
	}{
		{[][]string{{"Col1.1", "Col1.2", "Col1.3"}, {"Col2.1", "Col2.2", "Col2.3"}}, "Col1.1|Col1.2|Col1.3\nCol2.1|Col2.2|Col2.3"},
		{[][]string{{"Col1.1", "Col1.2", "Col1.3"}, {"Col2.1", "Col2.2.1\nCol2.2.2", "Col2.3"}}, "Col1.1|Col1.2  |Col1.3\nCol2.1|Col2.2.1|Col2.3\n      |Col2.2.2|      "},
		{[][]string{{"Col1.1", "Col1.2", "Col1.3"}, {"Col2.1", "Col2.2.1 Col2.2.2", "Col2.3"}}, "Col1.1|Col1.2    |Col1.3\nCol2.1|Col2.2.1  |Col2.3\n      |Col2.2.2  |      "},
	}

	for _, tc := range testCases {
		got := Table().SetGrid("|").SetMaxWidth(24).Rows(tc.in...).String()
		if got != tc.out {
			t.Errorf("table failed for '%s'.\nWanted:\n%s\nGot   :\n%s\n", tc.in, tc.out, got)
		}
	}
}

func TestTableByCol(t *testing.T) {
	testCases := []struct {
		in  [][]string
		out string
	}{
		{[][]string{{"Col1.1", "Col2.1"}, {"Col1.2", "Col2.2"}, {"Col1.3", "Col2.3"}}, "Col1.1|Col1.2|Col1.3\nCol2.1|Col2.2|Col2.3"},
		{[][]string{{"Col1.1", "Col2.1"}, {"Col1.2", "Col2.2.1\nCol2.2.2"}, {"Col1.3", "Col2.3"}}, "Col1.1|Col1.2  |Col1.3\nCol2.1|Col2.2.1|Col2.3\n      |Col2.2.2|      "},
		{[][]string{{"Col1.1", "Col2.1"}, {"Col1.2", "Col2.2.1 Col2.2.2"}, {"Col1.3", "Col2.3"}}, "Col1.1|Col1.2    |Col1.3\nCol2.1|Col2.2.1  |Col2.3\n      |Col2.2.2  |      "},
	}

	for _, tc := range testCases {
		got := Table().SetGrid("|").SetMaxWidth(24).Col(tc.in...).String()
		if got != tc.out {
			t.Errorf("table failed for '%s'.\nWanted:\n%s\nGot   :\n%s\n", tc.in, tc.out, got)
		}
	}
}
