package table

import (
	"testing"
)

func TestTable(t *testing.T) {
	testCases := []struct {
		in  [][]string
		out string
	}{
		{[][]string{{"Col1.1", "Col1.2", "Col1.3"}, {"Col2.1", "Col2.2", "Col2.3"}}, "Col1.1|Col1.2|Col1.3\nCol2.1|Col2.2|Col2.3"},
		{[][]string{{"Col1.1", "Col1.2", "Col1.3"}, {"Col2.1", "Col2.2.1\nCol2.2.2", "Col2.3"}}, "Col1.1|Col1.2  |Col1.3\nCol2.1|Col2.2.1|Col2.3\n      |Col2.2.2|      "},
		{[][]string{{"Col1.1", "Col1.2", "Col1.3"}, {"Col2.1", "Col2.2\nCol2.2.1", "Col2.3"}}, "Col1.1|Col1.2  |Col1.3\nCol2.1|Col2.2  |Col2.3\n      |Col2.2.1|      "},
		{[][]string{{"Col1.1", "Col1.2", "Col1.3"}, {"Col2.1", "Col2.2.1 Col2.2.2", "Col2.3"}}, "Col1.1|Col1.2    |Col1.3\nCol2.1|Col2.2.1  |Col2.3\n      |Col2.2.2  |      "},
		{[][]string{{"", "Col1.2", "Col1.3"}, {"Col2.1", "Col2.2.1 Col2.2.2", "Col2.3"}}, "      |Col1.2    |Col1.3\nCol2.1|Col2.2.1  |Col2.3\n      |Col2.2.2  |      "},
		{[][]string{{"\x1b[31mCol1", "\x1b[34mCol2\nCol2.1\x1b[0m\nCol2.2", "Col3\nCol3.1"}}, "\x1b[31mCol1\x1b[0m|\x1b[34mCol2\x1b[0m  |Col3  \n    |\x1b[34mCol2.1\x1b[0m|Col3.1\n    |Col2.2|      "},
	}

	for _, tc := range testCases {
		got := New().SetGrid(&Grid{Columns: "|"}).SetMaxWidth(24).AddRows(tc.in...).Draw()
		if got != tc.out {
			t.Errorf("table failed for '%#v'.\nWanted:\n%s\nGot   :\n%s\n", tc.in, tc.out, got)
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
		got := New().SetGrid(&Grid{Columns: "|"}).SetMaxWidth(24).AddCol(tc.in...).Draw()
		if got != tc.out {
			t.Errorf("table failed for '%s'.\nWanted:\n%s\nGot   :\n%s\n", tc.in, tc.out, got)
		}
	}
}

func TestTableWithGrid(t *testing.T) {
	testTable := [][]string{{"Col1.1", "Col1.2", "Col1.3"}, {"Col2.1", "Col2.2", "Col2.3"}}
	testTable2 := [][]string{{"Col1.1", "Col1.2", "Col1.3"}, {"Col2.1", "Col2.2", "Col2.3"}, {"Col3.1", "Col3.2", "Col3.3"}}

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
			out:     "Col1.1|Col1.2|Col1.3\nCol2.1|Col2.2|Col2.3",
		},
		{
			inTbody: testTable,
			inSep:   &Grid{Columns: "|", BodyRows: "-"},
			out:     "Col1.1|Col1.2|Col1.3\n------|------|------\nCol2.1|Col2.2|Col2.3",
		},
		{
			inTbody: testTable,
			inSep:   &Grid{Columns: "|", BodyRows: "-**-"},
			out:     "Col1.1|Col1.2|Col1.3\n-**--*|-**--*|-**--*\nCol2.1|Col2.2|Col2.3",
		},
		{
			inTheader: []string{"A", "B", "C"},
			inTbody:   testTable2,
			inSep:     &Grid{Columns: "|"},
			out:       "A     |B     |C     \nCol1.1|Col1.2|Col1.3\nCol2.1|Col2.2|Col2.3\nCol3.1|Col3.2|Col3.3",
		},
		{
			inTheader: []string{"A", "B", "C"},
			inTbody:   testTable2,
			inSep:     &Grid{Columns: "|", Header: "="},
			out:       "A     |B     |C     \n======|======|======\nCol1.1|Col1.2|Col1.3\nCol2.1|Col2.2|Col2.3\nCol3.1|Col3.2|Col3.3",
		},
		{
			inTheader: []string{"A", "B", "C"},
			inTbody:   testTable2,
			inSep:     &Grid{Columns: "|", BodyRows: "-"},
			out:       "A     |B     |C     \n------|------|------\nCol1.1|Col1.2|Col1.3\n------|------|------\nCol2.1|Col2.2|Col2.3\n------|------|------\nCol3.1|Col3.2|Col3.3",
		},
		{
			inTheader: []string{"A", "B", "C"},
			inTbody:   testTable2,
			inTfooter: []string{"", "S", "42"},
			inSep:     &Grid{Columns: "|", BodyRows: "-"},
			out:       "A     |B     |C     \n------|------|------\nCol1.1|Col1.2|Col1.3\n------|------|------\nCol2.1|Col2.2|Col2.3\n------|------|------\nCol3.1|Col3.2|Col3.3\n------|------|------\n      |S     |42    ",
		},
		{
			inTheader: []string{"A", "B", "C"},
			inTbody:   testTable2,
			inTfooter: []string{"", "S", "42"},
			inSep:     &Grid{Columns: "|", Header: "*", BodyRows: "-", Footer: "="},
			out:       "A     |B     |C     \n******|******|******\nCol1.1|Col1.2|Col1.3\n------|------|------\nCol2.1|Col2.2|Col2.3\n------|------|------\nCol3.1|Col3.2|Col3.3\n======|======|======\n      |S     |42    ",
		},
	}

	for _, tc := range testCases {
		got := New().SetGrid(tc.inSep).SetMaxWidth(24).AddRows(tc.inTbody...).SetHeader(tc.inTheader...).SetFooter(tc.inTfooter...).Draw()
		if got != tc.out {
			t.Errorf("table failed for sep='%#v'.\nWanted:\n%s\nGot   :\n%s\n", tc.inSep, tc.out, got)
		}
	}
}
