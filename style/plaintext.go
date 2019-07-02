package style

import (
	"strings"
	"unicode"

	"github.com/pirmd/cli/style/text"
)

var (
	//DefaultTxtWidth is the text width used to wrap text for markups that
	//supports it (PlainText notably)
	DefaultTxtWidth = 80

	//IndentPrefix is the suit of bytes used to indent text for markups that
	//supports it (PlainText notably) Use of '\t' is not recommended as it does
	//not cohabite correctly with wrapping func that cannot guess a '\t' length
	IndentPrefix = "    "

	indent2Prefix string //IndentPrefix x2
)

//core is a minimal styler providing function that almost everybody wants to
//have
var core = New(
	FormatMap{
		FmtUpper:            strings.ToUpper,
		FmtLower:            strings.ToLower,
		FmtTrimSpace:        strings.TrimSpace,
		FmtTrimLeadingSpace: func(s string) string { return strings.TrimLeftFunc(s, unicode.IsSpace) },
	},
	nil,
)

//PlainText is a Styler that provides a minimum style for plain texts. Wrap
//format wraps text to the maximum length specified by DefaultTxtWidth.
//
//Chaining multiples Wrap or Tab will, in most cases, void the result.
var PlainText = core.Extend(New(
	FormatMap{
		FmtDocHeader: Sprintf("%s\n"),
		FmtHeader:    Sprintf("\n%s\n"),
		FmtParagraph: Sprintf("\n%s\n"),
		FmtLine:      Sprintf("%s\n"),
		FmtList:      Sprintf("- %s\n"),
		FmtDefTerm:   Sprintf("\n%s:\n"),
		FmtDefDesc:   Sprintf("%s\n"),

		FmtWrap: func(s string) string { return text.Wrap(s, DefaultTxtWidth) },
		FmtTab:  func(s string) string { return text.Tab(s, IndentPrefix, DefaultTxtWidth) },
		FmtTab2: func(s string) string { return text.Tab(s, indent2Prefix, DefaultTxtWidth) },
	},

	//tableFn
	func(rows ...[]string) string { return "\n" + text.DrawTable(DefaultTxtWidth, " ", "-", " ", rows...) },
))

//XXX: add a Numbered Item
//XXX: create a Tab(n) function -> FormatFn

func init() {
	indent2Prefix = IndentPrefix + IndentPrefix
}
