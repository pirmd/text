package style

import (
	"fmt"
	"strings"
)

//Man is a sub-set of roff markup featuring common used macro for building
//man pages
var Man = core.Extend(New(
	FormatMap{
		FmtBold:   Sprintf("\\fB%s\\fP"),
		FmtItalic: Sprintf("\\fI%s\\fP"),

		FmtDocHeader: Sprintf(".TH %s\n"),
		FmtHeader:    Sprintf("\n.SH %s\n"),
		FmtHeader2:   Sprintf("\n.SS %s\n"),
		FmtHeader3:   Sprintf("\n.SS %s\n"),
		FmtParagraph: Sprintf(".PP\n%s\n"),
		FmtCode:      Sprintf(".PP\n.RS\n.nf\n%s\n.fi\n.RE\n"),

		FmtLine:   Sprintf("%s\n"),
		FmtEscape: escapeMandoc,
	},

	//tabFn
	func(level int) FormatFn {
		return func(s string) string {
			return fmt.Sprintf(".RS %d\n%s\n.RE\n", level*5, s)
		}
	},

	//listFn -> needs mdoc macros set
	func(level int) func(...string) string {
		return func(items ...string) string {
			return fmt.Sprintf(".RS %d\n%s\n.RE\n", level*5, strings.Join(items, "\n"))
		}
	},

	//listItemFn -> needs mdoc macros set
	func(level int) FormatFn {
		bullet := [...]string{"\u2043", "\u2022", "\u25E6"}[level%len(ListBullets)] + " "
		return func(s string) string {
			return fmt.Sprintf(".TP %d\n%s\n%s\n", len(bullet), bullet, s)
		}
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

	//defineFn
	func(term, desc string) string {
		return fmt.Sprintf("\n.TP\n\\fB%s\\fP\n%s\n", term, desc)
	},
))

//Mdoc is a sub-set of roff markup featuring common used mdoc macro for
//building man pages
var Mdoc = Man.Extend(New(
	FormatMap{
		FmtDocHeader: Sprintf(".Dt %s\n"),
		FmtHeader:    Sprintf("\n.Sh %s\n"),
		FmtHeader2:   Sprintf("\n.Ss %s\n"),
		FmtHeader3:   Sprintf("\n.Ss %s\n"),
		FmtParagraph: Sprintf(".Pp\n%s\n"),
		FmtCode:      Sprintf(".Pp\n.Rs\n.nf\n%s\n.fi\n.Re\n"),

		FmtEscape: escapeMdoc,
	},

	//tabFn
	func(level int) FormatFn {
		prefix := strings.Repeat("indent", level)
		return func(s string) string {
			return fmt.Sprintf(".Bd -ragged -offset %s\n%s\n.Ed\n", prefix, s)
		}
	},

	//listFn -> needs mdoc macros set
	func(level int) func(...string) string {
		bullet := [...]string{"-dash", "-bullet"}[level%2] + " "
		return func(items ...string) string {
			return fmt.Sprintf("\n.Bl %s\n%s.El\n", bullet, strings.Join(items, "\n"))
		}
	},

	//listItemFn -> needs mdoc macros set
	func(level int) FormatFn {
		return func(s string) string {
			return fmt.Sprintf(".It\n%s\n", s)
		}
	},

	//tableFn -> needs tbl
	nil,

	//defineFn
	func(term, desc string) string {
		return fmt.Sprintf("\n.Bl -tag\n.It \\fB%s\\fP\n%s\n.El\n", term, desc)
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

	return s
}

func escapeMdoc(s string) string {
	s = escapeMandoc(s)

	//Ugly trick to prevent valid mdoc macros options to be escaped when
	//styling functions are chained
	//TODO(pirmd): can something better be imagined?
	s = strings.ReplaceAll(s, ".Bl \\-bullet", ".Bl -bullet")
	s = strings.ReplaceAll(s, ".Bl \\-dash", ".Bl -dash")
	s = strings.ReplaceAll(s, ".Bl \\-tag", ".Bl -tag")

	return s
}
