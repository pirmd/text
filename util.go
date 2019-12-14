package text

import (
	"github.com/mattn/go-runewidth"

	"github.com/pirmd/text/ansi"
)

// visualLen calculates the length of string ignoring any ANSI escape sequences
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
//
// When needed, Truncate terminates the string by an ansi.Reset sequence to
// inhibate any visual effects coming from the truncation step.
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
// imit, Pad returns it as-is
func visualPad(s string, size int, padRune rune) string {
	var pad []rune
	for i := visualLen(s); i < size; i++ {
		pad = append(pad, padRune)
	}
	return s + string(pad)
}

// visualRepeat repeats s until given size is reached
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

// interruptANSI interrupts at each line any ANSI SGR rendition and continue it
// at the next line (usefull to work with text in column to avoid voiding
// neighbourgh text)
func interruptANSI(s []string) {
	var sgr ansi.Sequence

	for i, line := range s {
		curSGR := sgr
		_ = ansi.Walk(line, func(c rune, esc string) error {
			if c == -1 {
				curSGR.Combine(esc)
			}
			return nil
		})
		s[i] = sgr.Esc() + s[i] + curSGR.Reset()
		sgr = curSGR
	}
}

func runeWidth(c rune) int {
	return runewidth.RuneWidth(c)
}
