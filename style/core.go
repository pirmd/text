package style

import (
	"strings"
)

var (
	_ Styler = (*Core)(nil) //Makes sure that Core implements Styler
)

//Core is a minimal Styler that basically does nothing but providing functions
//that almost everybody wants to have (wrapping some functions from package
//'strings'). Using Core to format text will also probably lead to unreadable
//outputs.
//
//Its only use is probably to serve as a basis to developp new Stylers.
type Core struct{}

//Upper changes a string to upper case
func (st *Core) Upper(s string) string {
	return strings.ToUpper(s)
}

//Lower changes a string to lower case
func (st *Core) Lower(s string) string {
	return strings.ToLower(s)
}

//TitleCase changes all letters that begin words to their title case.
func (st *Core) TitleCase(s string) string {
	return strings.Title(s)
}

//Black does nothing (this style does not support colors)
func (st *Core) Black(s string) string {
	return s
}

//Red does nothing (this style does not support colors)
func (st *Core) Red(s string) string {
	return s
}

//Green does nothing (this style does not support colors)
func (st *Core) Green(s string) string {
	return s
}

//Yellow does nothing (this style does not support colors)
func (st *Core) Yellow(s string) string {
	return s
}

//Blue does nothing (this style does not support colors)
func (st *Core) Blue(s string) string {
	return s
}

//Magenta does nothing (this style does not support colors)
func (st *Core) Magenta(s string) string {
	return s
}

//Cyan does nothing (this style does not support colors)
func (st *Core) Cyan(s string) string {
	return s
}

//White does nothing (this style does not support colors)
func (st *Core) White(s string) string {
	return s
}

//Inverse does nothing (this style does not support colors)
func (st *Core) Inverse(s string) string {
	return s
}

//Bold does nothing (this style does not support emphasis)
func (st *Core) Bold(s string) string {
	return s
}

//Italic does nothing (this style does not support emphasis)
func (st *Core) Italic(s string) string {
	return s
}

//Underline does nothing (this style does not support emphasis)
func (st *Core) Underline(s string) string {
	return s
}

//Strike does nothing (this style does not support emphasis)
func (st *Core) Strike(s string) string {
	return s
}

//Tab does nothing (this style does not support text indenting/wrapping)
func (st *Core) Tab(lvl int) func(string) string {
	return func(s string) string {
		return s
	}
}

//Header returns text as a chapter's header
func (st *Core) Header(lvl int) func(s string) string {
	return func(s string) string { return s + "\n" }
}

//Paragraph returns text as a new paragraph.
func (st *Core) Paragraph(s string) string {
	return s + "\n"
}

//Line returns text as a new line.
func (st *Core) Line(s string) string {
	return s + "\n"
}

//List returns a new bullet-list. It returns one line per list item.
//This style does not support nested-list so level is not taken into account.
func (st *Core) List(lvl int) func(...string) string {
	return func(items ...string) string {
		return strings.Join(items, "\n")
	}
}

//ListItem returns a new bullet list's item. It returns the input preceded by a
//hyphen ("- ")
func (st *Core) ListItem(s string) string {
	//TODO(pirmd): maybe allows to differenciate bullets type ("-", "+", "*")?
	return "- " + s
}

//Define returns a term definition
func (st *Core) Define(term string, desc string) string {
	return term + ": " + desc + "\n"
}

//Table draws a basic table according to style's table drawing function. It
//returns one row per line with a "|"-separated columns for each row.
//It implies that tables with multiple-lines per columns will be awfully
//difficult to understand.
func (st *Core) Table(rows ...[]string) string {
	var r []string
	for _, row := range rows {
		r = append(r, strings.Join(row, " | "))
	}
	return strings.Join(r, "\n")
}

//Escape does nothing.
func (st *Core) Escape(s string) string {
	//XXX: implements Escape
	return s
}

//Auto auto-styles the input text.
func (st *Core) Auto(s string) string {
	//XXX: implements Auto
	return s
}
