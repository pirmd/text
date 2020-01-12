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
	// WithColor highlights differences in colors
	WithColor = Highlighter{
		Same:     func(dL, dR string) (string, string) { return dL, dR },
		Deleted:  func(dL, dR string) (string, string) { return ansi.SetBlue(dL), dR },
		Inserted: func(dL, dR string) (string, string) { return dL, ansi.SetRed(dR) },
		Type: func(dT Type, diffT string) string {
			switch dT {
			case IsInserted, IsDifferent:
				return ansi.SetRedBG(ansi.SetWhite(diffT))
			case IsDeleted:
				return ansi.SetBlueBG(ansi.SetWhite(diffT))
			}
			return diffT
		},
	}

	// WithNonPrintable ensures that non easily spotable differences are showed
	// by aliasing non visible runes with visible equivalent ones.
	WithNonPrintable = Highlighter{
		Same:     func(dL, dR string) (string, string) { return showNonPrintable(dL), showNonPrintable(dR) },
		Deleted:  func(dL, dR string) (string, string) { return showNonPrintable(dL), showNonPrintable(dR) },
		Inserted: func(dL, dR string) (string, string) { return showNonPrintable(dL), showNonPrintable(dR) },
		Type:     func(dT Type, diffT string) string { return diffT },
	}

	// WithSoftTabs replaces any tabs ('\t') by four consecutives spaces so
	// that it does not voids any further text formatting (like showing diff in
	// columns)
	WithSoftTabs = Highlighter{
		Same:     func(dL, dR string) (string, string) { return expandTabs(dL), expandTabs(dR) },
		Deleted:  func(dL, dR string) (string, string) { return expandTabs(dL), expandTabs(dR) },
		Inserted: func(dL, dR string) (string, string) { return expandTabs(dL), expandTabs(dR) },
		Type:     func(dT Type, diffT string) string { return diffT },
	}
)

// Highlighter represents a set of functions to decorate a Diff when pretty
// printing it.
type Highlighter struct {
	Same, Deleted, Inserted func(string, string) (string, string)
	Type                    func(Type, string) string
}

type highlighters []Highlighter

func (h highlighters) formatLR(dL, dR string, dT Type) (string, string) {
	diffL, diffR := dL, dR
	for _, highlight := range h {
		switch dT {
		case IsSame:
			diffL, diffR = highlight.Same(diffL, diffR)
		case IsInserted:
			diffL, diffR = highlight.Inserted(diffL, diffR)
		case IsDeleted:
			diffL, diffR = highlight.Deleted(diffL, diffR)
		}
	}

	return diffL, diffR
}

func (h highlighters) formatT(dT Type) string {
	diffT := dT.String()
	for _, highlight := range h {
		diffT = highlight.Type(dT, diffT)
	}
	return diffT
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
