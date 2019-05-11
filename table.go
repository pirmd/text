package text

import (
	"cli/style/termsize"
	"fmt"
	"strings"
)

const (
	//MaxTableWidth defines the maximum size of a table
	//It tries to get initialized by reading the terminal width or fall-back
	//to 80
	MaxTableWidth = 80
)

type table struct {
	maxWidth int
	sepV     string //TODO(pirmd): horizontal separator is not yet implemented

	cells [][]string
}

//Table returns a new empty table. New tables default to no grid
//with a maximum width being th eterminal size if it can be
//determined or MaxTableWidth
func Table() *table {
	if w, err := termsize.Width(); err == nil {
		return &table{
			maxWidth: w,
			sepV:     " ",
		}
	}

	return &table{
		maxWidth: MaxTableWidth,
		sepV:     " ",
	}
}

//SetMaxWidth manully set the table maximum width
func (t *table) SetMaxWidth(w int) *table {
	t.maxWidth = w
	return t
}

//SetGrid manually defines the grid separators
func (t *table) SetGrid(sepV string) *table {
	t.sepV = sepV
	return t
}

//Rows add a list of rows to the table
func (t *table) Rows(rows ...[]string) *table {
	t.cells = append(t.cells, rows...)
	return t
}

//Col add a list of columns to the table
//Col add col content at the end of existing rows, meaning
//that if rows are not of the same size or if col add more
//rows, results will not be an 'aligned' column
func (t *table) Col(col ...[]string) *table {
	for _, column := range col {
		for i, cell := range column {
			for row := len(t.cells); row <= i; row++ {
				//columns features more rows than actually available in the table
				//we complete by adding an empty row
				t.cells = append(t.cells, []string{})
			}
			t.cells[i] = append(t.cells[i], cell)
		}
	}
	return t
}

//Title sets the titles of the table
func (t *table) Title(row ...string) *table {
	t.cells = append([][]string{row}, t.cells...)
	return t
}

//Draw draws the table.
//Columns length are determined in order to maximize the use of the table maximum width
//Text is automatically wrapped to fit into the columns size.
//This algorithm to define colums' size has limitation and will provide unreadable output
//if available table's width is too short compared to content length.
func (t *table) Draw() string {
	colLen := colMaxLen(t.maxWidth, visualLen(t.sepV), t.cells)

	rows := []string{}
	for _, row := range t.cells {
		col := []string{}
		for i, cell := range row {
			col = append(col, Justify(cell, colLen[i]))
		}

		rows = append(rows, columnize(t.sepV, col...))
	}

	return strings.Join(rows, "\n")
}

//String returns a string representation of the table
func (t *table) String() string {
	return t.Draw()
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

			r[icol] = append(r[icol], visualPad(scol, colLen, ' '))
		}

		for l := len(col); l < len(r); l++ {
			r[l] = append(r[l], emptyCol[i])
		}
	}

	return r
}

func colMaxLen(maxWidth int, sepLen int, cells [][]string) []int {
	colLen := []int{}

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
	var fn func([]int, int, int) ([]int, int, int) //recursive function that gradually select columns that remains over size limits

	fn = func(colLen []int, maxWidth int, maxCol int) (overLimit []int, width int, max int) {
		max, width = maxWidth/len(colLen), maxWidth

		for _, l := range colLen {
			if l > max {
				overLimit = append(overLimit, l)
			} else {
				width -= l
			}
		}

		switch {
		case len(overLimit) == 0: //All columns are already of the right size
			maxCol = max
		case len(overLimit) < len(colLen): //We have found new under-limits items, try again with remaining space
			overLimit, width, max = fn(overLimit, width, max)
			//default: we have found no new under-limit item, have to stop
		}

		return
	}

	_, _, max := fn(colLen, maxWidth, maxWidth/len(colLen))
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
