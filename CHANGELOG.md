# Changelog
All notable changes to this project will be documented in this file.

Format is based on [Keep a Changelog] (https://keepachangelog.com/en/1.0.0/).
Versionning adheres to [Semantic Versioning] (https://semver.org/spec/v2.0.0.html).

## [0.5.1] - 2020-02-22
## Modified
- FIX bug in Table where empty columns where not properly managed.
## Removed
- Simplify Table API by deleting non used wrapper around NewTable.

## [0.5.0] - 2020-02-22
## Modified
- Add a flag to select whether wrap - and functions based on wrap - should cut
  long words that excede line limits or not.

## [0.4.2] - 2020-02-22
## Modified
- FIX a bug when wrapping line with "words" that are longer than the line limit
  (flush was not properly done)

## [0.4.1] - 2020-02-21
## Modified
- remove dependency to github.com/primd/verify to avoid further cyclic import issue
- FIX bug when formatting table column that need interrupting ANSI formatting
  at the end of each lines of a table cell (ANSI ofrmatting interruption need
  to happen before padding a line with space to reach expected column width)

## [0.4.0] - 2020-02-16
## Added
- remove broken diff algorithms and replace it by a brand new diff sub-module
  that supports LCS and Patience diff.
## Modified
- gather all logic related to ANSI escape sequence management in a dedicated
  sub-module 
- improve table, wrap, truncate, justify... and affiliated functions to cleanly
  terminate interrupted ANSI sequences

## [0.3.1] - 2019-11-11
## Removed
- separate text from rest of github.com/pirmd/cli and
  github.com/pirmd/cli/style repositories

## [0.3.0] - 2019-11-10

## [0.2.0] - 2019-08-11
### Added
- add support for table's horizontal separators
- add a new DrawTable() function
- add a new TabWithBullet() function
### Modified
- rename Table.Title() to Table.Captions(), which should hopefully be less
  misleading
- correct BUG where text.Indent() and text.Wrap() where not accepting non
  visible chars (like AINSI colors). It also modified text.Indent() and
  text.Tab() api: prefix is a string and not a []byte
- allow command exxecution if no args have been specified by user

## [0.1.0] - 2019-05-11
### Added
- wraping, justifying, table, text diff formatting
