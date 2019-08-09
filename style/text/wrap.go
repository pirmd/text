package text

import (
	"strings"
	"unicode"
)

//Indent inserts a name/bullet/number at the beginning of the string, then
//indents it (add prefix at the beginning and before any new line).
//
//Tag is superposed to the indent prefix to obtain the first line prefix, if
//tag length is greater than prefix, prefix is completed by trailing spaces.
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

//Wrap wraps a text by ensuring that no text's line will be longer than the provided
//limit.
//
//Wrap calculates the visual width of a string, discarding non printable rune or ansi
//coloring/styling escape sequence.
//
//Wrap is pretty bad when called multiple times as it cannot differenciates line breaks
//coming from previous wrapping from line breaks introduced by the end-user in the
//first place. It is therfore better to avoid multiple calls to Wrap
//
//BUG: Wrap works at word limits (space) so it does not behaves properly for words
//longer than the limit (it does not split words so it feedbacks  line longer than
//limit)
func Wrap(txt string, limit int) string {
	return strings.Join(wrap(txt, limit), "\n")
}

//Tab wraps and indents the given text.
//
//Tab will aditionally add the given tag in front of the first line. Tag is
//superposed to the indent prefix to obtain the first line prefix, if tag's
//length is greater than prefix, prefix is completed by trailing spaces.
//
//Tab calculates the correct wraping limits taking indent's prefix length. It
//does not work if prefix is made of tabs as indent's tag/prefix length is
//unknown (like '\t').
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

func wrap(txt string, limit int) []string {
	var c rune
	var ws []string
	var line, word string
	var linelen int

	for _, c = range txt {
		switch {
		case c == '\n':
			wordlen := visualLen(word)
			if linelen+wordlen > limit {
				ws = append(ws, line)
				line, linelen = "", 0
			}

			ws = append(ws, line+word)
			word, line, linelen = "", "", 0

		case unicode.IsSpace(c):
			wordlen := visualLen(word)
			switch l := linelen + wordlen; {
			case linelen == 0 && l >= limit:
				ws = append(ws, word)
			case l > limit:
				ws = append(ws, line)
				line = word + string(c)
				linelen = wordlen + 1
			case l == limit:
				ws = append(ws, line+word)
				line, linelen = "", 0
			default:
				line += word + string(c)
				linelen += wordlen + 1
			}
			word = ""

		default:
			word = word + string(c)
		}
	}

	wordlen := visualLen(word)
	switch l := linelen + wordlen; {
	case l == 0 && c == '\n':
		return append(ws, "")
	case l == 0:
		return ws
	case l > limit && linelen == 0:
		return append(ws, word)
	case l > limit:
		return append(ws, line, word)
	default:
		return append(ws, line+word)
	}
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
