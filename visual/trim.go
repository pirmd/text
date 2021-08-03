package visual

import (
	"bytes"
	"unicode"

	"github.com/pirmd/text/ansi"
)

// TrimSpace trims any leading and trailing space (as defined by Unicode) from
// a slice of bytes.
func TrimSpace(s []byte) []byte {
	var trimmed bytes.Buffer
	var spaceBuf bytes.Buffer
	var spaceBufSGR ansi.Sequence

	isLeadingSpaces := true
	_ = ansi.Walk(s, func(advance int, c rune, esc string) error {
		switch {
		case len(esc) > 0:
			spaceBuf.WriteString(esc)
			spaceBufSGR.Combine(esc)

		case unicode.IsSpace(c):
			if !isLeadingSpaces {
				spaceBuf.WriteRune(c)
			}

		default:
			isLeadingSpaces = false

			trimmed.Write(spaceBuf.Bytes())
			spaceBuf.Reset()
			spaceBufSGR.Reset()

			trimmed.WriteRune(c)
		}

		return nil
	})

	trimmed.WriteString(spaceBufSGR.String())
	return trimmed.Bytes()
}

// TrimLeadingSpace trims any leading space (as defined by Unicode) from
// a slice of bytes.
func TrimLeadingSpace(s []byte) []byte {
	var trimmed bytes.Buffer
	var spaceBufSGR ansi.Sequence

	isLeadingSpaces := true
	_ = ansi.Walk(s, func(advance int, c rune, esc string) error {
		switch {
		case len(esc) > 0:
			if isLeadingSpaces {
				spaceBufSGR.Combine(esc)
			} else {
				trimmed.WriteString(esc)
			}

		case unicode.IsSpace(c):
			if !isLeadingSpaces {
				trimmed.WriteRune(c)
			}

		default:
			isLeadingSpaces = false
			trimmed.WriteString(spaceBufSGR.String())
			spaceBufSGR.Reset()

			trimmed.WriteRune(c)
		}

		return nil
	})

	trimmed.WriteString(spaceBufSGR.String())
	return trimmed.Bytes()
}

// TrimTrailingSpace trims any trailing space (as defined by Unicode) from a
// slice of bytes.
func TrimTrailingSpace(s []byte) []byte {
	var trimmed bytes.Buffer
	var spaceBuf bytes.Buffer
	var spaceBufSGR ansi.Sequence

	_ = ansi.Walk(s, func(advance int, c rune, esc string) error {
		switch {
		case len(esc) > 0:
			spaceBuf.WriteString(esc)
			spaceBufSGR.Combine(esc)

		case unicode.IsSpace(c):
			spaceBuf.WriteRune(c)

		default:
			trimmed.Write(spaceBuf.Bytes())
			spaceBuf.Reset()
			spaceBufSGR.Reset()

			trimmed.WriteRune(c)
		}

		return nil
	})

	trimmed.WriteString(spaceBufSGR.String())
	return trimmed.Bytes()
}

// TrimSuffix trims trailing rune r.
func TrimSuffix(s []byte, r rune) []byte {
	var trimmed bytes.Buffer
	var buf bytes.Buffer
	var bufSGR ansi.Sequence

	_ = ansi.Walk(s, func(advance int, c rune, esc string) error {
		switch {
		case len(esc) > 0:
			buf.WriteString(esc)
			bufSGR.Combine(esc)

		case c == r:
			buf.WriteRune(c)

		default:
			trimmed.Write(buf.Bytes())
			buf.Reset()
			bufSGR.Reset()

			trimmed.WriteRune(c)
		}

		return nil
	})

	trimmed.WriteString(bufSGR.String())
	return trimmed.Bytes()
}
