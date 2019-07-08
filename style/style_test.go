package style_test

import (
	"github.com/pirmd/verify"
	"testing"

	"github.com/pirmd/cli/style"
)

func testText(st style.Styler) (s string) {
	//XXX:H := Chainf(Upper, Blue, Header)
	//XXX:H2 := Chainf(Title, Red, Header2)
	//XXX:P := Chainf(Paragraph, Wrap)

	s = st.Header(1)("Introduction")
	s += st.Paragraph("This small piece of text aims at demonstrating and testing package '" + st.Underline("style") + "'.")
	s += st.Paragraph("It is writen by a " + st.Bold("non-native") + " English speaker, so pardon any faults.")

	s += st.Header(1)("Examples of available styles")

	s += st.Header(2)("Demonstrating common formatting")
	s += st.Paragraph("Section " + st.Underline("Introduction") + " already demonstrates useful styles from package 'styles', this section completes them with most of the others possibilities.")
	s += st.Paragraph("Notably, package " + st.Underline("style") + " can print in " + st.Red("red") + " or " + st.Bold(st.Green("green bold")) + " (if chosen style supports it).")
	s += st.Paragraph("Several levels of tabulations can be used:")
	s += st.Tab(1)(st.Paragraph("(Level 1) Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."))
	s += st.Tab(2)(st.Paragraph("(Level 2) Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."))
	s += st.Tab(4)(st.Paragraph("(Level 4) Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."))

	s += st.Header(2)("Demonstrating lists")
	s += st.Paragraph("It also knows how to format " + st.Italic("lists") + ": ")
	s += st.List(0)(
		"This very long and detailed sentence is here to demonstrate that list can be formatted and wrapped. It should hopefully be so long that it will not fulfill the maximum number of authorized chars per line is reached.",
		"It also can support sub-lists:\n"+st.List(1)(
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
			"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		),
		"It is also possible to have a list's item that contains several paragraphs\n"+
			st.Paragraph("For example, this paragraph that I made artificially long to verify that wrapping is working correctly inside list"),
	)
	s += st.Paragraph("It also knows how to " + st.Italic("define") + " terms:")
	s += st.Define("style", "A particular procedure by which something is done; a manner or way.")
	s += st.Paragraph("Even using Tabs:")
	s += st.Tab(1)(st.Define("lorem ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."))

	s += st.Header(2)("Demonstrating tables")
	s += st.Paragraph("Package 'style' supports drawing tables for most basic cases. Using Tab to align tables should be done carefully as table way to guess optimum columns size does not take into account the tabulation")
	s += st.Table(
		[]string{"Column1", "Column2", "Column3"},
		[]string{"Basic column", "This one is here\nto demonstrate\nthat colums with several lines work too", "Last but not least shows " + st.Bold("formating") + " within the table"},
		[]string{"", "This second row is here to test multi-lines rows format", "Also possibly a good opportunity to check escaping *sequence*"},
	)
	s += st.Paragraph("It is also possible to use tabs with tables:")
	s += st.Tab(1)(
		st.Table(
			[]string{"Column1", "Column2", "Column3"},
			[]string{"Basic column", "This one is here\nto demonstrate\nthat colums with several lines work too", "Last but not least shows " + st.Bold("formating") + " within the table"},
			[]string{"", "This second row is here to test multi-lines rows format", "Also possibly a good opportunity to check escaping *sequence*"},
		),
	)

	return
}

func TestStyleCore(t *testing.T) {
	out := testText(&style.Core{ListBullets: []string{"-", "*", "+"}})
	verify.MatchGolden(t, out, "Styling with 'Core' style failed")
}

func TestStylePlainText(t *testing.T) {
	out := testText(style.Plaintext)
	verify.MatchGolden(t, out, "Styling with 'Plaintext' style failed")
}
