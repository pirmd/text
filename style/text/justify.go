package text

import (
	"strings"
)

//ElipsisRune defines the rune that is append to truncated strings to
//visually indicates that a string has been truncated
var ElipsisRune = '\u2026'

//Truncate returns a string that is shorter or equal to the provided
//limit. The length limit is calculated ignoring non-visible runes and ansi
//coloring/styling escape sequences.
//
//Truncated strings are returned with ElipsisRune as their last rune.
func Truncate(s string, size int) string {
	if visualLen(s) > size {
		return visualTruncate(s, size-1) + string(ElipsisRune)
	}
	return s
}

//Justify wraps a text to the given maximum size and makes sure that
//returned lines are of exact provided length by padding them as needed.
func Justify(s string, sz int) string {
	if visualLen(s) == 0 {
		return strings.Repeat(" ", sz)
	}

	ws := wrap(s, sz)
	for i, l := range ws {
		ws[i] = visualPad(l, sz, ' ')
	}
	return strings.Join(ws, "\n")
}

//ExactSize returns a string of the exact given size either
//by padding or truncating it.
func ExactSize(s string, size int) string {
	if visualLen(s) > size {
		return visualTruncate(s, size-1) + string(ElipsisRune)
	}
	return visualPad(s, size, ' ')
}
