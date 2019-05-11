# INTRODUCTION
`cli` is a collection of functions that I often use when developping simple cli
apps.

It is separate in sub-packages:
 - `app`: contains functions to build a cli application, featuring several
   levels of sub-commands, retrieving of flags and args, help and manpage
   buiding.

 - `style`: contains functions to decorate a text using several idioms
   (plaintext, (non-)colored term,  mandoc, markdown). It can be extended

 - `style/text`: contains text manipulation functions (like identation,
   wraping, columnize,...) that can differenciate printable from non-printable
   sequences (like ANSI colored sequences).  Tables formatting helpers are
   proposed as well as text diff formatting

 - `formatter`: contains functions to quickly build a string out of an object.
   It can be useful to pretty print a familly of objects, adopting a given
   formatting scheme based on their type.

 - `input`: contains functions to input or edit text from cli

# INSTALLATION
Everything should work fine using go standard commands (`build`, `get`,
`install`...).

# USAGE
Running `godoc` should give you helpful guidelines on availbales features.

# CONTRIBUTION
If you feel like to contribute, just follow github guidelines on
[forking](https://help.github.com/articles/fork-a-repo/) then [send a pull
request](https://help.github.com/articles/creating-a-pull-request/)

[modeline]: # ( vim: set fenc=utf-8 spell spl=en: )
