package style

//Markdown is a sub-set of markdown markup
var Markdown = core.Extend(Styler{
	FmtBold:   Sprintf("**%s**"),
	FmtItalic: Sprintf("*%s*"),

	FmtHeader:    Sprintf("\n# %s"),
	FmtParagraph: Sprintf("\n%s\n"),
	FmtLine:      Sprintf("%s\n"),
	FmtList:      Sprintf("\n- %s\n"),
	FmtDefTerm:   Sprintf("\n%s\n"),
	FmtDefDesc:   Sprintf(": %s\n"),
	FmtCode:      Sprintf("`%s`"),
})
