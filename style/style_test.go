package style

import (
	"github.com/pirmd/verify"
	"testing"
)

func testText() (s string) {
	S := Chainf(Upper, Blue, Header)
	P := Chainf(Paragraph, Tab)
	L := Chainf(List, Tab)

	s = TrimLeadingSpace(S("Introduction"))
	s += P("This small piece of text aims at demonstrating and testing my styling package '%s'.", Underline("style"))
	s += P("It is writen by a %s English speaker, so pardon any faults", Bold("non-native"))
	s += S("Examples of available styles")
	s += P("Section %s already demonstrates useful styles from package 'styles', this section completes with most of the others possibilities.", Underline("Introduction"))
	s += P("Notably, package '%s' can print in %s or %s.", Underline("style"), Red("red"), Bold(Green("bold green")))
	s += P("It also knows how to format %s:", Italic("lists"))
	s += L("This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped. It should hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached.")
	s += L("It is also possible to check that paragraph inside lists are respected.\nAs you can see here in this simple example.")
	s += P("It also knows how to %s terms:", Italic("define"))
	s += Tab(DefTerm("style"))
	s += Tab2(DefDesc("A particular procedure by which something is done; a manner or way."))
	s += S("Demonstrating tables")
	s += P("Package 'style' supports drawing tables for most basic cases. Using Tab to align tables should be done carefully as table way to guess optimum columns size does not take into account the tabulation")
	s += Table(
		[]string{"Column1", "Column2", "Column3"},
		[]string{"Basic column", "This one is here\nto demonstrate\nthat several lines\ncolumn work", "Last but not least shows " + Bold("formating") + " within the table"},
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

func TestStyleMandoc(t *testing.T) {
	CurrentStyler = Mandoc
	out := testText()
	verify.MatchGolden(t, out, "Styling with 'Mandoc' Markup failed")
}

func TestStyleMarkdown(t *testing.T) {
	CurrentStyler = Markdown
	DefaultTxtWidth = 60
	out := testText()
	verify.MatchGolden(t, out, "Styling with 'Markdown' Markup failed")
}
