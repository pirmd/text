// Package style provides a basic set of functions to decorate a text
// (formating, changing case or colors, ...) for different idioms.
//
// Provided styles can be combined and extended at runtime.
//
// A set of examples are to be found in the test package or in cli/app/help.go
// that uses some of package style api to generate help information in
// different format (plain text, colored text, mandoc or markdown).
//
// Package style proposed also a simplifed way to specify styles based on a
// minimal markup syntax. This syntax can also be extended by linking a
// regexp to a style to apply.
//
// Package design is limited raising some limitations when chaining styling
// functions, notably chaining styling's functions might end up in screwed-up
// output as it can stack multiple Escaping or Wraping functions with
// unexpected results.
package style
