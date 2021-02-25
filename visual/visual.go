package visual

import (
	"strings"
	"unicode"

	"github.com/mattn/go-runewidth"

	"github.com/pirmd/text/ansi"
)

// Len calculates the length of string ignoring any ANSI escape sequences.
func Len(s string) int {
	var l int
	_ = ansi.Walk(s, func(c rune, esc string) error {
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

	_ = ansi.Walk(s, func(c rune, esc string) error {
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

	ts.WriteString(sgr.Reset())
	return ts.String()
}

// PadRight completes a string with provided pattern until its "visual" size
// reaches the provided limit. If the string's "visual" size is already above
// limit, PadRight returns it as-is.
func PadRight(s string, pattern string, sz int) string {
	if l := Len(s); l < sz {
		return s + Repeat(pattern, sz-l)
	}
	return s
}

// PadLeft prefixes a string with provided pattern until its "visual" size
// reaches the provided limit. If the string's "visual" size is already above
// limit, PadLeft returns it as-is.
func PadLeft(s string, pattern string, sz int) string {
	if l := Len(s); l < sz {
		return Repeat(pattern, sz-l) + s
	}
	return s
}

// PadCenter equally prefixes and complete a string with provided pattern until
// its "visual" size reaches the provided limit. If the string's "visual" size
// is already above limit, PadCenter returns it as-is.
func PadCenter(s string, pattern string, sz int) string {
	if l := Len(s); l < sz {
		right := (sz - l) / 2
		return Repeat(pattern, right) + s + Repeat(pattern, (sz-l)-right)
	}
	return s
}

// Repeat repeats s until given size is reached.
func Repeat(s string, sz int) string {
	l := Len(s)
	if l >= sz {
		return s
	}

	var rs strings.Builder

	var i int
	for i <= sz {
		rs.WriteString(s)
		i += l
	}

	if i == sz {
		return rs.String()
	}

	return Truncate(rs.String(), sz)
}

// Runewidth returns the visual width of a rune.
func Runewidth(c rune) int {
	return runewidth.RuneWidth(c)
}

// InterruptFormattingAtEOL interrupts at each line any ANSI SGR rendition and
// continues it at the next line (useful to work with text in column to avoid
// voiding neighbour text).
func InterruptFormattingAtEOL(s []string) {
	var sgr ansi.Sequence
	var prevEsc string

	for i, line := range s {
		_ = ansi.Walk(line, func(c rune, esc string) error {
			if c == -1 {
				sgr.Combine(esc)
			}
			return nil
		})
		s[i] = prevEsc + s[i] + sgr.Reset()
		prevEsc = sgr.Esc()
	}
}

// Wrap wraps a text by ensuring that each of its line's "visible" length
// is lower or equal to the provided limit. Wrap works with word limits being
// spaces.
//
// If a "word" is encountered that is longer than the limit, it is split in
// chunks of 'limit' length.
func Wrap(s string, sz int) (ws []string) {
	var line, word string
	var linelen, wordlen int

	_ = ansi.Walk(s, func(c rune, esc string) error {
		switch {
		case c == -1:
			word += esc

		case c == '\n':
			if linelen+wordlen <= sz {
				line += word
			} else {
				ws = append(ws, line)
				line = word
			}

			ws = append(ws, line)
			line, linelen = "", 0
			word, wordlen = "", 0

		case unicode.IsSpace(c):
			switch l := linelen + wordlen; {
			case l == sz:
				ws = append(ws, line+word)
				line, linelen = "", 0

			case l > sz:
				ws = append(ws, line)
				line = word + string(c)
				linelen = wordlen + Runewidth(c)

			default:
				line += word + string(c)
				linelen += wordlen + Runewidth(c)
			}
			word, wordlen = "", 0

		default:
			if wordlen += Runewidth(c); wordlen > sz {
				if line != "" {
					ws = append(ws, line)
					line, linelen = "", 0
				}

				// word is longer than the sz, we split it at the current
				// position.
				//TODO(pirmd): find some way to split word at more meaningful
				//position (like at [()[]/.)
				ws = append(ws, word)
				word, wordlen = "", Runewidth(c)
			}
			word += string(c)

		}

		return nil
	})

	switch l := linelen + wordlen; {
	case l == 0 && strings.HasSuffix(s, "\n"):
		ws = append(ws, "")
	case l == 0:
	case l > sz && linelen == 0:
		ws = append(ws, word)
	case l > sz:
		if line != "" {
			ws = append(ws, line)
		}
		ws = append(ws, word)
	default:
		ws = append(ws, line+word)
	}

	return
}
