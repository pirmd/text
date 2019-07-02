//go:generate go run helpers_generate.go

package style

//Format group all available styles to format text
type Format int

const (
	//FmtUpper changes a string case to upper case
	FmtUpper Format = iota
	//FmtLower changes a string case to lower case
	FmtLower
	//FmtTitle changes all letters that begin words to their title case.
	FmtTitle
	//FmtTrimSpace removes any leading or trailing space
	FmtTrimSpace
	//FmtTrimLeadingSpace removes any leading space
	FmtTrimLeadingSpace
	//FmtBlack decorates a string foreground color to black
	FmtBlack
	//FmtRed decorates a string foreground color to red
	FmtRed
	//FmtGreen decorates a string foreground color to green
	FmtGreen
	//FmtYellow decorates a string foreground color to yellow
	FmtYellow
	//FmtBlue decorates a string foreground color to blue
	FmtBlue
	//FmtMagenta decorates a string foreground color to magenta
	FmtMagenta
	//FmtCyan decorates a string foreground color to cyan
	FmtCyan
	//FmtWhite decorates a string foreground color to white
	FmtWhite
	//FmtBold decorates a string case to bold
	FmtBold
	//FmtItalic decorates a string case to italic
	FmtItalic
	//FmtUnderline decorates a string to be underlined
	FmtUnderline
	//FmtInverse decorates a string by inverting its fore- and back-ground
	//colors
	FmtInverse
	//FmtStrike decorates a string to be strikethrough
	FmtStrike
	//FmtWrap wraps a text
	FmtWrap
	//FmtTab indents then wraps a text
	FmtTab
	//FmtTab2 indents two times then wraps a text
	FmtTab2
	//FmtDocHeader displays text as information in the document headers
	FmtDocHeader
	//FmtHeader displays text as a section header
	FmtHeader
	//FmtHeader2 displays text as a section header of rank 2
	FmtHeader2
	//FmtHeader3 displays text as a section header of rank 3
	FmtHeader3
	//FmtParagraph displays text as a new paragraph
	FmtParagraph
	//FmtLine displays text as a new line
	FmtLine
	//FmtList displays text as a new simple list member
	FmtList
	//FmtList2 displays text as a new sub-list member
	FmtList2
	//FmtDefTerm displays text as the term part of a definition list
	FmtDefTerm
	//FmtDefDesc displays text as the description part of a definition list
	FmtDefDesc
	//FmtCode displays text as a code block
	FmtCode
	//FmtEscape escapes the input text. FmtEscape is automatically applied and
	//does not usually needed to be called manually
	FmtEscape
	//FmtAuto auto-styles the input text. Usually used with a Markup FmtAuto is
	//automatically applied and does not usually needed to be called manually
	FmtAuto
)
