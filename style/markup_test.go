package style

import (
	"fmt"
	"testing"
	"github.com/pirmd/verify"
)

func testTextSyntax() (s string) {
	Section := New(Upper, Blue, Header)

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

	return s
}

func TestSyntaxToColorTerm(t *testing.T) {
	termWidth = 60
	CurrentStyler = ColorTerm.WithAutostyler(LightMarkup)
	got := testText()
	verify.MatchGolden(t, got, "Syntax conversion to 'ColorTerm' Markup failed")
}

func TestSyntaxLightMarkup(t *testing.T) {
	testStyler := Styler{
		FmtBold:   Sprintf("BOLD{%s}"),
		FmtItalic: Sprintf("ITALIC{%s}"),
		FmtCode:   Sprintf("CODE{%s}"),
	}

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
