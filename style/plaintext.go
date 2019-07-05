package style

import (
	"fmt"
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

	//ListBullets list the bullets to be added for each list items. Bullets are
	//chosen in that order per list nested-level
	ListBullets = [...]string{"-", "*", "+"}
	//ListBullets = [...]string{"\u2043", "\u2022", "\u25E6"}
)

//core is a minimal styler providing function that almost everybody wants to
//have
var core = New(
	FormatMap{
		FmtUpper:           strings.ToUpper,
		FmtLower:           strings.ToLower,
		FmtTitle:           strings.Title,
		FmtTrimSpace:       strings.TrimSpace,
		FmtNoLeadingSpace:  func(s string) string { return strings.TrimLeftFunc(s, unicode.IsSpace) },
		FmtNoTrailingSpace: func(s string) string { return strings.TrimRightFunc(s, unicode.IsSpace) },
	},
	nil,
	nil,
	nil,
	nil,
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
		FmtHeader2:   Sprintf("\n%s"),
		FmtHeader3:   Sprintf("\n%s"),
		FmtParagraph: Sprintf("\n%s\n"),
		FmtLine:      Sprintf("%s\n"),
		FmtWrap:      func(s string) string { return text.Wrap(s, DefaultTxtWidth) },
	},

	//tabFn
	func(level int) FormatFn {
		prefix := strings.Repeat(IndentPrefix, level)
		return func(s string) string {
			return text.Tab(s, prefix, DefaultTxtWidth)
		}
	},

	//listFn
	nil,

	//listItemFn
	func(level int) FormatFn {
		prefix := strings.Repeat(IndentPrefix, level)
		bullet := ListBullets[level%len(ListBullets)] + " "
		return func(s string) string {
			return "\n" + text.TabWithBullet(s, bullet, prefix, DefaultTxtWidth)
		}
	},

	//tableFn
	func(rows ...[]string) string { return "\n" + text.DrawTable(DefaultTxtWidth, " ", "-", " ", rows...) },

	//defineFn
	func(term, desc string) string {
		desc = text.Tab(desc, IndentPrefix, DefaultTxtWidth)
		return fmt.Sprintf("\n%s:\n%s\n", term, desc)
	},
))

//XXX: add a Numbered Item
