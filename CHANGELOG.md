# Changelog
All notable changes to this project will be documented in this file.

Format is based on [Keep a Changelog] (https://keepachangelog.com/en/1.0.0/).
Versionning adheres to [Semantic Versioning] (https://semver.org/spec/v2.0.0.html).

## [unrelesaed]
## Modified
- gather all logic related to ANSI escape sequence management in a dedicated
  sub-module 
- improve wrap, truncate, justify... function to cleanly terminate interrupted
  ANSI sequences

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
