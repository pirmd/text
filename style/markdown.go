package style

//Markdown is a sub-set of markdown markup
var Markdown = PlainText.Extend(New(
	FormatMap{
		FmtBold:   Sprintf("**%s**"),
		FmtItalic: Sprintf("*%s*"),
		FmtStrike: Sprintf("~~%s~~"),

		FmtHeader: Sprintf("\n# %s\n"),
		FmtCode:   Sprintf("`%s`"),
	},
	nil,
))

//XXX: Introduce Code and Bloc, transfer them to plaintext?
//XXX: Introduce escaping logic
