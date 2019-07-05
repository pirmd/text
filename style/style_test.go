package style

import (
	"testing"

	"github.com/pirmd/verify"
)

func testText() (s string) {
	H := Chainf(Upper, Blue, Header)
	H2 := Chainf(Title, Red, Header2)
	P := Chainf(Paragraph, Wrap)

	s = NoLeadingSpace(H("Introduction"))
	s += P("This small piece of text aims at demonstrating and testing my styling package '%s'.", Underline("style"))
	s += P("It is writen by a %s English speaker, so pardon any faults.", Bold("non-native"))
	s += H("Examples of available styles")

	s += H2("Demonstrating common formatting")
	s += P("Section %s already demonstrates useful styles from package 'styles', this section completes them with most of the others possibilities.", Underline("Introduction"))
	s += P("Notably, package '%s' can print in %s or %s (if chosen style supports it).", Underline("style"), Red("red"), Bold(Green("bold green")))
	s += P("Several levels of tabulations can be used:")
	s += Tab(1)(Paragraph("(Level 1) Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."))
	s += Tab(2)(Paragraph("(Level 2) Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."))
	s += Tab(4)(Paragraph("(Level 4) Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."))

	s += H2("Demonstrating lists")
	s += P("It also knows how to format %s:", Italic("lists"))
	s += List(0)(
		ListItem(0)("This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped. It should hopefully be so long that it will not fulfill the maximum number of authorized chars per line is reached."),
		ListItem(0)("It also can support sub-lists:")+List(1)(
			ListItem(1)("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
			ListItem(1)("Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."),
			ListItem(1)("Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
		),
		ListItem(0)(
			Line("It is also possible to have a list's item that contains several paragraphs")+
				Paragraph("For example, this paragraph that I made artificially long to verify that wrapping is working correctly inside list"),
		),
	)
	s += P("It also knows how to %s terms:", Italic("define"))
	s += Define("style", "A particular procedure by which something is done; a manner or way.")
	s += Define("lorem ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")

	s += H2("Demonstrating tables")
	s += P("Package 'style' supports drawing tables for most basic cases. Using Tab to align tables should be done carefully as table way to guess optimum columns size does not take into account the tabulation")
	s += Table(
		[]string{"Column1", "Column2", "Column3"},
		[]string{"Basic column", "This one is here\nto demonstrate\nthat colums with several lines work too", "Last but not least shows " + Bold("formating") + " within the table"},
		[]string{"", "This second row is here to test multi-lines rows format", "Also possibly a good opportunity to check escaping *sequence*"},
	)

	return
}

func TestStylePlainText(t *testing.T) {
	CurrentStyler = PlainText
	DefaultTxtWidth = 60
	out := testText()
	verify.MatchGolden(t, out, "Styling with 'PlainText' Markup failed")
}

func TestStyleTerm(t *testing.T) {
	CurrentStyler = Term
	termWidth = 60
	out := testText()
	verify.MatchGolden(t, out, "Styling with 'Term' Markup failed")
}

func TestStyleColorTerm(t *testing.T) {
	CurrentStyler = ColorTerm
	termWidth = 60
	out := testText()
	verify.MatchGolden(t, out, "Styling with 'ColorTerm' Markup failed")
}

func TestStyleMan(t *testing.T) {
	CurrentStyler = Man
	out := testText()
	verify.MatchGolden(t, out, "Styling with 'Man' Markup failed")
}

func TestStyleMdoc(t *testing.T) {
	CurrentStyler = Mdoc
	out := testText()
	verify.MatchGolden(t, out, "Styling with 'Mdoc' Markup failed")
}

func TestStyleMarkdown(t *testing.T) {
	CurrentStyler = Markdown
	DefaultTxtWidth = 60
	out := testText()
	verify.MatchGolden(t, out, "Styling with 'Markdown' Markup failed")
}
