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

// Len calculates the "visual" length of string.
func Len(s string) int {
	var l int
	_ = ansi.WalkString(s, func(n int, c rune, esc string) error {
		if c > -1 {
			l += Runewidth(c)
		}
		return nil
	})

	return l
}

// Truncate truncates the string so that its "visible" length is lower or equal
// to the provided limit.
// When needed, Truncate terminates the string by an ansi.Reset sequence
// to inhibit any visual effects coming from the truncation step.
func Truncate(s string, limit int) string {
	var ts strings.Builder
	var l int
	var sgr ansi.Sequence

	_ = ansi.WalkString(s, func(n int, c rune, esc string) error {
		if c > -1 {
			ts.WriteRune(c)
			l += Runewidth(c)
		}

		if esc != "" {
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
	ts, l := TrimSpace(s)
	if l < sz {
		return ts + strings.Repeat(" ", sz-l)
	}
	return s
}

// PadLeft prefixes a string with spaces until its "visual" size
// reaches the provided limit. If the string's "visual" size is already above
// limit, PadLeft returns it as-is.
func PadLeft(s string, sz int) string {
	ts, l := TrimSpace(s)
	if l < sz {
		return strings.Repeat(" ", sz-l) + ts
	}
	return s
}

// PadCenter equally prefixes and complete a string with spaces until its
// "visual" size reaches the provided limit. If the string's "visual" size is
// already above limit, PadCenter returns it as-is.
func PadCenter(s string, sz int) string {
	ts, l := TrimSpace(s)
	if l < sz {
		right := (sz - l) / 2
		return strings.Repeat(" ", right) + ts + strings.Repeat(" ", (sz-l)-right)
	}
	return s
}

// Repeat repeats s until given "visual" size is reached.
func Repeat(s string, sz int) string {
	var rs strings.Builder

	i, l := 0, Len(s)
	for i <= sz {
		rs.WriteString(s)
		i += l
	}

	if i == sz {
		return rs.String()
	}

	return Truncate(rs.String(), sz)
}

// Cut cuts a string at end-of-line if the line "visual" length is shorter than
// the given limit or at word boundary (space) to stay as close as possible
// under the given limit.
// Should a word exists that is longer than the limit, the word is split in
// chunks of given limit.
// If provided limit is zero or less than zero, Cut only acts at end-of-line.
func Cut(s string, sz int) (chunks []string) {
	return cut(s, sz, false)
}

// LazyCut cuts a string at end-of-line if the line "visual" length is shorter
// than the given limit or at word boundary (space) to stay as close as
// possible under the given limit.
// Should a word exists that is longer than the limit, the word is not split to
// fit into the given limit ans is kept as is.
// If provided limit is zero or less than zero, Cut only acts at end-of-line.
func LazyCut(s string, sz int) (chunks []string) {
	return cut(s, sz, true)
}

// TrimSpace  trims all leading and trailing white space (as defined by
// Unicode). TrimSpace returns the trimmed strings as well as its "visual"
// length.
func TrimSpace(s string) (string, int) {
	var trimmed strings.Builder
	var spaceBuf strings.Builder
	var spaceBufSGR ansi.Sequence
	var l, buflen int

	isLeadingSpaces := true
	_ = ansi.WalkString(s, func(n int, c rune, esc string) error {
		switch {
		case c == -1:
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
	return trimmed.String(), l
}

// TrimSuffix trims trailing rune r.
func TrimSuffix(s string, r rune) string {
	var trimmed strings.Builder
	var buf strings.Builder
	var bufSGR ansi.Sequence

	_ = ansi.WalkString(s, func(n int, c rune, esc string) error {
		switch {
		case c == -1:
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

func cut(s string, sz int, lazy bool) (chunks []string) {
	var line strings.Builder
	var word strings.Builder
	var linelen, wordlen int

	flushln := func() {
		chunks = append(chunks, line.String())
		line.Reset()
		linelen = 0
	}

	flushwd := func() {
		line.WriteString(word.String())
		linelen += wordlen
		word.Reset()
		wordlen = 0
	}

	flushrune := func(c rune) {
		word.WriteRune(c)
		wordlen += Runewidth(c)
	}

	_ = ansi.WalkString(s, func(n int, c rune, esc string) error {
		switch {
		case c == -1:
			word.WriteString(esc)

		case c == '\n':
			if sz > 0 && linelen+wordlen > sz {
				flushln()
			}
			flushwd()
			flushln()

		case unicode.IsSpace(c):
			switch l := linelen + wordlen; {
			case sz > 0 && l == sz:
				flushwd()
				flushln()

			case sz > 0 && l > sz:
				flushln()
				flushrune(c)
				flushwd()

			default:
				flushrune(c)
				flushwd()
			}

		default:
			// word is longer than sz, we split it at the current position.
			if sz > 0 && wordlen+Runewidth(c) > sz {
				if linelen > 0 {
					flushln()
				}

				if !lazy {
					// TODO(pirmd): try to split at meaningful rune (like at ()[]/.)
					flushwd()
					flushln()
				}
			}
			flushrune(c)
		}

		return nil
	})

	if sz > 0 && linelen+wordlen > sz {
		flushln()
	}
	flushwd()
	flushln()

	return
}
