package style

import (
	"strings"
	"unicode"
)

//Markdown is a sub-set of markdown markup
var Markdown = core.Extend(Styler{
	FmtBold:   Sprintf("**%s**"),
	FmtItalic: Sprintf("*%s*"),
	FmtStrike: Sprintf("~~%s~~"),

	FmtHeader:    Sprintf("\n# %s"),
	FmtParagraph: Sprintf("\n%s\n"),
	FmtLine:      Sprintf("%s\n"),
	FmtList:      Sprintf("\n- %s\n"),
	FmtDefTerm:   Sprintf("\n%s\n"),
	FmtDefDesc:   Sprintf(": %s\n"),
	FmtCode:      Sprintf("`%s`"),

	FmtTrimSpaceLeft: func(s string) string { return strings.TrimLeftFunc(s, unicode.IsSpace) },
})
