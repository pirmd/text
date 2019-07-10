package style

import (
	"strconv"
	"strings"

	"github.com/pirmd/cli/style/text"
)

var (
	_ Styler = (*TextSyntax)(nil) //Makes sure that TextSyntax implements Styler

	//Plaintext is a customized style.TextSyntax Styler to write plain text. It
	//allows for maximum 80 chars per line, indenting is made of 4 spaces and
	//list bullets are made of unicode hyphen and bullet.
	Plaintext = &TextSyntax{
		TextWidth:    80,
		IndentPrefix: "    ",
		ListBullets:  []string{"\u2043 ", "\u2022 ", "\u25E6 "},
	}
)

//TextSyntax implements Styler interface to provide basic formatting to write plain
//texts.  It supports text indenting and wraping as well as table but does not
//provide color nor text emphasis supports.
type TextSyntax struct {
	*CoreSyntax

	//TextWidth specified the maximum length of a text line
	//If st.TextWidth is null or negative, wraping is disabled (in practice, you
	//still want st.TextSyntaxwidth to be large enough to obtain a readable output).
	TextWidth int

	//IndentPrefix is the string used to indent text.
	//Use of "\t" is not recommended as it does not cohabite correctly with
	//wrapping func that cannot guess a '\t' length.
	//An empty string disable indenting and tabulation-like features.
	IndentPrefix string

	//ListBullets list the bullets added to each list items. Bullets are chosen
	//in the given order following the list nested-level (if nested-level is
	//greater than bullets number it restarts from 1).
	//If you want some spaces between the bullet and the start of text, you
	//have to include it (avoid "\t" nevertheless).
	ListBullets []string

	indentLvl   int
	enumerators []int

	//If true, adds a line break before paragraphs, headers or lists. It is
	//automatically set-up the first time one any of these formats is used.
	needBR bool
}

//Tab changes the tabulation level.
//If the tabulation level is positive, it wraps then indents provided text.
//Wraping is done according to st.TextWidth value, if st.TextWidth is null, Tab
//only indents and doesn't wrap the provided text.
//Indenting is done using st.IndentPrefix string, that is repeated for each
//tab-level.
func (st *TextSyntax) Tab(lvl int) func(string) string {
	oldlvl := st.indentLvl
	st.indentLvl = lvl
	return func(s string) string {
		st.indentLvl = oldlvl
		return s
	}
}

//Header returns text as a chapter's header.
func (st *TextSyntax) Header(lvl int) func(s string) string {
	switch {
	case lvl <= 0:
		return func(s string) string { return "" }
	case lvl == 1:
		return func(s string) string { return st.br() + st.Upper(s) + "\n" }
	default:
		return func(s string) string { return st.br() + st.TitleCase(s) }
	}
}

//Paragraph returns text as a new paragraph
func (st *TextSyntax) Paragraph(s string) string {
	return st.br() + st.tab(s, st.indentLvl, "") + "\n"
}

//List returns a new list (ordered or bulleted) with the proper nested level.
//List works in conjunction with either BulletedItem or OrderedItem.
func (st *TextSyntax) List(lvl int) func(...string) string {
	oldlvl := st.indentLvl
	st.indentLvl = lvl

	for st.indentLvl >= len(st.enumerators) {
		st.enumerators = append(st.enumerators, make([]int, 2)...)
	}

	return func(items ...string) string {
		st.enumerators[st.indentLvl] = 0
		st.indentLvl = oldlvl
		return strings.Join(items, "\n")
	}
}

//BulletedItem returns a new bullet-list item.
//It adds a bullet in front of the provided string according to st.BulletList
//and the list's level.
//A Tab is inserted before each item according to the list level.
func (st *TextSyntax) BulletedItem(s string) string {
	bullet := st.ListBullets[st.indentLvl%len(st.ListBullets)]
	return st.br() + st.tab(s, st.indentLvl, bullet)
}

//OrderedItem returns a new ordered-list item.
//It adds an enumerator in front of the provided string according to the
//current enumerator increment of the corresponding list's level (therefore
//only one enumerator can be followed for a given list's level).
//A Tab is inserted before each item according to the list level.
func (st *TextSyntax) OrderedItem(s string) string {
	st.enumerators[st.indentLvl]++
	enum := strconv.Itoa(st.enumerators[st.indentLvl]) + ". "
	return st.br() + st.tab(s, st.indentLvl, enum)
}

//Define returns a term definition
func (st *TextSyntax) Define(term string, desc string) string {
	term = st.tab(term, st.indentLvl, "") + "\n"
	desc = st.tab(desc, st.indentLvl+1, "") + "\n"
	return st.br() + term + desc
}

//Table draws a table out of the provided rows.
//Table column width are guessed automatically and are arranged so that the table
//fits into st.TextWidth.
func (st *TextSyntax) Table(rows ...[]string) string {
	//TODO(pirmd): ensure that IdentPrefix can work with ANSI code inside (notably
	//here where len is used, should be text.visualLen?)
	width := st.TextWidth - (st.indentLvl * len(st.IndentPrefix))
	//TODO(pirmd): introduce way to chose/define Table grid
	table := text.DrawTable(width, " ", "-", " ", rows...)

	return st.br() + st.tab(table, st.indentLvl, "") + "\n"
}

func (st *TextSyntax) br() string {
	if st.needBR {
		return "\n"
	}
	st.needBR = true
	return ""
}

func (st *TextSyntax) tab(s string, lvl int, tag string) string {
	prefix := strings.Repeat(st.IndentPrefix, lvl)

	if st.TextWidth > 0 {
		return text.Tab(s, tag, prefix, st.TextWidth)
	}
	return text.Indent(s, tag, prefix)
}
