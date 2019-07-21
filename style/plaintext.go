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
		TextWidth:   80,
		TabWidth:    4,
		ListBullets: []string{"\u2043 ", "\u2022 ", "\u25E6 "},
	}
)

//TextSyntax implements Styler interface to provide basic formatting to write plain
//texts.  It supports text indenting and wraping as well as table but does not
//provide color nor text emphasis supports.
type TextSyntax struct {
	*CoreSyntax

	//TextWidth specified the maximum length of a text line
	//If stx.TextWidth is null or negative, wraping is disabled (in practice, you
	//still want stx.TextSyntaxwidth to be large enough to obtain a readable output).
	TextWidth int

	//TabWidth is the number of spaces added in front of text for each
	//tabulation's level
	TabWidth int

	//ListBullets list the bullets added to each list items. Bullets are chosen
	//in the given order following the list nested-level (if nested-level is
	//greater than bullets number it restarts from 1).
	//If you want some spaces between the bullet and the start of text, you
	//have to include it (avoid "\t" nevertheless).
	ListBullets []string

	indentLvl int
	nestLvl   int

	//If true, adds a line break before paragraphs, headers or lists. It is
	//automatically set-up the first time one any of these formats is used.
	needBR bool
}

//Tab increases the tabulation level for the provided text.
//Wraping is done according to stx.TextWidth value, if stx.TextWidth is null, Tab
//only indents and doesn't wrap the provided text.
func (stx *TextSyntax) Tab() func(string) string {
	stx.indentLvl++
	return func(s string) string {
		stx.indentLvl--
		return s
	}
}

//Header returns text as a chapter's header.
func (stx *TextSyntax) Header(lvl int) func(s string) string {
	switch {
	case lvl <= 0:
		return func(s string) string { return "" }
	case lvl == 1:
		return func(s string) string { return stx.br() + stx.Upper(s) + "\n" }
	default:
		return func(s string) string { return stx.br() + stx.TitleCase(s) }
	}
}

//Paragraph returns text as a new paragraph
func (stx *TextSyntax) Paragraph(s string) string {
	return stx.br() + stx.tab(s, stx.indentLvl, "") + "\n"
}

//BulletedList returns a new bulleted-list (each list item has a leading
//bullet).
//It automatically indents each item.
func (stx *TextSyntax) BulletedList() func(items ...string) string {
	stx.nestLvl++
	stx.indentLvl++

	bullet := stx.ListBullets[stx.indentLvl%len(stx.ListBullets)]

	return func(items ...string) string {
		var s string
		for i, item := range items {
			if i == 0 {
				s = stx.tab(item, stx.indentLvl, bullet)
			} else {
				s = s + "\n" + stx.tab(item, stx.indentLvl, bullet)
			}
			if !strings.HasSuffix(s, "\n") {
				s += "\n"
			}
		}

		stx.indentLvl--
		stx.nestLvl--
		return stx.br() + s
	}
}

//OrderedList returns a new ordered-list (each list item has a leading
//auto-incrementing enumerator).
//It automatically indents each item.
func (stx *TextSyntax) OrderedList() func(items ...string) string {
	stx.nestLvl++
	stx.indentLvl++

	return func(items ...string) string {
		var s string
		for i, item := range items {
			enum := strconv.Itoa(i+1) + ". "

			if i == 0 {
				s = stx.tab(item, stx.indentLvl, enum)
			} else {
				s = s + "\n" + stx.tab(item, stx.indentLvl, enum)
			}
			if !strings.HasSuffix(s, "\n") {
				s += "\n"
			}
		}

		stx.indentLvl--
		stx.nestLvl--
		return stx.br() + s
	}
}

//Define returns a term definition
func (stx *TextSyntax) Define(term string, desc string) string {
	term = stx.tab(term, stx.indentLvl, "") + "\n"
	desc = stx.tab(desc, stx.indentLvl+1, "") + "\n"
	return stx.br() + term + desc
}

//Table draws a table out of the provided rows.
//Table column width are guessed automatically and are arranged so that the table
//fits into stx.TextWidth.
func (stx *TextSyntax) Table(rows ...[]string) string {
	width := stx.TextWidth - (stx.indentLvl * stx.TabWidth)
	//TODO(pirmd): introduce way to chose/define Table grid
	table := text.DrawTable(width, " ", "-", " ", rows...)

	return stx.br() + stx.tab(table, stx.indentLvl, "") + "\n"
}

func (stx *TextSyntax) br() string {
	if stx.needBR {
		return "\n"
	}
	stx.needBR = true
	return ""
}

func (stx *TextSyntax) tab(s string, lvl int, tag string) string {
	prefix := strings.Repeat(" ", lvl*stx.TabWidth)

	if stx.TextWidth > 0 {
		if stx.nestLvl > 0 {
			width := stx.TextWidth - (stx.nestLvl-1)*stx.TabWidth
			prefix = strings.Repeat(" ", (lvl-stx.nestLvl)*stx.TabWidth)
			return text.Tab(s, tag, prefix, width)
		}
		return text.Tab(s, tag, prefix, stx.TextWidth)
	}
	return text.Indent(s, tag, prefix)
}
