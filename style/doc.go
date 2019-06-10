// Package style provides a simple set of functions to decorate a text
// (formating, changing case or colors, ...) for different idioms.
//
// Provided styles can be extended and selected at runtime using CurrentStyler
// for ease of use.
//
// A set of examples are to be found in the test package or in cli/app/help.go
// that uses some of package style api to generate help information in
// different format (plain text, colored text, mandoc or markdown).
//
// Package style proposed also a simplifed way to specify styles based on a
// really simple markup syntax. This syntax can also be extended by linking a
// regexp to a style to apply.
//
// Package style approach is naive on-purpose to obtain an hopefully simple
// library.
package style
