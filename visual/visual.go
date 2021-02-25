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
// When needed, visualTruncate terminates the string by an ansi.Reset sequence
// to inhibate any visual effects coming from the truncation step.
func Truncate(s string, limit int) (trunc string) {
	var l int
	var sgr ansi.Sequence

	_ = ansi.Walk(s, func(c rune, esc string) error {
		if c > -1 {
			trunc += string(c)
			l += Runewidth(c)
		}

		if esc != "" {
			trunc += esc
			sgr.Combine(esc)
		}

		if l >= limit {
			return ansi.ErrStopWalk
		}

		return nil
	})

	return trunc + sgr.Reset()
}

// Pad completes a string with provided rune until its "visible" lentgh
// reaches the provided limit.  If the string visible size is already above
// imit, Pad returns it as-is.
func Pad(s string, size int, padRune rune) string {
	var pad []rune
	for i := Len(s); i < size; i++ {
		pad = append(pad, padRune)
	}
	return s + string(pad)
}

// Repeat repeats s until given size is reached.
func Repeat(s string, size int) string {
	l := Len(s)
	r := s
	for i := l; i <= size; i = i + l {
		if i == size {
			return r
		}
		r = r + s
	}
	return Truncate(r, size)
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
// If a "word" is encountered that is longer than the limit, it is truncated or
// left as is depending of truncateLongWords flag.
func Wrap(s string, limit int, truncateLongWords bool) (ws []string) {
	var line, word string
	var linelen, wordlen int

	_ = ansi.Walk(s, func(c rune, esc string) error {
		switch {
		case c == -1:
			word += esc

		case c == '\n':
			if linelen+wordlen <= limit {
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
			case l == limit:
				ws = append(ws, line+word)
				line, linelen = "", 0

			case l > limit:
				ws = append(ws, line)
				line = word + string(c)
				linelen = wordlen + Runewidth(c)

			default:
				line += word + string(c)
				linelen += wordlen + Runewidth(c)
			}
			word, wordlen = "", 0

		default:
			if wordlen += Runewidth(c); wordlen > limit {
				if line != "" {
					ws = append(ws, line)
					line, linelen = "", 0
				}

				// word is longer than the limit, we truncated it
				if truncateLongWords {
					ws = append(ws, word)
					word, wordlen = "", Runewidth(c)
				}
			}
			word += string(c)

		}

		return nil
	})

	switch l := linelen + wordlen; {
	case l == 0 && strings.HasSuffix(s, "\n"):
		ws = append(ws, "")
	case l == 0:
	case l > limit && linelen == 0:
		ws = append(ws, word)
	case l > limit:
		if line != "" {
			ws = append(ws, line)
		}
		ws = append(ws, word)
	default:
		ws = append(ws, line+word)
	}

	return
}
