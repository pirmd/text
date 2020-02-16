package diff

import (
	"strings"
	"unicode"

	"github.com/pirmd/text/ansi"
)

const (
	tabAlias          = "\u21e5\u21e5\u21e5\u21e5"
	spaceAlias        = '\u00B7'
	nonPrintableAlias = '\ufffd'
)

var (
	// DefaultHighlighter is the highlither used by default by PrettyPrint if
	// no specific highlighters are supplied. You can ovverride it as needed.
	// It defaults to WithoutMissingContent.
	DefaultHighlighter = WithoutMissingContent

	// WithoutMissingContent highlights differences by hidding any missing part
	// on left or right.
	WithoutMissingContent = Highlighter{
		Same:      func(dL, dR, dT string) (string, string, string) { return dL, dR, dT },
		Deleted:   func(dL, dR, dT string) (string, string, string) { return dL, "", dT },
		Inserted:  func(dL, dR, dT string) (string, string, string) { return "", dR, dT },
		Different: func(dL, dR, dT string) (string, string, string) { return dL, dR, dT },
	}

	// WithColor highlights differences in colors: inserted are in red, deleted
	// are in blue, missing content are cossed-out.
	WithColor = Highlighter{
		Same: func(dL, dR, dT string) (string, string, string) { return dL, dR, dT },
		Deleted: func(dL, dR, dT string) (string, string, string) {
			return ansi.SetBlue(dL), ansi.SetCrossedOut(dR), ansi.SetBlueBG(ansi.SetWhite(dT))
		},
		Inserted: func(dL, dR, dT string) (string, string, string) {
			return ansi.SetCrossedOut(dL), ansi.SetRed(dR), ansi.SetRedBG(ansi.SetWhite(dT))
		},
		Different: func(dL, dR, dT string) (string, string, string) {
			return ansi.SetRed(dL), ansi.SetRed(dR), ansi.SetRedBG(ansi.SetWhite(dT))
		},
	}

	// WithNonPrintable ensures that non easily spotable differences are showed
	// by aliasing non visible runes with visible equivalent ones.
	WithNonPrintable = Highlighter{
		Same: func(dL, dR, dT string) (string, string, string) {
			return showNonPrintable(dL), showNonPrintable(dR), dT
		},
		Deleted: func(dL, dR, dT string) (string, string, string) {
			return showNonPrintable(dL), showNonPrintable(dR), dT
		},
		Inserted: func(dL, dR, dT string) (string, string, string) {
			return showNonPrintable(dL), showNonPrintable(dR), dT
		},
		Different: func(dL, dR, dT string) (string, string, string) {
			return showNonPrintable(dL), showNonPrintable(dR), dT
		},
	}

	// WithSoftTabs replaces any tabs ('\t') by four consecutives spaces so
	// that it does not voids any further text formatting (like showing diff in
	// columns)
	WithSoftTabs = Highlighter{
		Same:      func(dL, dR, dT string) (string, string, string) { return expandTabs(dL), expandTabs(dR), dT },
		Deleted:   func(dL, dR, dT string) (string, string, string) { return expandTabs(dL), expandTabs(dR), dT },
		Inserted:  func(dL, dR, dT string) (string, string, string) { return expandTabs(dL), expandTabs(dR), dT },
		Different: func(dL, dR, dT string) (string, string, string) { return expandTabs(dL), expandTabs(dR), dT },
	}
)

// Highlighter represents a set of functions to decorate a Diff when pretty
// printing it.
type Highlighter struct {
	Same, Deleted, Inserted, Different func(string, string, string) (string, string, string)
}

type highlighters []Highlighter

func newHighlighters(h ...Highlighter) highlighters {
	if len(h) > 0 {
		return highlighters(h)
	}
	return []Highlighter{DefaultHighlighter}
}

func (h highlighters) Format(delta Delta) (string, string, string) {
	diffL, diffR, diffT := delta.Value(), delta.Value(), delta.Type().String()
	for _, highlight := range h {
		switch delta.Type() {
		case IsSame:
			diffL, diffR, diffT = highlight.Same(diffL, diffR, diffT)
		case IsInserted:
			diffL, diffR, diffT = highlight.Inserted(diffL, diffR, diffT)
		case IsDeleted:
			diffL, diffR, diffT = highlight.Deleted(diffL, diffR, diffT)
		case IsDifferent:
			diffL, diffR, diffT = highlight.Different(diffL, diffR, diffT)
		}
	}

	return diffL, diffR, diffT
}

func showNonPrintable(s string) string {
	sWithSoftTabs := strings.ReplaceAll(s, "\t", tabAlias)

	return strings.Map(func(r rune) rune {
		switch {
		case r == '\n':
			return r

		case unicode.IsSpace(r):
			return spaceAlias

		case !unicode.IsPrint(r):
			return nonPrintableAlias

		default:
			return r
		}
	}, sWithSoftTabs)
}

func expandTabs(s string) string {
	return strings.ReplaceAll(s, "\t", "    ")
}
