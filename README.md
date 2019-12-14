# PACKAGE STYLE
[![GoDoc](https://godoc.org/github.com/pirmd/text?status.svg)](https://godoc.org/github.com/pirmd/text)&nbsp; 
[![Go Report Card](https://goreportcard.com/badge/github.com/pirmd/text)](https://goreportcard.com/report/github.com/pirmd/text)&nbsp;

`text` provides text manipulation functions (like indentation, wrapping,
columnize,...) that can differentiate printable from non-printable sequences
(like ANSI colored sequences).

`text` provides helpers to format tables and to print diff between text.

## EXAMPLE

```go
package main

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

```

A slightly more complete example can be found in [testing file](style_test.go) or in 
[clapp](https://github.com/pirmd/clapp) package's [help](https://github.com/pirmd/clapp/blob/master/help.go)
or [manpage](https://github.com/pirmd/clapp/blob/master/manpage.go) files generation.

## INSTALLATION
Everything should work fine using go standard commands (`build`, `get`,
`install`...).

## USAGE
Running `go doc github.com/pirmd/text` should give you helpful guidelines on
available features.

## CONTRIBUTION
If you feel like to contribute, just follow github guidelines on
[forking](https://help.github.com/articles/fork-a-repo/) then [send a pull
request](https://help.github.com/articles/creating-a-pull-request/)

[modeline]: # ( vim: set fenc=utf-8 spell spl=en: )
