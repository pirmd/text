package style

import (
	"fmt"
	"github.com/pirmd/cli/style/text"
)

//Markdown is a sub-set of markdown markup
var Markdown = PlainText.Extend(New(
	FormatMap{
		FmtBold:   Sprintf("**%s**"),
		FmtItalic: Sprintf("*%s*"),
		FmtStrike: Sprintf("~~%s~~"),

		FmtHeader:  Sprintf("\n# %s\n"),
		FmtHeader2: Sprintf("\n## %s\n"),
		FmtHeader3: Sprintf("\n### %s\n"),

		FmtCode: Sprintf("`%s`"),
	},
	nil,
	nil,
	nil,
	nil,

	//defineFn
	func(term, desc string) string {
		s := fmt.Sprintf("\n%s\n:%s\n", term, desc)
		return text.Wrap(s, DefaultTxtWidth)
	},
))

//XXX: Introduce Code and Bloc, transfer them to plaintext?
//XXX: Introduce escaping logic
