package table_test

import (
	"fmt"

	"github.com/pirmd/text/ansi"
	"github.com/pirmd/text/table"
)

func Example() {
	tab := table.New().SetMaxWidth(80).SetGrid(&table.Grid{Columns: " | ", Header: "=", BodyRows: "-", Footer: "="})

	tab.SetHeader("Column1", "Column2", "Column3")
	tab.AddRows(
		[]string{"Basic column", "This one is here\nto demonstrate\nthat columns with several lines work too.", "Any " + ansi.Bold("formatted") + " string can be inserted too without breaking the table."},
		[]string{"", "This second row is here to test multi-lines rows format.", "Also possibly a second chance to verify that multi-lines is working."},
	)
	tab.SetFooter("", "Grand Total:", ansi.Bold(ansi.Green("42")))

	fmt.Println(tab)
}

//Output:
//Column1      | Column2                         | Column3
//============ | =============================== | ===============================
//Basic column | This one is here                | Any [1mformatted[22m string can be
//             | to demonstrate                  | inserted too without breaking
//             | that columns with several lines | the table.
//             | work too.                       |
//------------ | ------------------------------- | -------------------------------
//             | This second row is here to test | Also possibly a second chance
//             | multi-lines rows format.        | to verify that multi-lines is
//             |                                 | working.
//============ | =============================== | ===============================
//             | Grand Total:                    | [1m[32m42[39m[22m
