package visual

import (
	"strings"
	"unicode"

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

// PadRight completes a string with spaces until its "visual" size
// reaches the provided limit. If the string's "visual" size is already above
// limit, PadRight returns it as-is.
func PadRight(s string, sz int) string {
	ts, l := TrimSpaceString(s)
	if l < sz {
		return ts + strings.Repeat(" ", sz-l)
	}
	return s
}

// PadLeft prefixes a string with spaces until its "visual" size
// reaches the provided limit. If the string's "visual" size is already above
// limit, PadLeft returns it as-is.
func PadLeft(s string, sz int) string {
	ts, l := TrimSpaceString(s)
	if l < sz {
		return strings.Repeat(" ", sz-l) + ts
	}
	return s
}

// PadCenter equally prefixes and complete a string with spaces until its
// "visual" size reaches the provided limit. If the string's "visual" size is
// already above limit, PadCenter returns it as-is.
func PadCenter(s string, sz int) string {
	ts, l := TrimSpaceString(s)
	if l < sz {
		right := (sz - l) / 2
		return strings.Repeat(" ", right) + ts + strings.Repeat(" ", (sz-l)-right)
	}
	return s
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

// TrimSpace  trims all leading and trailing space (as defined by Unicode) from
// a slice of bytes.
// TrimSpace returns also the "visual" width of trimmed slice.
func TrimSpace(s []byte) ([]byte, int) {
	var trimmed strings.Builder
	var spaceBuf strings.Builder
	var spaceBufSGR ansi.Sequence
	var l, buflen int

	isLeadingSpaces := true
	_ = ansi.Walk(s, func(advance int, c rune, esc string) error {
		switch {
		case len(esc) > 0:
			spaceBuf.WriteString(esc)
			spaceBufSGR.Combine(esc)

		case unicode.IsSpace(c):
			if !isLeadingSpaces {
				spaceBuf.WriteRune(c)
				buflen += Runewidth(c)
			}

		default:
			isLeadingSpaces = false

			trimmed.WriteString(spaceBuf.String())
			l += buflen
			spaceBuf.Reset()
			buflen = 0
			spaceBufSGR.Reset()

			trimmed.WriteRune(c)
			l += Runewidth(c)
		}

		return nil
	})

	trimmed.WriteString(spaceBufSGR.String())
	return []byte(trimmed.String()), l
}

// TrimSpaceString trims all leading and trailing spaces (as defined by
// Unicode) from a string.
// TrimSpaceString returns also the "visual" width of trimmed string.
func TrimSpaceString(s string) (string, int) {
	out, sz := TrimSpace([]byte(s))
	return string(out), sz
}

// TrimSuffix trims trailing rune r.
func TrimSuffix(s string, r rune) string {
	var trimmed strings.Builder
	var buf strings.Builder
	var bufSGR ansi.Sequence

	_ = ansi.WalkString(s, func(advance int, c rune, esc string) error {
		switch {
		case len(esc) > 0:
			buf.WriteString(esc)
			bufSGR.Combine(esc)

		case c == r:
			buf.WriteRune(c)

		default:
			trimmed.WriteString(buf.String())
			buf.Reset()
			bufSGR.Reset()

			trimmed.WriteRune(c)
		}

		return nil
	})
	trimmed.WriteString(bufSGR.String())
	return trimmed.String()
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
