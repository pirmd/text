package style

import (
	"github.com/pirmd/cli/style/termsize"
)

var (
	_ Styler = (*ColorText)(nil) //Makes sure that ColorText implements Styler

	//Term is a customized style.Text Style to write plain text to the
	//terminal. It extends Plaintext style by adapting text maximum length to
	//the terminal width. If terminal width cannot be detected, it will
	//fallback to a 80 maximum chars per line.
	Term = &Text{
		TextWidth:    termwidth(),
		IndentPrefix: "    ",
		ListBullets:  []string{"\u2043 ", "\u2022 ", "\u25E6 "},
	}

	//ColorTerm is a customized style.ColorText Style to write plain text in
	//color to the terminal. It extends Plaintext style by adapting text
	//maximum length to the terminal width. If terminal width cannot be
	//detected, it will fallback to a 80 maximum chars per line.
	ColorTerm = &ColorText{&Text{
		TextWidth:    termwidth(),
		IndentPrefix: "    ",
		ListBullets:  []string{"\u2043 ", "\u2022 ", "\u25E6 "},
	}}
)

//ColorText implements Styler interface to provide basic formatting to write
//plain texts using ANSI colors. It supports in addition to ANSI colors and
//text emphasis, text indenting and wraping as well as table.
type ColorText struct {
	*Text
}

//Black changes a string foreground color to black
func (st *ColorText) Black(s string) string {
	return "\x1b[30m" + s + "\x1b[0m"
}

//Red changes a string foreground color to red
func (st *ColorText) Red(s string) string {
	return "\x1b[31m" + s + "\x1b[0m"
}

//Green changes a string foreground color to green
func (st *ColorText) Green(s string) string {
	return "\x1b[32m" + s + "\x1b[0m"
}

//Yellow changes a string foreground color to yellow
func (st *ColorText) Yellow(s string) string {
	return "\x1b[33m" + s + "\x1b[0m"
}

//Blue changes a string foreground color to blue
func (st *ColorText) Blue(s string) string {
	return "\x1b[34m" + s + "\x1b[0m"
}

//Magenta changes a string foreground color to magenta
func (st *ColorText) Magenta(s string) string {
	return "\x1b[35m" + s + "\x1b[0m"
}

//Cyan changes a string foreground color to cyan
func (st *ColorText) Cyan(s string) string {
	return "\x1b[36m" + s + "\x1b[0m"
}

//White changes a string foreground color to white
func (st *ColorText) White(s string) string {
	return "\x1b[37m" + s + "\x1b[0m"
}

//Inverse changes a string by inverting its fore- and back-ground
//colors
func (st *ColorText) Inverse(s string) string {
	return "\x1b[7m" + s + "\x1b[27m"
}

//Bold changes a string case to bold
func (st *ColorText) Bold(s string) string {
	return "\x1b[1m" + s + "\x1b[22m"
}

//Italic changes a string case to italic
func (st *ColorText) Italic(s string) string {
	return "\x1b[3m" + s + "\x1b[23m"
}

//Underline changes a string to be underlined
func (st *ColorText) Underline(s string) string {
	return "\x1b[4m" + s + "\x1b[24m"
}

//Crossout changes a string to be strikethrough
func (st *ColorText) Crossout(s string) string {
	return "\x1b[9m" + s + "\x1b[29m"
}

//Define returns a term definition
func (st *ColorText) Define(term string, desc string) string {
	term = st.tab(st.Bold(term), st.indentLvl, "") + "\n"
	desc = st.tab(desc, st.indentLvl+1, "") + "\n"
	return st.br() + term + desc
}

func termwidth() int {
	const DefaultWidth = 80

	if w, err := termsize.Width(); err == nil {
		return w
	}
	return DefaultWidth
}
