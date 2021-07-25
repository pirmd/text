package ansi

import (
	"bytes"
	"errors"
	"unicode"
	"unicode/utf8"
)

var (
	// ErrStopWalk is used as a return value from WalkFunc to indicate to Walk
	// to interrupt its iteration.
	ErrStopWalk = errors.New("stop walk")

	// ErrIncompleteRune is the error raised if utf-8 decoding error occurs
	// during write operation.
	ErrIncompleteRune = errors.New("utf-8 decoding errors")

	// ErrIncompleteAnsiEscSeq is the error raised if an incomplete ANSI
	// sequence is encountered.
	ErrIncompleteAnsiEscSeq = errors.New("incomplete ANSI escape sequence")
)

// WalkFunc is the type of the function that is called while going through a
// string that might contains ANSI escape codes.
// Walking occurs either rune per rune if the rune does not belong to an ANSI
// escape sequence or ANSI escape sequence per ANSI escape sequence in the
// other case. If WalkFunc is called for an ANSI Escape Sequence, the curRune
// will be -1.
// If WalkFunc returns the ErrStopWalk error, Walk will interrupt the walk.
type WalkFunc func(advance int, curRune rune, curSGREsc string) error

// Walk walks through a slice of bytes that can contain ANSI escape codes and
// run WalkFunc either on each rune that is not part of an escape sequence or
// on each ANSI escape sequence.
func Walk(p []byte, fn WalkFunc) (err error) {
	advance := 0
	inEscSeq := false
	curEsc := new(bytes.Buffer)

	for len(p[advance:]) > 0 {
		c, sz := utf8.DecodeRune(p[advance:])
		if c == utf8.RuneError {
			err = ErrIncompleteAnsiEscSeq
			return
		}
		advance += sz

		switch {
		case c == '\x1b':
			inEscSeq = true
			curEsc.WriteRune(c)

		case inEscSeq:
			curEsc.WriteRune(c)
			if unicode.IsLetter(c) || c == '~' {
				err = fn(advance, -1, curEsc.String())
				inEscSeq = false
				curEsc.Reset()
			}

		default:
			err = fn(advance, c, curEsc.String())
		}

		if err != nil {
			if err == ErrStopWalk {
				err = nil
			}

			return
		}
	}

	if curEsc.Len() > 0 {
		err = ErrIncompleteAnsiEscSeq
	}

	return
}

// WalkString walks through a string that can contain ANSI escape codes and run
// WalkFunc either on each rune that is not part of an escape sequence or on
// each ANSI escape sequence.
func WalkString(s string, fn WalkFunc) (err error) {
	return Walk([]byte(s), fn)
}
