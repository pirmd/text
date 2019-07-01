package style

//Mandoc is a sub-set of troff markup featuring common used macro for building
//man pages
var Mandoc = core.Extend(New(
	FormatMap{
		FmtBold:   Sprintf("\\fB%s\\fP"),
		FmtItalic: Sprintf("\\fI%s\\fP"),

		FmtDocHeader: Sprintf(".TH %s\n"),
		FmtHeader:    Sprintf("\n.SH %s\n"),
		FmtParagraph: Sprintf(".PP\n%s\n"),
		FmtLine:      Sprintf(".br\n%s\n"),
		FmtList:      Sprintf(".RS\n%s\n.RE\n"),
		FmtDefTerm:   Sprintf("\n.TP\n%s\n"),
		FmtDefDesc:   Sprintf("%s\n"),
		FmtCode:      Sprintf(".PP\n.RS\n.nf\n%s\n.fi\n.RE\n"),

		FmtEscape: escapeMandoc,
	},
	nil,
))

//XXX: add mandoc tbl extension

func escapeMandoc(s string) string {
	var b []byte
	var isEscaped bool

	for _, c := range s {
		switch {
		case c == '\\':
			if isEscaped {
				b = append(b, '\\', '\\')
				isEscaped = false
			}
			isEscaped = true

		case c == '-' || c == '_' || c == '&' || c == '~':
			b = append(b, '\\', byte(c))
			isEscaped = false

		default:
			if isEscaped { //isEscaped is triggerd but escape nothing known, it is eventually a standalone '\' that needs to b eescaped itself
				b = append(b, '\\')
			}
			b = append(b, byte(c))
			isEscaped = false
		}
	}
	return string(b)
}
