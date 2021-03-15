package text

import (
	"os"
	"strings"

	"golang.org/x/term"

	"github.com/pirmd/text/table"
	"github.com/pirmd/text/visual"
)

// Indent inserts a name/bullet/number at the beginning of the string, then
// indents it (add prefix at the beginning and before any new line).
//
// Tag is superposed to the indent prefix to obtain the first line prefix, if
// tag length is greater than prefix, prefix is completed by trailing spaces.
func Indent(s string, tag, prefix string) string {
	lT, lP := visual.Len(tag), visual.Len(prefix)

	switch {
	case lT > lP:
		prefix = visual.PadRight(prefix, lT)
	case lT < lP:
		tag = visual.Truncate(prefix, lP-lT) + tag
	}

	return indent(s, tag, prefix)
}

// Wrap wraps a text by ensuring that each of its line's "visual" length is
// lower or equal to the provided limit. Wrap works with word limits being
// spaces.
//
// If a "word" is encountered that is longer than the limit, it is split in
// chunks of 'limit' length.
func Wrap(txt string, sz int) string {
	return strings.Join(visual.Cut(txt, sz), "\n")
}

// LazyWrap wraps a text by ensuring that each of its line's "visual" length
// is lower or equal to the provided limit. Wrap works with word limits being
// spaces.
//
// If a "word" is encountered that is longer than the limit, it is not split in
// chunks of 'limit' length and is kept as is.
func LazyWrap(txt string, sz int) string {
	return strings.Join(visual.LazyCut(txt, sz), "\n")
}

// Tab wraps and indents the given text.
//
// Tab will additionally add the given tag in front of the first line. Tag is
// superposed to the indent prefix to obtain the first line prefix, if tag's
// length is greater than prefix, prefix is completed by trailing spaces.
//
// Tab calculates the correct wrapping limits taking indent's prefix length. It
// does not work if prefix is made of tabs as indent's tag/prefix length is
// unknown (like '\t').
func Tab(s string, tag, prefix string, sz int) string {
	lT, lP := visual.Len(tag), visual.Len(prefix)

	var r string
	switch {
	case lT > lP:
		prefix = visual.PadRight(prefix, lT)
		r = Wrap(s, sz-lT)
	case lT < lP:
		tag = visual.Truncate(prefix, lP-lT) + tag
		r = Wrap(s, sz-lP)
	default:
		r = Wrap(s, sz-lP)
	}

	return indent(r, tag, prefix)
}

// LazyTab wraps and indents the given text.
//
// LazyTab operates like Tab but relies on LazyWrap and does not split long
// words to fit the provided "visual" limit.
func LazyTab(s string, tag, prefix string, sz int) string {
	lT, lP := visual.Len(tag), visual.Len(prefix)

	var r string
	switch {
	case lT > lP:
		prefix = visual.PadRight(prefix, lT)
		r = LazyWrap(s, sz-lT)
	case lT < lP:
		tag = visual.Truncate(prefix, lP-lT) + tag
		r = LazyWrap(s, sz-lP)
	default:
		r = LazyWrap(s, sz-lP)
	}

	return indent(r, tag, prefix)
}

// Justify wraps a text to the given maximum size and makes sure that returned
// lines are of exact provided size by wrapping and padding them as needed.
func Justify(s string, sz int) string {
	return Left(s, sz)
}

// Left aligns text on the left and pads it with spaces to reach the given
// "visual" length.
func Left(s string, sz int) string {
	if len(s) == 0 {
		return strings.Repeat(" ", sz)
	}

	ws := visual.Cut(s, sz)
	for i, l := range ws {
		ws[i] = visual.PadRight(l, sz)
	}
	return strings.Join(ws, "\n")
}

// Right aligns text on the right and pads it with spaces to reach the given
// "visual" length.
func Right(s string, sz int) string {
	if len(s) == 0 {
		return strings.Repeat(" ", sz)
	}

	ws := visual.Cut(s, sz)
	for i, l := range ws {
		ws[i] = visual.PadLeft(l, sz)
	}
	return strings.Join(ws, "\n")
}

// Center centers text and pads it with spaces to reach the given "visual"
// length.
func Center(s string, sz int) string {
	if len(s) == 0 {
		return strings.Repeat(" ", sz)
	}

	ws := visual.Cut(s, sz)
	for i, l := range ws {
		ws[i] = visual.PadCenter(l, sz)
	}
	return strings.Join(ws, "\n")
}

// Columnize organizes supplied strings in a side-by-side fashion.
func Columnize(columns ...string) string {
	tab := table.New()

	if fd := int(os.Stdout.Fd()); term.IsTerminal(fd) {
		if w, _, err := term.GetSize(fd); err == nil {
			tab.SetMaxWidth(w)
		}
	}

	col := make([][]string, len(columns))
	for i := range columns {
		col[i] = strings.Split(strings.TrimSuffix(columns[i], "\n"), "\n")
	}

	return tab.AddCol(col...).String()
}

// Rowwise organizes supplied strings in a table-like fashion. Each string is
// displayed as a table row, each '\t' delimiting a column.
func Rowwise(rows ...string) string {
	tab := table.New()

	if fd := int(os.Stdout.Fd()); term.IsTerminal(fd) {
		if w, _, err := term.GetSize(fd); err == nil {
			tab.SetMaxWidth(w)
		}
	}

	return tab.AddTabbedRows(rows...).String()
}

// Tabulate organizes supplied string in a table-like fashion. Each '\t\n'
// delimiting a table's row and each '\t' a column.
func Tabulate(tabbedtext string) string {
	tab := table.New()

	if fd := int(os.Stdout.Fd()); term.IsTerminal(fd) {
		if w, _, err := term.GetSize(fd); err == nil {
			tab.SetMaxWidth(w)
		}
	}

	return tab.AddTabbedText(tabbedtext).String()
}

func indent(s string, firstPrefix, prefix string) string {
	var indented strings.Builder
	var isNewLine bool

	indented.WriteString(firstPrefix)

	for _, c := range s {
		//add indent prefix if we have a non-empty newline
		if isNewLine && c != '\n' {
			indented.WriteString(prefix)
		}
		indented.WriteRune(c)
		isNewLine = (c == '\n')
	}

	return indented.String()
}
