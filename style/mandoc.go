package style

import (
	"strconv"
	"strings"
)

//XXX st *XXXSyntax -> stx *XXXSyntax

var (
	_ Styler = (*ManSyntax)(nil)  //Makes sure that Man implements Styler

	//Man is a customized style.ManSyntax to write manpages.
	Man = &ManSyntax{
        ListBullets: []string{"\u2043 ", "\u2022 ", "\u25E6 "},
		IndentWidth: 4,
	}
)

//ManSyntax is a Styler that provides a sub-set of roff markup featuring common
//used macros for building man pages
type ManSyntax struct {
	*CoreSyntax

    //IndentWidth is the scaling factor (number of spaces) to indent text.
	IndentWidth int

	//ListBullets list the bullets added to each list items. Bullets are chosen
	//in the given order following the list nested-level (if nested-level is
	//greater than bullets number it restarts from 1).
	//If you want some spaces between the bullet and the start of text, you
	//have to include it (avoid "\t" nevertheless).
	ListBullets []string

	indentLvl   int
	enumerators []int
}

//Bold changes a string case to bold.
func (st *ManSyntax) Bold(s string) string {
	return "\\fB" + s + "\\fP"
}

//Italic changes a string case to italic.
func (st *ManSyntax) Italic(s string) string {
	return "\\fI" + s + "\\fP"
}

//Tab changes the tabulation level.
//If the tabulation level is positive, it wraps then indents provided text.
//Indenting is done using st.IndentWidth.
func (st *ManSyntax) Tab(lvl int) func(string) string {
	oldlvl := st.indentLvl
	st.indentLvl = lvl

	return func(s string) string {
		st.indentLvl = oldlvl
		return s
	}
}

//Header returns text as a chapter's header
func (st *ManSyntax) Header(lvl int) func(s string) string {
	switch {
	case lvl <= 0:
		return func(s string) string { return "" }
	case lvl == 1:
		return func(s string) string { return "\n.SH " + s + "\n" }
	default:
		return func(s string) string { return "\n.SS " + s + "\n" }
	}
}

//Metadata returns formatted metadata information (title, author(s), date)
//The man page section need to be found in title argument.
//XXX: Metadata(map[string]string) string
func (st *ManSyntax) Metadata(title, authors, date string) string {
	return ".TH " + st.Upper(title) + date + "\n"
}

//Paragraph returns text as a new paragraph.
func (st *ManSyntax) Paragraph(s string) string {
    return st.tab(".PP\n" + s + "\n")
}

//List returns a new bulleted-list. It returns one line per list item.
func (st *ManSyntax) List(lvl int) func(...string) string {
	oldlvl := st.indentLvl
	st.indentLvl = lvl + 1

	for st.indentLvl >= len(st.enumerators) {
		st.enumerators = append(st.enumerators, make([]int, 2)...)
	}

	return func(items ...string) string {
		st.enumerators[st.indentLvl] = 0
		st.indentLvl = oldlvl
		return st.tab(strings.Join(items, "\n"))
	}
}

//BulletedItem returns a new bullet-list item.
//It adds a bullet in front of the provided string according to st.BulletList
//and the list's level.
func (st *ManSyntax) BulletedItem(s string) string {
	bullet := st.ListBullets[st.indentLvl%len(st.ListBullets)]
    return ".TP " + strconv.Itoa(len(bullet)) + "\n" + bullet + "\n" + s + "\n"
}

//OrderedItem returns a new ordered-list item.
//It adds an enumerator in front of the provided string according to the
//current enumerator increment of the corresponding list's level (therefore
//only one enumerator can be followed for a given list's level).
func (st *ManSyntax) OrderedItem(s string) string {
	st.enumerators[st.indentLvl]++
	enum := strconv.Itoa(st.enumerators[st.indentLvl]) + ". "
    return ".TP " + strconv.Itoa(len(enum)) + "\n" + enum + "\n" + s + "\n"
}

//Define returns a term definition
func (st *ManSyntax) Define(term string, desc string) string {
    def := ".TP\n" + st.Bold(term) + "\n" + desc + "\n"
	return st.tab(def)
}

//Table draws a table out of the provided rows (using tbl).
func (st *ManSyntax) Table(rows ...[]string) string {
	if len(rows) == 0 {
		return ""
	}

	var r []string
	for _, row := range rows {
		for i, c := range row {
			row[i] = "T{\n" + c + "\nT}"
		}
		r = append(r, strings.Join(row, "\t"))
	}

	var layout []string
	for range rows[0] {
		layout = append(layout, "lx")
	}

    table := ".TS\nallbox;\n" + strings.Join(layout, " ") + ".\n" + strings.Join(r, "\n") + "\n.TE\n"
	return st.tab(table)
}

//Link is not supported in this style.
func (st *ManSyntax) Link(txt string, url string) string {
	return st.Italic(url)
}

//Img is not supported in this style.
func (st *ManSyntax) Img(txt string, url string) string {
	return st.Link(txt, url)
}

//Escape escapes the provided text.
func (st *ManSyntax) Escape(s string) string {
	var toEsc = [...]string{"-", "_", "&", "~"}

	//Assume that if supplied string contains already escaped char, it was
	//already escaped (chaining of styling's functions)
	for _, e := range toEsc {
		if strings.Contains(s, "\\"+e) {
			return s
		}
	}

	for _, e := range toEsc {
		s = strings.ReplaceAll(s, e, "\\"+e)
	}

	return s
}

func (st *ManSyntax) tab(s string) string {
    if st.indentLvl > 0 {
        return ".RS " + strconv.Itoa(st.indentLvl*st.IndentWidth) + "\n" + s + "\n.RE\n"
    }
    return s
}
