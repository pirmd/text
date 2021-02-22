package util

import (
	"github.com/mattn/go-runewidth"

	"github.com/pirmd/text/ansi"
)

// VisualLen calculates the length of string ignoring any ANSI escape sequences.
func VisualLen(s string) int {
	var l int
	_ = ansi.Walk(s, func(c rune, esc string) error {
		if c > -1 {
			l += Runewidth(c)
		}
		return nil
	})

	return l
}

// VisualTruncate truncates the string so that its "visible" length is lower or equal
// to the provided limit.
// When needed, visualTruncate terminates the string by an ansi.Reset sequence
// to inhibate any visual effects coming from the truncation step.
func VisualTruncate(s string, limit int) (trunc string) {
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

// VisualPad completes a string with provided rune until its "visible" lentgh
// reaches the provided limit.  If the string visible size is already above
// imit, Pad returns it as-is.
func VisualPad(s string, size int, padRune rune) string {
	var pad []rune
	for i := VisualLen(s); i < size; i++ {
		pad = append(pad, padRune)
	}
	return s + string(pad)
}

// VisualRepeat repeats s until given size is reached.
func VisualRepeat(s string, size int) string {
	l := VisualLen(s)
	r := s
	for i := l; i <= size; i = i + l {
		if i == size {
			return r
		}
		r = r + s
	}
	return VisualTruncate(r, size)
}

// InterruptANSI interrupts at each line any ANSI SGR rendition and continues it
// at the next line (useful to work with text in column to avoid voiding
// neighbourgh text).
func InterruptANSI(s []string) {
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

// Runewidth returns the visual width of a rune.
func Runewidth(c rune) int {
	return runewidth.RuneWidth(c)
}
