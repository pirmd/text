package text_test

import (
	"fmt"

	"github.com/pirmd/text"
	"github.com/pirmd/text/ansi"
)

func Example() {
	tab := text.NewTable().SetMaxWidth(80).SetGrid(" ", "-", " ")

	tab.Rows(
		[]string{"Column1", "Column2", "Column3"},
		[]string{"Basic column", "This one is here\nto demonstrate\nthat colums with several lines work too", "Any " + ansi.SetBold("formatted") + " string can be inserted too without beaking the table."},
		[]string{"", "This second row is here to test multi-lines rows format", "Also possibly a second chance to verify that multi-lines is working"},
	)

	fmt.Println(tab.Draw())
}
