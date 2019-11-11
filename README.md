# PACKAGE STYLE
[![GoDoc](https://godoc.org/github.com/pirmd/text?status.svg)](https://godoc.org/github.com/pirmd/text)&nbsp; 
[![Go Report Card](https://goreportcard.com/badge/github.com/pirmd/text)](https://goreportcard.com/report/github.com/pirmd/text)&nbsp;

`text` provides text manipulation functions (like indentation, wrapping,
columnize,...) that can differentiate printable from non-printable sequences
(like ANSI colored sequences). Tables formatting helpers are proposed as well
as text diff formatting.

## EXAMPLE

```go
package main

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
