package table_test

import (
	"fmt"

	"github.com/pirmd/text/ansi"
	"github.com/pirmd/text/table"
)

func Example() {
	tab := table.New().SetMaxWidth(80).SetGrid(" ", "-", " ")

	tab.Rows(
		[]string{"Column1", "Column2", "Column3"},
		[]string{"Basic column", "This one is here\nto demonstrate\nthat colums with several lines work too", "Any " + ansi.Bold("formatted") + " string can be inserted too without beaking the table."},
		[]string{"", "This second row is here to test multi-lines rows format", "Also possibly a second chance to verify that multi-lines is working"},
	)

	fmt.Println(tab.Draw())
}
