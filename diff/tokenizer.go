package diff

import (
	"strings"
	"unicode"
)

// Tokenizer is a function that transform a string into a bunch of tokens
type Tokenizer func(string) []string

// ByLines splits a string by lines (limited by '\n').
func ByLines(s string) []string {

	// TODO(pirmd): PrettyPrint output is not properly working in case inputs
	// doas not end with a line break and is involved in an intermediate diff
	// operation. Therefore I have to add a fake end-of-line at this stage just
	// in case.  Something better has to be done to solve that
	if !strings.HasSuffix(s, "\n") {
		tokens := strings.SplitAfter(s+"\n", "\n")
		return tokens[:len(tokens)-1]
	}
	return strings.SplitAfter(s, "\n")
}

// ByRunes splits a string by runes.
func ByRunes(s string) []string {
	return strings.Split(s, "")
}

// ByWords splits a string by words (group of letters).
func ByWords(s string) (split []string) {
	var lasti int

	for i, r := range s {
		if !unicode.IsLetter(r) {
			if lasti != i {
				split = append(split, s[lasti:i])
			}
			split = append(split, string(r))
			lasti = i + 1
		}
	}

	if lasti < len(s) {
		split = append(split, s[lasti:])
	}

	return
}
