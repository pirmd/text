package style

import (
	"fmt"
	"strings"
)

//Mandoc is a sub-set of troff markup featuring common used mdoc macro for
//building man pages
var Mandoc = core.Extend(New(
	FormatMap{
		FmtBold:   Sprintf("\\fB%s\\fP"),
		FmtItalic: Sprintf("\\fI%s\\fP"),

		FmtDocHeader: Sprintf(".TH %s\n"),
		FmtHeader:    Sprintf("\n.SH %s\n"),
		FmtHeader2:   Sprintf("\n.SS %s\n"),
		FmtHeader3:   Sprintf("\n.SS %s\n"),
		FmtParagraph: Sprintf(".PP\n%s\n"),
		FmtLine:      Sprintf("%s\n"),
		FmtDefTerm:   Sprintf("\n.TP\n%s\n"),
		FmtDefDesc:   Sprintf("%s\n"),
		FmtCode:      Sprintf(".PP\n.RS\n.nf\n%s\n.fi\n.RE\n"),

		//Needs mdoc macros set
		FmtList:      Sprintf("\n.Bl -dash\n%s.El\n"),
		FmtListItem:  Sprintf(".It\n%s\n"),
		FmtList2:     Sprintf(".Bl -bullet\n%s.El"),
		FmtList2Item: Sprintf(".It\n%s\n"),

		FmtEscape: escapeMandoc,
	},

	//tableFn
	func(rows ...[]string) string {
		if len(rows) == 0 {
			return ""
		}

		var r []string
		for _, row := range rows {
			for i, c := range row {
				row[i] = fmt.Sprintf("T{\n%s\nT}", c)
			}
			r = append(r, strings.Join(row, "\t"))
		}

		var layout []string
		for range rows[0] {
			layout = append(layout, "lx")
		}

		return fmt.Sprintf("\n.TS\nallbox;\n%s.\n%s\n.TE\n", strings.Join(layout, " "), strings.Join(r, "\n"))
	},
))

func escapeMandoc(s string) string {
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

	//Ugly trick to prevent valid mdoc macros options to be escaped when
	//styling functions are chained
	//TODO(pirmd): can something better be imagined?
	s = strings.ReplaceAll(s, ".Bl \\-bullet", ".Bl -bullet")
	s = strings.ReplaceAll(s, ".Bl \\-dash", ".Bl -dash")

	return s
}
