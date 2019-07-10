package style

import (
	"strconv"
	"strings"
)

var (
	_ Styler = (*ManSyntax)(nil) //Makes sure that Man implements Styler

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
func (stx *ManSyntax) Bold(s string) string {
	return "\\fB" + s + "\\fP"
}

//Italic changes a string case to italic.
func (stx *ManSyntax) Italic(s string) string {
	return "\\fI" + s + "\\fP"
}

//Tab changes the tabulation level.
//If the tabulation level is positive, it wraps then indents provided text.
//Indenting is done using stx.IndentWidth.
func (stx *ManSyntax) Tab(lvl int) func(string) string {
	oldlvl := stx.indentLvl
	stx.indentLvl = lvl

	return func(s string) string {
		stx.indentLvl = oldlvl
		return s
	}
}

//Header returns text as a chapter's header
func (stx *ManSyntax) Header(lvl int) func(s string) string {
	switch {
	case lvl <= 0:
		return func(s string) string { return "" }
	case lvl == 1:
		return func(s string) string { return "\n.SH " + s + "\n" }
	default:
		return func(s string) string { return "\n.SS " + s + "\n" }
	}
}

//Metadata returns formatted metadata information.
//Used metadata are: "title", "date" and "mansection"
func (stx *ManSyntax) Metadata(mdata map[string]string) string {
	return ".TH " + stx.Upper(mdata["title"]) + " " + mdata["mansection"] + " " + mdata["date"] + "\n"
}

//Paragraph returns text as a new paragraph.
func (stx *ManSyntax) Paragraph(s string) string {
	return stx.tab(".PP\n" + s + "\n")
}

//List returns a new bulleted-listx. It returns one line per list item.
func (stx *ManSyntax) List(lvl int) func(...string) string {
	oldlvl := stx.indentLvl
	stx.indentLvl = lvl + 1

	for stx.indentLvl >= len(stx.enumerators) {
		stx.enumerators = append(stx.enumerators, make([]int, 2)...)
	}

	return func(items ...string) string {
		stx.enumerators[stx.indentLvl] = 0
		stx.indentLvl = oldlvl
		return stx.tab(strings.Join(items, "\n"))
	}
}

//BulletedItem returns a new bullet-list item.
//It adds a bullet in front of the provided string according to stx.BulletList
//and the list's level.
func (stx *ManSyntax) BulletedItem(s string) string {
	bullet := stx.ListBullets[stx.indentLvl%len(stx.ListBullets)]
	return ".TP " + strconv.Itoa(len(bullet)) + "\n" + bullet + "\n" + s + "\n"
}

//OrderedItem returns a new ordered-list item.
//It adds an enumerator in front of the provided string according to the
//current enumerator increment of the corresponding list's level (therefore
//only one enumerator can be followed for a given list's level).
func (stx *ManSyntax) OrderedItem(s string) string {
	stx.enumerators[stx.indentLvl]++
	enum := strconv.Itoa(stx.enumerators[stx.indentLvl]) + ". "
	return ".TP " + strconv.Itoa(len(enum)) + "\n" + enum + "\n" + s + "\n"
}

//Define returns a term definition
func (stx *ManSyntax) Define(term string, desc string) string {
	def := ".TP\n" + stx.Bold(term) + "\n" + desc + "\n"
	return stx.tab(def)
}

//Table draws a table out of the provided rows (using tbl).
func (stx *ManSyntax) Table(rows ...[]string) string {
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
	return stx.tab(table)
}

//Link is not supported in this style.
func (stx *ManSyntax) Link(txt string, url string) string {
	return stx.Italic(url)
}

//Img is not supported in this style.
func (stx *ManSyntax) Img(txt string, url string) string {
	return stx.Link(txt, url)
}

//Escape escapes the provided text.
func (stx *ManSyntax) Escape(s string) string {
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

func (stx *ManSyntax) tab(s string) string {
	if stx.indentLvl > 0 {
		return ".RS " + strconv.Itoa(stx.indentLvl*stx.IndentWidth) + "\n" + s + "\n.RE\n"
	}
	return s
}
