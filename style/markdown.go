package style

import (
	"strings"
)

var (
	_ Styler = (*MarkdownSyntax)(nil) //Makes sure that MarkdownSyntax implements Styler

	//Markdown is a customized style.MarkdownSyntax Styler to write plain text using
	//the markdown idiom.
	//Text wrapping is activated (80 chars maximum per line) so that reading
	//directly markdown text is easier to the eye.
	Markdown = &MarkdownSyntax{&TextSyntax{
		TextWidth:   80,
		TabWidth:    4,
		ListBullets: []string{"- ", "* ", "+ "},
	}}
)

//MarkdownSyntax implements Styler interface to provide formatting to write plain
//texts using the markdown idiom.
type MarkdownSyntax struct {
	*TextSyntax
}

//Header returns text as a chapter's header.
func (stx *MarkdownSyntax) Header(lvl int) func(s string) string {
	switch {
	case lvl <= 0:
		return func(s string) string { return "" }
	case lvl == 1:
		return func(s string) string { return stx.br() + "# " + stx.Upper(s) + "\n" }
	default:
		hashes := strings.Repeat("#", lvl) + " "
		return func(s string) string { return stx.br() + hashes + stx.TitleCase(s) + "\n" }
	}
}

//Metadata returns formatted metadata information
//Recognized metadata are: "title", "date", "authors"
func (stx *MarkdownSyntax) Metadata(mdata map[string]string) string {
	return stx.br() + "% " + mdata["title"] + "\n% " + mdata["authors"] + "\n% " + mdata["date"] + "\n"
}

//Bold changes a string case to bold
func (stx *MarkdownSyntax) Bold(s string) string {
	return "__" + s + "__"
}

//Italic changes a string case to italic
func (stx *MarkdownSyntax) Italic(s string) string {
	return "*" + s + "*"
}

//Crossout changes a string to be strikethrough
func (stx *MarkdownSyntax) Crossout(s string) string {
	return "~~" + s + "~~"
}

//Define returns a term definition
func (stx *MarkdownSyntax) Define(term string, desc string) string {
	return stx.br() + stx.tab(term+"\n:"+desc, stx.indentLvl, "") + "\n"
}

//Escape escapes the provided text.
func (stx *MarkdownSyntax) Escape(s string) string {
	var toEscape = [...]string{"\\", "`", "*", "_", "{", "}", "[", "]", "(", ")", ">", "#", "+", "-", ".", "!"}

	//Assume that if supplied string contains already escaped char, it was
	//already escaped (chaining of styling's functions)
	for _, e := range toEscape {
		if strings.Contains(s, "\\"+e) {
			return s
		}
	}

	for _, e := range toEscape {
		s = strings.ReplaceAll(s, e, "\\"+e)
	}

	return s
}
