package text

import (
	"unicode"
)

const ansiEscResetSeq = "\x1b[0m"

//TODO(pirmd): these functions calculate length of a text that contains ANSI
//            escape sequences that modify the text's style (color / bold and cie)
//            it will not filter properly the string if other ANSI sequences are
//            used, which should not be a problem for common use cases of this module.

func visualLen(s string) int {
	var isEscSeq bool
	var length int

	for _, c := range s {
		switch {
		case c == '\x1b' || c == '\x9b':
			isEscSeq = true

		case isEscSeq && c == 'm':
			isEscSeq = false

		case !isEscSeq && !unicode.IsMark(c) && unicode.IsGraphic(c):
			length++
		}
	}

	return length
}

//TODO(pirmd): when an ansi escape sequence are found, an ansi reset sequence
//is appended to the final truncated even if there is no specific need (reset
//sequence already exists)

func visualTruncate(s string, size int) string {
	var isEscSeq, hasEscSeq bool
	var length int

	for i, c := range s {
		switch {
		case c == '\x1b' || c == '\x9b':
			isEscSeq = true

		case isEscSeq && c == 'm':
			isEscSeq, hasEscSeq = false, true

		case !isEscSeq && !unicode.IsMark(c) && unicode.IsGraphic(c):
			length++
		}

		if length == size {
			i++
			if i == len(s) {
				return s
			}
			if hasEscSeq {
				return s[:i] + ansiEscResetSeq
			}
			return s[:i]
		}
	}

	return s
}

func visualPad(s string, size int, padRune rune) string {
	var pad []rune
	for i := visualLen(s); i < size; i++ {
		pad = append(pad, padRune)
	}
	return s + string(pad)
}
