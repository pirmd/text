package style

import (
	"github.com/pirmd/cli/style/termsize"
	"github.com/pirmd/cli/style/text"
)

var (
	termWidth = DefaultTxtWidth
)

//Term extends PlainText markup to display texts on terminals
var Term = PlainText.Extend(New(
	FormatMap{
		FmtWrap: func(s string) string { return text.Wrap(s, termWidth) },
		FmtTab:  func(s string) string { return text.Tab(s, IndentPrefix, termWidth) },
		FmtTab2: func(s string) string { return text.Tab(s, indent2Prefix, termWidth) },
	},

	//tableFn
	func(rows ...[]string) string { return "\n" + text.DrawTable(termWidth, " ", "-", rows...) },
))

//ColorTerm extends PlainText markup with colors and text styles that can be
//understood by terminals that supports colors
var ColorTerm = Term.Extend(New(
	FormatMap{
		FmtBlack:     Sprintf("\x1b[30m%s\x1b[0m"),
		FmtRed:       Sprintf("\x1b[31m%s\x1b[0m"),
		FmtGreen:     Sprintf("\x1b[32m%s\x1b[0m"),
		FmtYellow:    Sprintf("\x1b[33m%s\x1b[0m"),
		FmtBlue:      Sprintf("\x1b[34m%s\x1b[0m"),
		FmtMagenta:   Sprintf("\x1b[35m%s\x1b[0m"),
		FmtCyan:      Sprintf("\x1b[36m%s\x1b[0m"),
		FmtWhite:     Sprintf("\x1b[37m%s\x1b[0m"),
		FmtBold:      Sprintf("\x1b[1m%s\x1b[22m"),
		FmtItalic:    Sprintf("\x1b[3m%s\x1b[23m"),
		FmtUnderline: Sprintf("\x1b[4m%s\x1b[24m"),
		FmtInverse:   Sprintf("\x1b[7m%s\x1b[27m"),
		FmtStrike:    Sprintf("\x1b[9m%s\x1b[29m"),

		FmtDefTerm: Sprintf("\n\x1b[1m%s\x1b[22m:\n"), //Defterm in Bold
	},
	nil,
))

func init() {
	if w, err := termsize.Width(); err == nil {
		termWidth = w
	}
}
