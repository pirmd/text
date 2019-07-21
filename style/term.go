package style

import (
	"github.com/pirmd/cli/style/termsize"
)

var (
	_ Styler = (*ColorTextSyntax)(nil) //Makes sure that ColorTextSyntax implements Styler

	//Term is a customized style.TextSyntax Style to write plain text to the
	//terminal. It extends textSyntax by adapting text maximum length to the
	//terminal width. If terminal width cannot be detected, it will fallback to
	//a 80 maximum chars per line.
	Term = &TextSyntax{
		TextWidth:   termwidth(),
		TabWidth:    4,
		ListBullets: []string{"\u2043 ", "\u2022 ", "\u25E6 "},
	}

	//ColorTerm is a customized style.ColorTextSyntax Style to write plain text
	//in color to the terminal. It extends TextSyntax by adapting text maximum
	//length to the terminal width. If terminal width cannot be detected, it
	//will fallback to a 80 maximum chars per line.
	ColorTerm = &ColorTextSyntax{&TextSyntax{
		TextWidth:   termwidth(),
		TabWidth:    4,
		ListBullets: []string{"\u2043 ", "\u2022 ", "\u25E6 "},
	}}
)

//ColorTextSyntax implements Styler interface to provide basic formatting to write
//plain texts using ANSI colors. It supports in addition to ANSI colors and
//text emphasis, text indenting and wraping as well as table.
type ColorTextSyntax struct {
	*TextSyntax
}

//Black changes a string foreground color to black
func (stx *ColorTextSyntax) Black(s string) string {
	return "\x1b[30m" + s + "\x1b[0m"
}

//Red changes a string foreground color to red
func (stx *ColorTextSyntax) Red(s string) string {
	return "\x1b[31m" + s + "\x1b[0m"
}

//Green changes a string foreground color to green
func (stx *ColorTextSyntax) Green(s string) string {
	return "\x1b[32m" + s + "\x1b[0m"
}

//Yellow changes a string foreground color to yellow
func (stx *ColorTextSyntax) Yellow(s string) string {
	return "\x1b[33m" + s + "\x1b[0m"
}

//Blue changes a string foreground color to blue
func (stx *ColorTextSyntax) Blue(s string) string {
	return "\x1b[34m" + s + "\x1b[0m"
}

//Magenta changes a string foreground color to magenta
func (stx *ColorTextSyntax) Magenta(s string) string {
	return "\x1b[35m" + s + "\x1b[0m"
}

//Cyan changes a string foreground color to cyan
func (stx *ColorTextSyntax) Cyan(s string) string {
	return "\x1b[36m" + s + "\x1b[0m"
}

//White changes a string foreground color to white
func (stx *ColorTextSyntax) White(s string) string {
	return "\x1b[37m" + s + "\x1b[0m"
}

//Inverse changes a string by inverting its fore- and back-ground
//colors
func (stx *ColorTextSyntax) Inverse(s string) string {
	return "\x1b[7m" + s + "\x1b[27m"
}

//Bold changes a string case to bold
func (stx *ColorTextSyntax) Bold(s string) string {
	return "\x1b[1m" + s + "\x1b[22m"
}

//Italic changes a string case to italic
func (stx *ColorTextSyntax) Italic(s string) string {
	return "\x1b[3m" + s + "\x1b[23m"
}

//Underline changes a string to be underlined
func (stx *ColorTextSyntax) Underline(s string) string {
	return "\x1b[4m" + s + "\x1b[24m"
}

//Crossout changes a string to be strikethrough
func (stx *ColorTextSyntax) Crossout(s string) string {
	return "\x1b[9m" + s + "\x1b[29m"
}

//Define returns a term definition
func (stx *ColorTextSyntax) Define(term string, desc string) string {
	term = stx.tab(stx.Bold(term), stx.indentLvl, "") + "\n"
	desc = stx.tab(desc, stx.indentLvl+1, "") + "\n"
	return stx.br() + term + desc
}

func termwidth() int {
	const DefaultWidth = 80

	if w, err := termsize.Width(); err == nil {
		return w
	}
	return DefaultWidth
}
