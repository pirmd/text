//`cli` is a collection of functions that I often use when developping simple cli apps.
//
//It is separate in sub-packages:
// . app:    contains functions to build a cli application, featuring
//           several levels of sub-commands, retrieving of flags and args,
//           help and manpage buiding.
//
// . style:  contains functions to decorate a text using several idioms
//           (plaintext, (non-)colored term,  mandoc, markdown). It can be
//           extended
//
// . style/text: contains text manipulation functions (like identation, wraping,
//               columnize,...) that can differentiate printable from non-printable
//               sequences (like ANSI colored sequences).
//               table formatting helpers are proposed as well as text diff formatting
//
// . formatter: contains functions to quickly build a string out of an
//              object. It can be useful to pretty print a familly of
//              objects, adopting a given formatting scheme based on
//              their type.
//
// . input:  contains functions to input or edit text from cli

package cli
