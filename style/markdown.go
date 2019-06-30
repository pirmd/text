package style

import (
	"github.com/pirmd/cli/style/text"
)

//Markdown is a sub-set of markdown markup
var Markdown = core.Extend(New(
	FormatMap{
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
	},

	//tableFn
	func(rows ...[]string) string { return "\n" + text.DrawTable(DefaultTxtWidth, " ", "-", rows...) },
))
