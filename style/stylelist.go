//go:generate go run helpers_generate.go

package style

//Format group all available styles to format text
type Format int

const (
	//Changes a string case to upper case
	FmtUpper Format = iota
	//Changes a string case to lower case
	FmtLower
	//Changes string foreground color to black
	FmtBlack
	//Changes string foreground color to red
	FmtRed
	//Changes string foreground color to green
	FmtGreen
	//Changes string foreground color to yellow
	FmtYellow
	//Changes string foreground color to blue
	FmtBlue
	//Changes string foreground color to magenta
	FmtMagenta
	//Changes string foreground color to cyan
	FmtCyan
	//Changes string foreground color to white
	FmtWhite
	//Changes a string case to bold
	FmtBold
	//Changes a string case to italic
	FmtItalic
	//Changes a string to be underlined
	FmtUnderline
	//Changes a string by inverting its fore- and back-ground colors
	FmtInverse
	//Changes a string to be strikethrough
	FmtStrike
	//Wraps a text
	FmtWrap
	//Indents then wraps a text
	FmtTab
	//Indents two times then wraps a text
	FmtTab2
	//Displays information in the document headers
	FmtDocHeader
	//Displays a section header
	FmtHeader
	//Displays a paragraph
	FmtParagraph
	//Displays a new line
	FmtNewLine
	//Displays a simple list
	FmtList
	//Displays the term part of a defintion list
	FmtDefTerm
	//Displays the description part of a definition list
	FmtDefDesc
	//Displays a code block
	FmtCode
	//Escape the input text. FmtEscape is automatically
	//applied and does not usually needed to be called manually
	FmtEscape
	//Auto-style the input text. Usually used with a Markup
	//FmtAuto is automatically applied and does not usually needed
	//to be called manually
	FmtAuto
)
