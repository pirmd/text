package style

import (
	"strings"
)

var (
	_ Styler = (*CoreSyntax)(nil) //Makes sure that CoreSyntax implements Styler
)

//CoreSyntax is a minimal Styler that basically does nothing but providing functions
//that almost everybody wants to have (wrapping some functions from package
//'strings'). Using CoreSyntax to format text will also probably lead to unreadable
//outputs.
//
//Its only use is probably to serve as a basis to developp new Stylers.
type CoreSyntax struct{}

//Upper changes a string to upper case
func (st *CoreSyntax) Upper(s string) string {
	return strings.ToUpper(s)
}

//Lower changes a string to lower case
func (st *CoreSyntax) Lower(s string) string {
	return strings.ToLower(s)
}

//TitleCase changes all letters that begin words to their title case.
func (st *CoreSyntax) TitleCase(s string) string {
	return strings.Title(s)
}

//Black does nothing (this style does not support colors)
func (st *CoreSyntax) Black(s string) string {
	return s
}

//Red does nothing (this style does not support colors)
func (st *CoreSyntax) Red(s string) string {
	return s
}

//Green does nothing (this style does not support colors)
func (st *CoreSyntax) Green(s string) string {
	return s
}

//Yellow does nothing (this style does not support colors)
func (st *CoreSyntax) Yellow(s string) string {
	return s
}

//Blue does nothing (this style does not support colors)
func (st *CoreSyntax) Blue(s string) string {
	return s
}

//Magenta does nothing (this style does not support colors)
func (st *CoreSyntax) Magenta(s string) string {
	return s
}

//Cyan does nothing (this style does not support colors)
func (st *CoreSyntax) Cyan(s string) string {
	return s
}

//White does nothing (this style does not support colors)
func (st *CoreSyntax) White(s string) string {
	return s
}

//Inverse does nothing (this style does not support colors)
func (st *CoreSyntax) Inverse(s string) string {
	return s
}

//Bold does nothing (this style does not support emphasis)
func (st *CoreSyntax) Bold(s string) string {
	return s
}

//Italic does nothing (this style does not support emphasis)
func (st *CoreSyntax) Italic(s string) string {
	return s
}

//Underline does nothing (this style does not support emphasis)
func (st *CoreSyntax) Underline(s string) string {
	return s
}

//Crossout does nothing (this style does not support emphasis)
func (st *CoreSyntax) Crossout(s string) string {
	return s
}

//Tab does nothing (this style does not support text indenting/wrapping)
func (st *CoreSyntax) Tab(lvl int) func(string) string {
	return func(s string) string {
		return s
	}
}

//Header returns text as a chapter's header
func (st *CoreSyntax) Header(lvl int) func(s string) string {
	switch {
	case lvl <= 0:
		return func(s string) string { return "" }
	default:
		return func(s string) string { return s + "\n" }
	}
}

//Metadata is not supported by this style
func (st *CoreSyntax) Metadata(title, authors, date string) string {
	return ""
}

//Paragraph returns text as a new paragraph.
func (st *CoreSyntax) Paragraph(s string) string {
	return s + "\n"
}

//List returns a new bulleted-list. It returns one line per list item.
func (st *CoreSyntax) List(lvl int) func(...string) string {
	return func(items ...string) string {
		return strings.Join(items, "\n")
	}
}

//BulletedItem returns a new bullet-list item.
//It adds an hyphen in front of each item. This style does not support
//nested-list so level is not taken into account.
func (st *CoreSyntax) BulletedItem(s string) string {
	return "- " + s
}

//OrderedItem returns a new ordered-list item.
//It adds a "#" in front of each item. This style does not support
//nested-list so level is not taken into account.
func (st *CoreSyntax) OrderedItem(s string) string {
	return "# " + s
}

//Define returns a term definition
func (st *CoreSyntax) Define(term string, desc string) string {
	return term + ": " + desc + "\n"
}

//Table draws a basic table according to style's table drawing function. It
//returns one row per line with a "|"-separated columns for each row.
//It implies that tables with multiple-lines per columns will be awfully
//difficult to understand.
func (st *CoreSyntax) Table(rows ...[]string) string {
	var r []string
	for _, row := range rows {
		r = append(r, strings.Join(row, " | "))
	}
	return strings.Join(r, "\n")
}

//Link returns links to internal or external resources
func (st *CoreSyntax) Link(txt string, url string) string {
	return "[" + txt + "](" + url + ")"
}

//Img returns a string pointing to an image
func (st *CoreSyntax) Img(txt string, url string) string {
	return "!" + st.Link(txt, url)
}

//Escape escapes the provided text.
func (st *CoreSyntax) Escape(s string) string {
	return s
}
