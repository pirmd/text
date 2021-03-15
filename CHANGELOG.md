# Changelog

## [0.6.1] - 2021-0315
- Add LazyWrap and visual.LazyCut to handle long words (like URL) waiting for a
  better solution for splitting such words.

## [0.6.0] - 2021-02-05
- Major refactor and rewrite of code that breaks API but hopefully simplify
  code, add new features and solve some bugs.

## [0.5.4] - 2020-05-20
- FIX bug that panic when drawing an empty table.

## [0.5.3] - 2020-05-08
- FIX bug when wrapping ANSI-formatted text into table columns.

## [0.5.2] - 2020-05-08
- FIX bug in diff ByLines tokenizer that were adding uneeded '\n'.

## [0.5.1] - 2020-02-22
- FIX bug in Table where empty columns where not properly managed.
- Simplify Table API by deleting non used wrapper around NewTable.

## [0.5.0] - 2020-02-22
- Add a flag to select whether wrap - and functions based on wrap - should cut
  long words that excede line limits or not.

## [0.4.2] - 2020-02-22
- FIX a bug when wrapping line with "words" that are longer than the line limit
  (flush was not properly done)

## [0.4.1] - 2020-02-21
- remove dependency to github.com/primd/verify to avoid further cyclic import issue
- FIX bug when formatting table column that need interrupting ANSI formatting
  at the end of each lines of a table cell (ANSI ofrmatting interruption need
  to happen before padding a line with space to reach expected column width)

## [0.4.0] - 2020-02-16
- remove broken diff algorithms and replace it by a brand new diff sub-module
  that supports LCS and Patience diff.
- gather all logic related to ANSI escape sequence management in a dedicated
  sub-module 
- improve table, wrap, truncate, justify... and affiliated functions to cleanly
  terminate interrupted ANSI sequences

## [0.3.1] - 2019-11-11
- separate text from rest of github.com/pirmd/cli and
  github.com/pirmd/cli/style repositories

## [0.3.0] - 2019-11-10

## [0.2.0] - 2019-08-11
- add support for table's horizontal separators
- add a new DrawTable() function
- add a new TabWithBullet() function
- rename Table.Title() to Table.Captions(), which should hopefully be less
  misleading
- correct BUG where text.Indent() and text.Wrap() where not accepting non
  visible chars (like AINSI colors). It also modified text.Indent() and
  text.Tab() api: prefix is a string and not a []byte
- allow command exxecution if no args have been specified by user

## [0.1.0] - 2019-05-11
- wraping, justifying, table, text diff formatting
