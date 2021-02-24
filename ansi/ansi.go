package ansi

import (
	"errors"
	"unicode"
)

var (
	// ErrStopWalk is used as a return value from WalkFunc to indicate to Walk
	// to interrupt its iteration.
	ErrStopWalk = errors.New("stop walk")
)

// WalkFunc is the type of the function that is called while going through a
// string that might contains ANSI escape Codes.
// Walking occurs either rune per rune if the rune does not belong to an ANSI
// escape sequence or ANSI escape sequence per ANSI escape sequence in the
// other case. If WalkFunc is called for an ANSI Escape Scope, the curRune
// argument will be -1.
// If WalkFunc returns the ErrStopWalk error, Walk will interrupt the walk.
type WalkFunc func(curRune rune, curSGREsc string) error

// Walk walks through a string that can contain ANSI escape codes (starts by
// \x1b[ and ends with any letter or ~) and run WalkFunc either on each rune
// that is not part of an escape sequence or on each ANSI escape sequence.
func Walk(s string, fn WalkFunc) (err error) {
	var inEscSeq bool
	var curEsc string

Loop:
	for _, c := range s {
		switch {
		case c == '\x1b':
			inEscSeq, curEsc = true, string(c)

		case inEscSeq:
			curEsc += string(c)
			// TODO(pirmd): can probably do better than this to capture ANSI
			// CSI based sequence
			if unicode.IsLetter(c) || c == '~' {
				if err = fn(-1, curEsc); err != nil {
					break Loop
				}
				inEscSeq, curEsc = false, ""
			}

		default:
			if err = fn(c, ""); err != nil {
				break Loop
			}
		}
	}

	if err != nil {
		if err == ErrStopWalk {
			err = nil
		}

		return
	}

	if curEsc != "" {
		err = fn(-1, curEsc)
	}

	if err == ErrStopWalk {
		err = nil
	}

	return
}
