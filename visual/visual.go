package visual

import (
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/pirmd/text/ansi"
)

// Runewidth returns the visual width of a rune.
func Runewidth(c rune) int {
	return runewidth.RuneWidth(c)
}

// Width returns the "visual" width of a slice of bytes.
func Width(p []byte) (w int) {
	_ = ansi.Walk(p, func(n int, c rune, esc string) error {
		if len(esc) == 0 {
			w += Runewidth(c)
		}
		return nil
	})

	return
}

// Stringwidth returns the "visual" width of string.
func Stringwidth(s string) int {
	return Width([]byte(s))
}

// Repeat repeats s until given "visual" size is reached.
func Repeat(s string, sz int) string {
	var rs strings.Builder

	i, l := 0, Stringwidth(s)
	for i <= sz {
		rs.WriteString(s)
		i += l
	}

	if i == sz {
		return rs.String()
	}

	return Truncate(rs.String(), sz)
}

// Truncate truncates the string so that its "visible" length is lower or equal
// to the provided limit.
// When needed, Truncate terminates the string by an ansi.Reset sequence
// to inhibit any visual effects coming from the truncation step.
func Truncate(s string, limit int) string {
	var ts strings.Builder
	var l int
	var sgr ansi.Sequence

	_ = ansi.WalkString(s, func(advance int, c rune, esc string) error {
		if c > -1 {
			ts.WriteRune(c)
			l += Runewidth(c)
		}

		if len(esc) > 0 {
			ts.WriteString(esc)
			sgr.Combine(esc)
		}

		if l >= limit {
			return ansi.ErrStopWalk
		}

		return nil
	})

	ts.WriteString(sgr.Off())
	return ts.String()
}

// Cut cuts a string at end-of-line if the line "visual" length is shorter than
// the given limit or at word boundary (space) to stay as close as possible
// under the given limit.
// Should a word exists that is longer than the limit, the word is split in
// pieces.
// If provided limit is zero or less than zero, Cut only acts at end-of-line.
func Cut(s string, sz int) (chunks []string) {
	return cut(s, NewCutter(sz))
}

// LazyCut cuts a string at end-of-line if the line "visual" length is shorter
// than the given limit or at word boundary (space) to stay as close as
// possible under the given limit.
// Should a word exists that is longer than the limit, the word is not split to
// fit into the given limit ans is kept as is.
// If provided limit is zero or less than zero, Cut only acts at end-of-line.
func LazyCut(s string, sz int) (chunks []string) {
	return cut(s, NewLazyCutter(sz))
}

func cut(s string, cutr *Cutter) (chunks []string) {
	line, cut := cutr.Split([]byte(s))

	for line != nil {
		chunks = append(chunks, string(line))
		line, cut = cutr.Split(cut)
	}

	if len(cut) > 0 {
		chunks = append(chunks, string(cut))
	}

	return
}
