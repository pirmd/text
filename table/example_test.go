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
		[]string{"Basic column", "Multi-line row:\n- first line\nSecond line is working too.", "Any very and interesting long line is also going to be adequatly wrapped at column boundaries."},
		[]string{"", "<- Empty columns are properly managed (as you can see on your left and right) ->", ""},
	)
	tab.SetFooter("but whatever you'll put in your table,", "the answer will be:", "42")

	fmt.Println(tab)
	// Output:
	// Column1                  | Column2                  | Column3
	// ======================== | ======================== | ========================
	// Basic column             | Multi-line row:          | Any very and interesting
	//                          | - first line             | long line is also going
	//                          | Second line is working   | to be adequatly wrapped
	//                          | too.                     | at column boundaries.
	// ------------------------ | ------------------------ | ------------------------
	//                          | <- Empty columns are     |
	//                          | properly managed (as you |
	//                          | can see on your left and |
	//                          | right) ->                |
	// ======================== | ======================== | ========================
	// but whatever you'll put  | the answer will be:      | 42
	// in your table,           |                          |

	// I'm not sure why but it seems this comment is needed or the test will fail.
}

func Example_with_colors() {
	tab := table.New().SetMaxWidth(80).SetGrid(&table.Grid{Columns: " | ", Header: ansi.Bold("-"), BodyRows: "-", Footer: ansi.Bold("-")})

	tab.SetHeader("Let's put", "Some fun", "With colors")
	tab.AddRows(
		[]string{"Basic column", ansi.Green("Multi-line row:\n- first line\nSecond line is working too."), "Any very and " + ansi.Underline("interesting") + " long line is also going to be adequatly wrapped at column boundaries."},
		[]string{"", ansi.Blue("<-") + " Empty columns are properly managed (as you can see on your left and right)" + ansi.Red("->"), ""},
	)
	tab.SetFooter("but whatever you'll put in your table,", "the answer will be:", ansi.Bold(ansi.Green("42")))

	fmt.Println(tab)
	// Output:
	// Let's put                | Some fun                 | With colors
	// [1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[0m | [1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[0m | [1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[0m
	// Basic column             | [32mMulti-line row:[0m          | Any very and [4minteresting[24m
	//                          | [32m- first line[0m             | long line is also going
	//                          | [32mSecond line is working [0m  | to be adequatly wrapped
	//                          | [32mtoo.[39m                     | at column boundaries.
	// ------------------------ | ------------------------ | ------------------------
	//                          | [34m<-[39m Empty columns are     |
	//                          | properly managed (as you |
	//                          | can see on your left and |
	//                          | right)[31m->[39m                 |
	// [1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[0m | [1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[0m | [1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[22m[1m-[0m
	// but whatever you'll put  | the answer will be:      | [1m[32m42[39m[22m
	// in your table,           |                          |

	// I'm not sure why but it seems this comment is needed or the test will fail.
}
