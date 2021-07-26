package visual

import (
	"bytes"
	"unicode"

	"github.com/pirmd/text/ansi"
)

var (
	// defaultWordBoundaries contains runes than are found at words boundaries,
	// without being necessarily followed by a space.
	defaultWordBoundaries = []byte{')', ']', '-', '|'}
)

// Cutter cuts a slice of bytes at end-of-line if the line's "visual"
// length is shorter than the maximum line limit or at word boundary to
// stay as close as possible under the maximum limit.
// Should a word exist that is longer than the limit, the word is either
// split in chunk or kept as is depending on Lazycut flag.
type Cutter struct {
	// Maxwidth is the maximum line's "visual" length at which the Cutter
	// force a split.
	Maxwidth int
	// Lazycut flag governs the Cutter's behavior in case a word is
	// encountered whose "visual" length is longer than Cutter's MaxWidth.
	// When true, the word is kept as is, if false, the word is cut in
	// place.
	Lazycut bool
	// WordBoundaries list the byte considered as word boundaries where a
	// Cutter can split a line.
	WordBoundaries []byte
}

// NewCutter creates a new Cutter.
func NewCutter(maxwidth int) *Cutter {
	return &Cutter{
		Maxwidth:       maxwidth,
		WordBoundaries: defaultWordBoundaries,
	}
}

// NewLazyCutter creates a new Cutter that lazily manage word longer than
// maxwidth.
func NewLazyCutter(maxwidth int) *Cutter {
	return &Cutter{
		Maxwidth:       maxwidth,
		Lazycut:        true,
		WordBoundaries: defaultWordBoundaries,
	}
}

// Split cuts s at the first encountered end-of-line or at the word boundary at which
// the "visual" width is as close as possible (but lower than) Cutter max width.
// Split returns the two parts of the operation's result. Should no end-of-line
// nor the "visual" width be reached, result will be nil, s.
func (cutr Cutter) Split(s []byte) ([]byte, []byte) {
	if ieol := cutr.nextEOL(s); ieol >= 0 {
		return s[:ieol], s[ieol:]
	}

	return nil, s
}

// SplitFunc implements bufio.SplitFunc for using Cutter with a bufio.Scanner.
func (cutr Cutter) SplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}

	if ieol := cutr.nextEOL(data); ieol >= 0 {
		return ieol, data[:ieol], nil
	}

	return 0, nil, nil
}

// nextEOL moves the Cutter's position after the next end-of-line or after the next word
// boundary as close as possible to Cutter's maximum width.
// nextEOL returns -1 if it reaches end of data without finding '\n' or without
// reaching maximum line's visual length.
// if maximum line's visual length is negative or null, nextEOL only returns
// next '\n' position.
func (cutr Cutter) nextEOL(data []byte) int {
	if cutr.Maxwidth <= 0 {
		if n := bytes.IndexRune(data, '\n'); n != -1 {
			return n + 1
		}
		return -1
	}

	ieol, iword := -1, -1
	i, w := 0, 0

	_ = ansi.Walk(data, func(n int, c rune, esc string) error {
		runewidth := Runewidth(c)

		switch {

		// we have found an ANSI escape sequence, we move forwards.
		// if we are at a word boundary of at the end of line, we also push
		// word/line to encompass the ANSI escape sequence.
		case len(esc) > 0:
			if ieol == i {
				ieol = n
			}
			if iword == i {
				iword = n
			}
			i = n

		// we have found the end-of-line, but we have continued to eventually
		// captured remaining consecutive ANSI escape sequences. No more escape
		// sequence are found, we break.
		case ieol >= 0:
			return ansi.ErrStopWalk

		case c == '\n':
			ieol, i = n, n
			return nil // end-of line is found, continue in case there is a remaining consecutive escape sequence.

		case unicode.IsSpace(c):
			if w+runewidth > cutr.Maxwidth {
				ieol = i
				return ansi.ErrStopWalk
			}

			iword, i = n, n
			w = w + runewidth

		case bytes.ContainsRune(cutr.WordBoundaries, c):
			if w+runewidth > cutr.Maxwidth {
				ieol = iword
				return ansi.ErrStopWalk
			}

			iword, i = n, n
			w = w + runewidth

		default:
			if w+runewidth > cutr.Maxwidth {
				switch {
				// current word will be longer than maximum width; we are lazy,
				// so we do nothing to prevent that.
				case iword == -1 && cutr.Lazycut:

				// current word will be longer than maximum width; we are not lazy,
				// so we cut it in pieces.
				case iword == -1:
					ieol = i
					return nil // end-of line is found, continue in case there is a remaining consecutive escape sequence.

				default:
					ieol = iword
					return ansi.ErrStopWalk
				}
			}

			i = n
			w = w + runewidth
		}

		return nil
	})

	return ieol
}
