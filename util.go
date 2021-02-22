package text

import (
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/pirmd/text/ansi"
)

// visualLen calculates the length of string ignoring any ANSI escape sequences.
func visualLen(s string) int {
	var l int
	_ = ansi.Walk(s, func(c rune, esc string) error {
		if c > -1 {
			l += runeWidth(c)
		}
		return nil
	})

	return l
}

// visualTruncate truncates the string so that its "visible" length is lower or equal
// to the provided limit.
// When needed, visualTruncate terminates the string by an ansi.Reset sequence
// to inhibate any visual effects coming from the truncation step.
func visualTruncate(s string, limit int) (trunc string) {
	var l int
	var sgr ansi.Sequence

	_ = ansi.Walk(s, func(c rune, esc string) error {
		if c > -1 {
			trunc += string(c)
			l += runeWidth(c)
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

// visualPad completes a string with provided rune until its "visible" lentgh
// reaches the provided limit.  If the string visible size is already above
// imit, Pad returns it as-is.
func visualPad(s string, size int, padRune rune) string {
	var pad []rune
	for i := visualLen(s); i < size; i++ {
		pad = append(pad, padRune)
	}
	return s + string(pad)
}

// visualRepeat repeats s until given size is reached.
func visualRepeat(s string, size int) string {
	l := visualLen(s)
	r := s
	for i := l; i <= size; i = i + l {
		if i == size {
			return r
		}
		r = r + s
	}
	return visualTruncate(r, size)
}

// interruptANSI interrupts at each line any ANSI SGR rendition and continues it
// at the next line (useful to work with text in column to avoid voiding
// neighbourgh text).
func interruptANSI(s []string) {
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

// justifyWithInterruptANSI wraps cells, properly interrupting ANSI sequence at
// line boundaries, then fill lines to meet columns size
func justifyWithInterruptANSI(s string, sz int, truncateLongWords bool) string {
	if len(s) == 0 {
		return strings.Repeat(" ", sz)
	}

	ws := wrap(s, sz, truncateLongWords)
	interruptANSI(ws)
	for i, l := range ws {
		ws[i] = visualPad(l, sz, ' ')
	}
	return strings.Join(ws, "\n")
}

// runeWidth returns the visual width of a rune.
func runeWidth(c rune) int {
	return runewidth.RuneWidth(c)
}
