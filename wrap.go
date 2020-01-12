package text

import (
	"strings"
	"unicode"

	"github.com/pirmd/text/ansi"
)

// Indent inserts a name/bullet/number at the beginning of the string, then
// indents it (add prefix at the beginning and before any new line).
//
// Tag is superposed to the indent prefix to obtain the first line prefix, if
// tag length is greater than prefix, prefix is completed by trailing spaces.
func Indent(s string, tag, prefix string) string {
	lB, lP := visualLen(tag), visualLen(prefix)

	switch {
	case lB > lP:
		prefix = visualPad(prefix, lB, ' ')
	case lB < lP:
		tag = visualTruncate(prefix, lP-lB) + tag
	}

	return indent(s, tag, prefix)
}

// Wrap wraps a text by ensuring that each of its line's "visible" length is
// lower or equal to the provided limit. Wrap works at word limits (space). If
// a word is encountered that is longer than the limit, it is truncated.
func Wrap(txt string, limit int) string {
	return strings.Join(wrap(txt, limit), "\n")
}

// Tab wraps and indents the given text.
//
// Tab will aditionally add the given tag in front of the first line. Tag is
// superposed to the indent prefix to obtain the first line prefix, if tag's
// length is greater than prefix, prefix is completed by trailing spaces.
//
// Tab calculates the correct wraping limits taking indent's prefix length. It
// does not work if prefix is made of tabs as indent's tag/prefix length is
// unknown (like '\t').
func Tab(s string, tag, prefix string, limit int) string {
	lB, lP := visualLen(tag), visualLen(prefix)

	var r string
	switch {
	case lB > lP:
		prefix = visualPad(prefix, lB, ' ')
		r = Wrap(s, limit-lB)
	case lB < lP:
		tag = visualTruncate(prefix, lP-lB) + tag
		r = Wrap(s, limit-lP)
	default:
		r = Wrap(s, limit-lP)
	}

	return indent(r, tag, prefix)
}

func indent(s string, firstPrefix, prefix string) string {
	var indented string
	var isNewLine bool

	for _, c := range s {
		//add indent prefix if we have a non-empty newline
		if isNewLine && c != '\n' {
			indented = indented + prefix
		}
		indented += string(c)
		isNewLine = (c == '\n')
	}

	return firstPrefix + indented
}

func wrap(s string, limit int) (ws []string) {
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
				linelen = wordlen + runeWidth(c)

			default:
				line += word + string(c)
				linelen += wordlen + runeWidth(c)
			}
			word, wordlen = "", 0

		default:
			if wordlen += runeWidth(c); wordlen > limit {
				// word is longer than the limit, we truncated it
				if line != "" {
					ws = append(ws, line)
				}
				ws = append(ws, word)
				word, wordlen = "" , runeWidth(c)
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
