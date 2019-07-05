# Changelog
All notable changes to this project will be documented in this file.

Format is based on [Keep a Changelog] (https://keepachangelog.com/en/1.0.0/).
Versionning adheres to [Semantic Versioning] (https://semver.org/spec/v2.0.0.html)

## [Unreleased]
### Added
- cli/app: add function to generate a help file in markdown format for a
  command 
- cli/app: add default support to print a version information taken from 'git
  describe and rev-parse' and set-up using 'ldflags -X' directive. Provides a
  simple shell script to facilitate the build/install directive for that purpose.
  This behaviour can be overwriten as wished
- cli/input: add function to fire up a tool for a user to manually merge
  files/data
- cli/style: add new styling functions: NoLeadingSpace, NoTrailingSpace and
  TrimSpace
- cli/style: add support to format tables, nested lists and tabs
- cli/style: add preliminary support to mdoc format
- cli/style/text: add support for table's horizontal separators
- cli/style/text: add a new DrawTable() function
- cli/style/text: add a new TabWithBullet() function
### Modified
- cli/style: rename NewLine to Line
- cli/style: give markdown some love, alig nas far as possible plain text
  format with markdown
- cli/style: clean api and (hopefully) propose better names
  (style.Func -> style.FormatFn, style.New -> style.Chainf)
- cli/style/text: rename Table.Title() to Table.Captions(), which
  should hopefully be less misleading
- cli/style/text: correct BUG where text.Indent() and text.Wrap() where not
  accepting non visible chars (like AINSI colors).
  It also modified text.Indent() and text.Tab() api: prefix is a string and not
  a []byte


## [0.1.0] - 2019-05-11
### Added
- commandline app definition with flags and args parsing and help and/or
  manpage generation
- manipulation of often used text styling for (color)term, manpage idiom
  and (very)light markdown
- wraping, justifying, table, text diff formatting
