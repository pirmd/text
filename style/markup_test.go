package style

import (
	"fmt"
	"github.com/pirmd/verify"
	"testing"
)

func testTextSyntax() (s string) {
	Section := Chainf(Upper, Blue, Header)

	s = Section("Introduction")
	s += Paragraph("This small piece of text aims at demonstrating and testing my styling package '*style*'.")
	s += Paragraph("It is writen by a *non-native* English speaker, so pardon any faults")
	s += Section("Available styles")
	s += Paragraph("Section _Introduction_ already demonstrates useful '*styles*''s Styles, this section completes with most of the others possibilities.")
	s += Paragraph("Notably, package '*style*' can print in %s or *%s*.", Red("red"), Green("bold green"))
	s += Paragraph("It also knows how to format _lists_:")
	s += List("this very long and detailed sentence is here to demonstrate that list can be formatted and wrapped because it hopefully be so long that it will not fulfill the maximum number of authorized char per lines is reached.")
	s += List("It is also possible to check that paragraph inside lists are respected.\nAs you can see here in this simple example.")
	s += Paragraph("It also knows how to _define_ terms:")
	s += DefTerm("style") + DefDesc("A particular procedure by which something is done; a manner or way.")
	s += Section("Demonstrating tables")
	s += Paragraph("Package '*style*' supports drawing tables for most basic cases. Using Tab to align tables should be done carefully as table way to guess optimum columns size does not take into account the tabulation")
	s += Table(
		[]string{"Column1", "Column2", "Column3"},
		[]string{"Basic column", "This one is here\nto demonstrate\nthat several lines\ncolumn work", "Last but not least, shows **formating** within the table"},
	)

	return s
}

func TestSyntaxToColorTerm(t *testing.T) {
	termWidth = 60
	CurrentStyler = ColorTerm.WithAutostyler(LightMarkup)
	got := testText()
	verify.MatchGolden(t, got, "Syntax conversion to 'ColorTerm' Markup failed")
}

func TestSyntaxLightMarkup(t *testing.T) {
	testStyler := New(FormatMap{
		FmtBold:   Sprintf("BOLD{%s}"),
		FmtItalic: Sprintf("ITALIC{%s}"),
		FmtCode:   Sprintf("CODE{%s}"),
	}, nil)

	testCases := []struct {
		in  string
		out string
	}{
		{"*Test*", "BOLD{Test}"},
		{"_Test_", "ITALIC{Test}"},
		{"_*Test*_", "ITALIC{BOLD{Test}}"},
		{"*Test* is a difficult *art*.", "BOLD{Test} is a difficult BOLD{art}."},
		{"*Test* is a difficult _art_.", "BOLD{Test} is a difficult ITALIC{art}."},
		{"Test **is** a difficult _art_.", "Test **is** a difficult ITALIC{art}."},
		{"*Don't* try this at home `rm -rf /`", "BOLD{Don't} try this at home CODE{rm -rf /}"},
	}

	for _, tc := range testCases {
		got := LightMarkup.render(testStyler, tc.in)
		verify.EqualString(t, got, tc.out, fmt.Sprintf("Fail to render markup for '%s'", tc.in))
	}
}
