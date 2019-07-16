package style

import (
	"strconv"
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
func (stx *CoreSyntax) Upper(s string) string {
	return strings.ToUpper(s)
}

//Lower changes a string to lower case
func (stx *CoreSyntax) Lower(s string) string {
	return strings.ToLower(s)
}

//TitleCase changes all letters that begin words to their title case.
func (stx *CoreSyntax) TitleCase(s string) string {
	return strings.Title(s)
}

//Black does nothing (this style does not support colors)
func (stx *CoreSyntax) Black(s string) string {
	return s
}

//Red does nothing (this style does not support colors)
func (stx *CoreSyntax) Red(s string) string {
	return s
}

//Green does nothing (this style does not support colors)
func (stx *CoreSyntax) Green(s string) string {
	return s
}

//Yellow does nothing (this style does not support colors)
func (stx *CoreSyntax) Yellow(s string) string {
	return s
}

//Blue does nothing (this style does not support colors)
func (stx *CoreSyntax) Blue(s string) string {
	return s
}

//Magenta does nothing (this style does not support colors)
func (stx *CoreSyntax) Magenta(s string) string {
	return s
}

//Cyan does nothing (this style does not support colors)
func (stx *CoreSyntax) Cyan(s string) string {
	return s
}

//White does nothing (this style does not support colors)
func (stx *CoreSyntax) White(s string) string {
	return s
}

//Inverse does nothing (this style does not support colors)
func (stx *CoreSyntax) Inverse(s string) string {
	return s
}

//Bold does nothing (this style does not support emphasis)
func (stx *CoreSyntax) Bold(s string) string {
	return s
}

//Italic does nothing (this style does not support emphasis)
func (stx *CoreSyntax) Italic(s string) string {
	return s
}

//Underline does nothing (this style does not support emphasis)
func (stx *CoreSyntax) Underline(s string) string {
	return s
}

//Crossout does nothing (this style does not support emphasis)
func (stx *CoreSyntax) Crossout(s string) string {
	return s
}

//Tab does nothing (this style does not support text indenting/wrapping)
func (stx *CoreSyntax) Tab(lvl int) func(string) string {
	return func(s string) string {
		return s
	}
}

//Header returns text as a chapter's header
func (stx *CoreSyntax) Header(lvl int) func(s string) string {
	switch {
	case lvl <= 0:
		return func(s string) string { return "" }
	default:
		return func(s string) string { return s + "\n" }
	}
}

//Metadata is not supported by this style
func (stx *CoreSyntax) Metadata(map[string]string) string {
	return ""
}

//Paragraph returns text as a new paragraph.
func (stx *CoreSyntax) Paragraph(s string) string {
	return s + "\n"
}

//BulletedList returns a new bulleted-list (each list item has a leading
//hyphen).
//This style does not support nested-list.
func (stx *CoreSyntax) BulletedList() func(items ...string) string {
	return func(items ...string) string {
		var s string
		for i, item := range items {
			if i == 0 {
				s = "- " + item
			} else {
				s = s + "\n" + "- " + item
			}
			if !strings.HasSuffix(s, "\n") {
				s += "\n"
			}
		}

		return "\n" + s
	}
}

//OrderedList returns a new ordered-list (each list item has a leading
//auto-incrementing enumerator).
//This style does not support nested-list.
func (stx *CoreSyntax) OrderedList() func(items ...string) string {
	return func(items ...string) string {
		var s string
		for i, item := range items {
			if i == 0 {
				s = "1. " + item
			} else {
				s = s + "\n" + strconv.Itoa(i+1) + ". " + item
			}
			if !strings.HasSuffix(s, "\n") {
				s += "\n"
			}
		}

		return "\n" + s
	}
}

//Define returns a term definition
func (stx *CoreSyntax) Define(term string, desc string) string {
	return term + ": " + desc + "\n"
}

//Table draws a basic table according to style's table drawing function. It
//returns one row per line with a "|"-separated columns for each row.
//It implies that tables with multiple-lines per columns will be awfully
//difficult to understand.
func (stx *CoreSyntax) Table(rows ...[]string) string {
	var r []string
	for _, row := range rows {
		r = append(r, strings.Join(row, " | "))
	}
	return strings.Join(r, "\n")
}

//Link returns links to internal or external resources
func (stx *CoreSyntax) Link(txt string, url string) string {
	return "[" + txt + "](" + url + ")"
}

//Img returns a string pointing to an image
func (stx *CoreSyntax) Img(txt string, url string) string {
	return "!" + stx.Link(txt, url)
}

//Escape escapes the provided text.
func (stx *CoreSyntax) Escape(s string) string {
	return s
}
