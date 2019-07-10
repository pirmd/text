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
		TextWidth:    80,
		IndentPrefix: "    ",
		ListBullets:  []string{"- ", "* ", "+ "},
	}}
)

//MarkdownSyntax implements Styler interface to provide formatting to write plain
//texts using the markdown idiom.
type MarkdownSyntax struct {
	*TextSyntax
}

//Header returns text as a chapter's header.
func (st *MarkdownSyntax) Header(lvl int) func(s string) string {
	switch {
	case lvl <= 0:
		return func(s string) string { return "" }
	case lvl == 1:
		return func(s string) string { return st.br() + "# " + st.Upper(s) + "\n" }
	default:
		hashes := strings.Repeat("#", lvl) + " "
		return func(s string) string { return st.br() + hashes + st.TitleCase(s) + "\n" }
	}
}

//Metadata returns formatted metadata information (title, author(s), date)
func (st *MarkdownSyntax) Metadata(title, authors, date string) string {
	return st.br() + "% " + title + "\n% " + authors + "\n% " + date + "\n"
}

//Bold changes a string case to bold
func (st *MarkdownSyntax) Bold(s string) string {
	return "__" + s + "__"
}

//Italic changes a string case to italic
func (st *MarkdownSyntax) Italic(s string) string {
	return "*" + s + "*"
}

//Crossout changes a string to be strikethrough
func (st *MarkdownSyntax) Crossout(s string) string {
	return "~~" + s + "~~"
}

//Define returns a term definition
func (st *MarkdownSyntax) Define(term string, desc string) string {
	return st.br() + st.tab(term+"\n:"+desc, st.indentLvl, "") + "\n"
}

//Escape escapes the provided text.
func (st *MarkdownSyntax) Escape(s string) string {
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
