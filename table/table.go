package table

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/term"

	"github.com/pirmd/text/internal/util"
)

const (
	// DefaultMaxWidth defines the maximum size of a table It tries to get
	// initialized by reading the terminal width or fall-back to 80.
	DefaultMaxWidth = 80
)

// Grid represents a set of Table's grid decoration.
type Grid struct {
	// Columns is the separator pattern between two columns.
	Columns string
	// Header is the separator pattern between the header and the table's body.
	// Default to BodyRows.
	Header string
	// BodyRows is the separator pattern between two consecutive table's body
	// rows.
	BodyRows string
	// Footer is the separator pattern between the table's body and the footer.
	// Default to BodyRows.
	Footer string
}

// Table represents a table.
type Table struct {
	// header contains the Table's header's row.
	header []string
	// body contains the Table"s body's rows.
	body [][]string
	// footer contains the Table's footer's row .
	footer []string

	// maxWidth is the maximum allowed width of the Table.
	// default to DefaultMaxWidth or to the terminal width if it can be
	// determined at runtime.
	maxWidth int
	// colWidth contains the expected width of the Table's columns. If colWidth
	// is empty it is automatically determined before drawing the table based
	// on Table's maxWidth and on the actual maximum width of the rows'
	// content.
	colWidth []int

	// sep contains the patterns to draw the Table's grid.
	sep *Grid
}

// New returns a new empty table, with no grid and a maximum width set-up to
// the terminal width to DefaultMaxWidth.
func New() *Table {
	return &Table{
		maxWidth: DefaultMaxWidth,
		sep:      &Grid{Columns: " "},
	}
}

// SetMaxWidth sets the table maximum width.
func (t *Table) SetMaxWidth(w int) *Table {
	t.maxWidth = w
	return t
}

// SetTermWidth set Table's maximum width to the current terminal's width. Does
// nothing if Terminal size cannot be determined (or it current output is not a
// terminal)
func (t *Table) SetTermWidth() *Table {
	if fd := int(os.Stdout.Fd()); term.IsTerminal(fd) {
		if w, _, err := term.GetSize(fd); err == nil {
			t.SetMaxWidth(w)
		}
	}
	return t
}

// SetColWidth sets the table's column width. If not set, Table will
// auto-determined the column width based on Table's max width.
func (t *Table) SetColWidth(w ...int) *Table {
	t.colWidth = w
	return t
}

// SetGrid defines the grid separators.
func (t *Table) SetGrid(sep *Grid) *Table {
	t.sep = sep
	if t.sep.Header == "" {
		t.sep.Header = t.sep.BodyRows
	}
	if t.sep.Footer == "" {
		t.sep.Footer = t.sep.BodyRows
	}
	return t
}

// SetHeader sets the Table's header (first row).
func (t *Table) SetHeader(row ...string) *Table {
	t.header = append([]string{}, row...)
	return t
}

// SetFooter sets the Table's footer (last row).
func (t *Table) SetFooter(row ...string) *Table {
	t.footer = append([]string{}, row...)
	return t
}

// AddRows adds a list of rows to the table's body.
func (t *Table) AddRows(rows ...[]string) *Table {
	for _, row := range rows {
		srow := []string{}
		for _, cell := range row {
			srow = append(srow, strings.TrimSuffix(cell, "\n"))
		}
		t.body = append(t.body, srow)
	}

	return t
}

// AddCol adds a list of columns to the table's body.
func (t *Table) AddCol(columns ...[]string) *Table {
	for _, col := range columns {
		for i, cell := range col {
			for row := len(t.body); row <= i; row++ {
				//columns features more rows than actually available in the
				//table we complete by adding an empty row
				t.body = append(t.body, []string{})
			}
			t.body[i] = append(t.body[i], strings.TrimSuffix(cell, "\n"))
		}
	}
	return t
}

// AddTabbedText adds to table's body text whose columns are separated by "\t"
// and rows by "\t\n".
func (t *Table) AddTabbedText(tabbedtext string) *Table {
	lines := strings.Split(strings.TrimSuffix(tabbedtext, "\t\n"), "\t\n")

	var rows [][]string
	for _, row := range lines {
		rows = append(rows, strings.Split(row, "\t"))
	}

	return t.AddRows(rows...)
}

// String returns a string representation of the table
func (t *Table) String() string {
	var s strings.Builder
	t.WriteTo(&s)
	return s.String()
}

// WriteTo actually draws the Table to an io.Writer.
// Columns width, if not manually defined, is automatically determined to fit
// Table maximum width Table's text is automatically wrapped to fit into the
// columns size.
func (t *Table) WriteTo(w io.Writer) (int64, error) {
	var nbytes int

	t.autoColWidth()

	sepHeader := t.buildSeparator(t.sep.Header)
	sepRow := t.buildSeparator(t.sep.BodyRows)
	sepFooter := t.buildSeparator(t.sep.Footer)

	if len(t.header) > 0 {
		n, err := t.writeRowTo(w, t.header)
		nbytes += n
		if err != nil {
			return int64(nbytes), err
		}

		if len(t.body) > 0 {
			n, err := t.writeSepTo(w, sepHeader)
			nbytes += n
			if err != nil {
				return int64(nbytes), err
			}
		}
	}

	lastRow := len(t.body) - 1
	for i, row := range t.body {
		n, err := t.writeRowTo(w, row)
		nbytes += n
		if err != nil {
			return int64(nbytes), err
		}

		if i != lastRow {
			n, err := t.writeSepTo(w, sepRow)
			nbytes += n
			if err != nil {
				return int64(nbytes), err
			}
		}
	}

	if len(t.footer) > 0 {
		if len(t.body) > 0 || len(t.header) > 0 {
			n, err := t.writeSepTo(w, sepFooter)
			nbytes += n
			if err != nil {
				return int64(nbytes), err
			}
		}

		n, err := t.writeRowTo(w, t.footer)
		nbytes += n
		if err != nil {
			return int64(nbytes), err
		}
	}

	return int64(nbytes), nil
}

func (t *Table) writeRowTo(w io.Writer, row []string) (int, error) {
	var nbytes int

	for i, subrow := range t.padRow(row) {
		if i > 0 {
			n, err := fmt.Fprint(w, "\n")
			nbytes += n
			if err != nil {
				return nbytes, err
			}
		}

		for j, cell := range subrow {
			if j > 0 {
				n, err := fmt.Fprint(w, t.sep.Columns)
				nbytes += n
				if err != nil {
					return nbytes, err
				}
			}

			n, err := fmt.Fprint(w, cell)
			nbytes += n
			if err != nil {
				return nbytes, err
			}
		}
	}

	return nbytes, nil
}

func (t *Table) padRow(row []string) [][]string {
	paddedRow := [][]string{}

	for icol, cell := range row {
		lines := columnize(cell, t.colWidth[icol])

		for len(lines) > len(paddedRow) {
			paddedRow = append(paddedRow, t.emptyLine())
		}

		for iline := range lines {
			paddedRow[iline][icol] = lines[iline]
		}
	}

	return paddedRow
}

func (t *Table) emptyLine() []string {
	l := make([]string, len(t.colWidth))
	for i := range t.colWidth {
		l[i] = strings.Repeat(" ", t.colWidth[i])
	}
	return l
}

func (t *Table) buildSeparator(pattern string) string {
	if pattern == "" {
		return ""
	}

	sep := make([]string, len(t.colWidth))
	for i := range t.colWidth {
		sep[i] = util.VisualRepeat(pattern, t.colWidth[i])
	}

	return strings.Join(sep, t.sep.Columns)
}

func (t *Table) writeSepTo(w io.Writer, sep string) (int, error) {
	var nbytes int

	n, err := fmt.Fprint(w, "\n")
	nbytes += n
	if err != nil {
		return nbytes, err
	}

	if sep != "" {
		n, err := fmt.Fprint(w, sep)
		nbytes += n
		if err != nil {
			return nbytes, err
		}

		n, err = fmt.Fprint(w, "\n")
		nbytes += n
		if err != nil {
			return nbytes, err
		}
	}

	return nbytes, nil

}

// autoColWidth calculates the column's width of the Table based on the Table's
// maxWidth and the cells maximum width.
func (t *Table) autoColWidth() {
	t.colWidth = make([]int, len(t.header))
	for i, cell := range t.header {
		l := cellWidth(cell)

		if i >= len(t.colWidth) {
			t.colWidth = append(t.colWidth, l)
		} else if t.colWidth[i] <= l {
			t.colWidth[i] = l
		}
	}

	for _, row := range t.body {
		for i, cell := range row {
			l := cellWidth(cell)

			if i >= len(t.colWidth) {
				t.colWidth = append(t.colWidth, l)
			} else if t.colWidth[i] <= l {
				t.colWidth[i] = l
			}
		}
	}

	for i, cell := range t.footer {
		l := cellWidth(cell)

		if i >= len(t.colWidth) {
			t.colWidth = append(t.colWidth, l)
		} else if t.colWidth[i] <= l {
			t.colWidth[i] = l
		}
	}

	maxUsableWidth := t.maxWidth - (len(t.colWidth)-1)*util.VisualLen(t.sep.Columns)
	max := findWidthLimit(t.colWidth, maxUsableWidth)
	for i, l := range t.colWidth {
		if l > max {
			t.colWidth[i] = max
		}
	}
}

// findWidthLimit uses an heuristic finding the best combination for the
// column's width so that the total Table's width is close to its maxWidth
// while keeping as much columns as possible close to their maximum width.
// In a nutshell:
// - allocate max width equally to all columns,
// - get back the extra space not needed that much space,
// - reallocate extra space to columns that need more space,
// - rinse and repeat until either no space left or no columns above allocated
// space.
func findWidthLimit(width []int, max int) int {
	var fn func([]int, int) ([]int, int, int) //recursive function that gradually selects columns that remains over size limits

	fn = func(width []int, max int) (overLimit []int, w int, m int) {
		if len(width) == 0 {
			return
		}

		m, w = max/len(width), max

		for _, l := range width {
			if l > m {
				overLimit = append(overLimit, l)
			} else {
				w -= l
			}
		}

		if len(overLimit) > 0 && len(overLimit) < len(width) {
			//We have found new under-limits items, try again with remaining space
			overLimit, w, m = fn(overLimit, w)
		}

		return
	}

	_, _, m := fn(width, max)
	return m
}

// cellWidth finds the width of a cell.
func cellWidth(cell string) int {
	var length int
	for _, line := range strings.Split(cell, "\n") {
		if l := util.VisualLen(line); l > length {
			length = l
		}
	}
	return length
}

// columnize justifies a text and properly interrupt ANSI sequences at line
// boundaries so that it can easily be combined with another text without
// voiding the formatting.
func columnize(s string, sz int) []string {
	if len(s) == 0 {
		return []string{strings.Repeat(" ", sz)}
	}

	ws := util.VisualWrap(s, sz, true)

	util.InterruptFormattingAtEOL(ws)

	for i, l := range ws {
		ws[i] = util.VisualPad(l, sz, ' ')
	}
	return ws
}
