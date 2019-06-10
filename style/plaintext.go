package style

import (
	"strings"

	"github.com/pirmd/cli/style/text"
)

var (
	//DefaultTxtWidth is the text width used to wrap text for markups that
	//supports it (PlainText notably)
	DefaultTxtWidth = 80

	//IndentPrefix is the suit of bytes used to indent text for markups that
	//supports it (PlainText notably) Use of '\t' is not recommended as it does
	//not cohabite correctly with wrapping func that cannot guess a '\t' length
	IndentPrefix = []byte("    ")

	indent2Prefix []byte //IndentPrefix x2
)

//core is a minimal styler providing function that almost everybody wants to
//have
var core = Styler{
	FmtUpper: strings.ToUpper,
	FmtLower: strings.ToLower,
}

//PlainText is a Styler that provides a minimum style for plain texts Wrap
//formatting wraps text to the maximum length of DefaultTxtWidth.
//
//Chaining multiples Wrap or Tab will in most cases void the result
var PlainText = core.Extend(Styler{
	FmtDocHeader: Sprintf("%s\n"),
	FmtHeader:    Sprintf("\n%s\n"),
	FmtParagraph: Sprintf("\n%s\n"),
	FmtLine:      Sprintf("%s\n"),
	FmtList:      Sprintf("\n- %s\n"),
	FmtDefTerm:   Sprintf("\n%s:\n"),
	FmtDefDesc:   Sprintf("%s\n"),

	FmtWrap: func(s string) string { return text.Wrap(s, DefaultTxtWidth) },
	FmtTab:  func(s string) string { return text.Tab(s, IndentPrefix, DefaultTxtWidth) },
	FmtTab2: func(s string) string { return text.Tab(s, indent2Prefix, DefaultTxtWidth) },
})

func init() {
	indent2Prefix = append(IndentPrefix, IndentPrefix...)
}
