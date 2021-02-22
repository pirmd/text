package text

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

const (
	// MaxTableWidth defines the maximum size of a table It tries to get
	// initialized by reading the terminal width or fall-back to 80
	MaxTableWidth = 80
)

// Table represents a table.
type Table struct {
	maxWidth              int
	sepV, sepC, sepH      string
	dontTruncateLongWords bool

	cells [][]string
}

// NewTable returns a new empty table, with no grid and a maximum width set-up
// to the terminal width or to MaxTableWidth if output is not a termila or if
// the terminal size cannot be determined.
func NewTable() *Table {
	t := &Table{
		maxWidth: MaxTableWidth,
		sepV:     " ",
	}

	if fd := int(os.Stdout.Fd()); term.IsTerminal(fd) {
		if w, _, err := term.GetSize(fd); err == nil {
			t.SetMaxWidth(w)
		}
	}

	return t
}

// SetMaxWidth sets the table maximum width.
func (t *Table) SetMaxWidth(w int) *Table {
	t.maxWidth = w
	return t
}

// SetGrid defines the grid separators, respectively vertical separator between
// columns and horizontal separating captions' row from table content (before
// and after first row, after last row) and between rows.
//
// An empty separator means no separation at all.
func (t *Table) SetGrid(sepV, sepC, sepH string) *Table {
	t.sepV, t.sepC, t.sepH = sepV, sepC, sepH
	return t
}

// DontTruncateLongWords prevents long words to be cut before words boundaries
// to fit Table's column maximum width. It is not recommanded as it might void
// table formatting.
func (t *Table) DontTruncateLongWords() *Table {
	t.dontTruncateLongWords = true
	return t
}

// Rows adds a list of rows to the table. Rows trims any trailing new line.
func (t *Table) Rows(rows ...[]string) *Table {
	for _, row := range rows {
		srow := []string{}
		for _, cell := range row {
			srow = append(srow, strings.TrimSuffix(cell, "\n"))
		}
		t.cells = append(t.cells, srow)
	}

	return t
}

// RawRows adds a list of rows to the table without triming any trailing new line.
func (t *Table) RawRows(rows ...[]string) *Table {
	t.cells = append(t.cells, rows...)
	return t
}

// Col adds a list of columns to the table Col add col content at the end of
// existing rows, meaning that if rows are not of the same size or if col add
// more rows, results will not be an 'aligned' column.
// Col trims any cell trailing new line.
func (t *Table) Col(col ...[]string) *Table {
	for _, column := range col {
		for i, cell := range column {
			for row := len(t.cells); row <= i; row++ {
				//columns features more rows than actually available in the
				//table we complete by adding an empty row
				t.cells = append(t.cells, []string{})
			}
			t.cells[i] = append(t.cells[i], strings.TrimSuffix(cell, "\n"))
		}
	}
	return t
}

// RawCol adds a list of colums to the table but unlike 'Col' does not trim any
// trailing new line.
func (t *Table) RawCol(col ...[]string) *Table {
	for _, column := range col {
		for i, cell := range column {
			for row := len(t.cells); row <= i; row++ {
				//columns features more rows than actually available in the
				//table we complete by adding an empty row
				t.cells = append(t.cells, []string{})
			}
			t.cells[i] = append(t.cells[i], cell)
		}
	}
	return t
}

// Captions sets the columns' captions (first row) of the table
func (t *Table) Captions(row ...string) *Table {
	t.cells = append([][]string{row}, t.cells...)
	return t
}

// Draw draws the table.
// Columns length are determined in order to maximize the use of the table
// maximum width Text is automatically wrapped to fit into the columns size.
//
// The algorithm that defines colums' size has limitation and will provide
// unreadable output if available table's width is too short compared to content
// length.
func (t *Table) Draw() string {
	if len(t.cells) == 0 {
		return ""
	}

	//TODO(pirmd): optimize this serie of string transformation (several
	//split-join sequences.
	maxColLen := colMaxLen(t.maxWidth, visualLen(t.sepV), t.cells)

	rows := []string{}
	for _, row := range t.cells {
		col := []string{}
		for i, cell := range row {
			justifiedCell := justifyWithInterruptANSI(cell, maxColLen[i], !t.dontTruncateLongWords)
			col = append(col, justifiedCell)
		}

		rows = append(rows, columnize(t.sepV, col...))
	}

	return strings.Join(t.addHorizontalGrid(rows, maxColLen), "\n")
}

// String returns a string representation of the table
func (t *Table) String() string {
	return t.Draw()
}

func (t *Table) addHorizontalGrid(rows []string, colLen []int) []string {
	var rowsWithGrid []string

	var sepH string
	if t.sepH != "" {
		var sep []string
		for _, l := range colLen {
			sep = append(sep, visualRepeat(t.sepH, l))
		}

		sepH = strings.Join(sep, t.sepV)
	}

	var sepC string
	if t.sepC != "" {
		var sep []string
		for _, l := range colLen {
			sep = append(sep, visualRepeat(t.sepC, l))
		}

		sepC = strings.Join(sep, t.sepV)
	}

	if t.sepH != "" {
		if len(rows) > 2 {
			for _, row := range rows[2:] {
				rowsWithGrid = append(rowsWithGrid, sepH, row)
			}
		}
	} else if len(rows) > 2 {
		rowsWithGrid = append(rowsWithGrid, rows[2:]...)
	}

	if t.sepC != "" {
		if len(rows) > 1 {
			rowsWithGrid = append([]string{sepC, rows[0], sepC, rows[1]}, rowsWithGrid...)
			rowsWithGrid = append(rowsWithGrid, sepC)
		} else {
			rowsWithGrid = []string{sepC, rows[0], sepC}
		}
	} else {
		if t.sepH != "" && len(rows) > 1 {
			//Situation where we don't want any caption separation but ask for an
			//horizontal separation
			rowsWithGrid = append([]string{rows[0], sepH, rows[1]}, rowsWithGrid...)
		} else if len(rows) > 1 {
			rowsWithGrid = append(rows[0:2], rowsWithGrid...)
		} else {
			rowsWithGrid = append([]string{rows[0]}, rowsWithGrid...)
		}
	}

	return rowsWithGrid
}

func columnize(sepV string, col ...string) (row string) {
	for i, r := range multilinesRow(col) {
		if i == 0 {
			row = strings.Join(r, sepV)
		} else {
			row += "\n" + strings.Join(r, sepV)
		}
	}
	return
}

func multilinesRow(row []string) [][]string {
	var r [][]string
	var emptyCol []string

	for i, s := range row {
		col := strings.Split(s, "\n")
		colLen := maxLen(col)
		emptyCol = append(emptyCol, fmt.Sprintf("%-*s", colLen, ""))

		for icol, scol := range col {
			if icol >= len(r) {
				//Text in current column has more lines than our current
				//row. Create a new one filling other columns with blank
				var newRow []string

				for l := 0; l < i; l++ {
					newRow = append(newRow, emptyCol[l])
				}
				r = append(r, newRow)
			}

			r[icol] = append(r[icol], scol)
		}

		for l := len(col); l < len(r); l++ {
			r[l] = append(r[l], emptyCol[i])
		}
	}

	return r
}

func colMaxLen(maxWidth int, sepLen int, cells [][]string) []int {
	colLen := []int{}

	if len(cells) == 0 {
		return colLen
	}

	for _, row := range cells {
		for i, cell := range row {
			l := maxLen(strings.Split(cell, "\n"))

			if i >= len(colLen) {
				colLen = append(colLen, l)
			} else if colLen[i] <= l {
				colLen[i] = l
			}
		}
	}

	max := findMaxColWidth(colLen, maxWidth-(len(colLen)-1)*sepLen)
	for i, l := range colLen {
		if l > max {
			colLen[i] = max
		}
	}

	return colLen
}

func findMaxColWidth(colLen []int, maxWidth int) int {
	var fn func([]int, int) ([]int, int, int) //recursive function that gradually select columns that remains over size limits

	fn = func(colLen []int, maxWidth int) (overLimit []int, width int, max int) {
		if len(colLen) == 0 {
			return
		}

		max, width = maxWidth/len(colLen), maxWidth

		for _, l := range colLen {
			if l > max {
				overLimit = append(overLimit, l)
			} else {
				width -= l
			}
		}

		if len(overLimit) > 0 && len(overLimit) < len(colLen) {
			//We have found new under-limits items, try again with remaining space
			overLimit, width, max = fn(overLimit, width)
		}

		return
	}

	_, _, max := fn(colLen, maxWidth)
	return max
}

func maxLen(col []string) int {
	var length int
	for _, cell := range col {
		if l := visualLen(cell); l > length {
			length = l
		}
	}
	return length
}
