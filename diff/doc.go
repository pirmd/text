// Package diff proposes a set of functions to compute differences between two
// strings. It offers standard implementation of (lcs based diff)[from
// https://en.m.wikipedia.org/wiki/Longest_common_subsequence_problem] as well as
// the (patience diff)[http://alfedenzo.livejournal.com/170301.html].
//
// On top of these algorythms, some home-brewed approach is supposed to provide
// more readable diff outputs. Main intend is to allow refining diff outputs
// by deeper looking into differences, for instance between almost similar
// lines, look at differences by words the, runes.
//
// `diff` package supports personalizing highlighters to pretty-print diff
// results.
package diff
