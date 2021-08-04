package txtwriter

import (
	"bytes"
	"unicode"

	"github.com/pirmd/text/ansi"
	"github.com/pirmd/text/visual"
)

// SetMaxWidth sets Writer's maximum width.
// Setting maxwith to a value less or equal to zero disable wrapping as well as
// alignment and padding steps.
func (w *Writer) SetMaxWidth(maxwidth int) *Writer {
	w.maxwidth = maxwidth

	return w
}

// LazyWrap sets Writer's wrapping mode to not break in the middle of words
// whose "visual" width are above Writer's maximum width.
func (w *Writer) LazyWrap() *Writer {
	w.lazywrap = true

	return w
}

func (w Writer) newCutter() *visual.Cutter {
	if w.lazywrap {
		return visual.NewLazyCutter(w.width())
	}

	return visual.NewCutter(w.width())

}

func (w *Writer) writeWithWrap(p []byte) (n int, err error) {
	cutr := w.newCutter()

	// Append the current buffer to p to catch-up with previous wrapping
	// operation that did not fully complete.
	in := append(w.curline.Bytes(), p...)

	in = trimLeadingSpace(in)
	line, tail := cutr.Split(in)
	for line != nil {
		// trim every trailing spaces and eol from line to get a cleaner visual output.
		// It does simplify the life of writeLine in getting proper
		// left/right/center alignment and padding. It also ensures that we
		// have to add a '\n' without worrying doubling it if already existing.
		line = visual.TrimTrailingSpace(line)
		if err = w.writeLine(line); err != nil {
			return
		}
		if _, err = w.out.Write([]byte{'\n'}); err != nil {
			return
		}

		// trim leading spaces from tail to get a proper visual alignment feeling.
		in = trimLeadingSpace(tail)
		line, tail = cutr.Split(in)
	}

	// Memorize remaining text that is not detected as either being ended by a
	// end-of-line or longer than maximum width.
	// Writer.Flush() is responsible to eventually output its content to
	// Writer.out.
	w.curline.Reset()
	w.curline.Write(tail)

	// TODO: how to properly capture the number of bytes actually processed by
	// Write (taking into account trimming operation and initial addition of
	// pending buffer)?
	n = len(p)

	return
}

// trimLeadingSpace is a copy of visual.TrimLeadingSpace that does not consider
// eol as a space.
func trimLeadingSpace(s []byte) []byte {
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
			if c == '\n' || !isLeadingSpaces {
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
