package text

import (
	"strings"

	"github.com/pirmd/text/visual"
)

// Indent inserts a name/bullet/number at the beginning of the string, then
// indents it (add prefix at the beginning and before any new line).
//
// Tag is superposed to the indent prefix to obtain the first line prefix, if
// tag length is greater than prefix, prefix is completed by trailing spaces.
func Indent(s string, tag, prefix string) string {
	lB, lP := visual.Len(tag), visual.Len(prefix)

	switch {
	case lB > lP:
		prefix = visual.Pad(prefix, lB, ' ')
	case lB < lP:
		tag = visual.Truncate(prefix, lP-lB) + tag
	}

	return indent(s, tag, prefix)
}

// Wrap wraps a text by ensuring that each of its line's "visible" length is
// lower or equal to the provided limit. Wrap works with word limits being
// spaces.
//
// If a "word" is encountered that is longer than the limit, it is truncated or
// left as is depending of truncateLongWords flag.
func Wrap(txt string, limit int, truncateLongWords bool) string {
	return strings.Join(visual.Wrap(txt, limit, truncateLongWords), "\n")
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
func Tab(s string, tag, prefix string, limit int, truncateLongWords bool) string {
	lB, lP := visual.Len(tag), visual.Len(prefix)

	var r string
	switch {
	case lB > lP:
		prefix = visual.Pad(prefix, lB, ' ')
		r = Wrap(s, limit-lB, truncateLongWords)
	case lB < lP:
		tag = visual.Truncate(prefix, lP-lB) + tag
		r = Wrap(s, limit-lP, truncateLongWords)
	default:
		r = Wrap(s, limit-lP, truncateLongWords)
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
