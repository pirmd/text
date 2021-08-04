package txtwriter

import (
	"bytes"

	"github.com/pirmd/text/visual"
)

type alignment int

const (
	alignLeft alignment = iota
	alignRight
	alignCenter
)

func (a alignment) String() string {
	return [...]string{"Left", "Right", "Center"}[a]
}

// AlignLeft will align text to the left (default).
func (w *Writer) AlignLeft() *Writer {
	w.alignment = alignLeft
	return w
}

// AlignRight will align text to the right.
func (w *Writer) AlignRight() *Writer {
	w.alignment = alignRight
	return w
}

// AlignCenter will center text.
func (w *Writer) AlignCenter() *Writer {
	w.alignment = alignCenter
	return w
}

// PadLeft will ad spaces  at the end of lines up-to reaching Writer's maximum
// visual width.
func (w *Writer) PadLeft() *Writer {
	w.padWithSpaces = true
	return w
}

// alignLine aligns a line to fit current Writer's format.
// It is assumed that p has been trimmed left and right from spaces (incl.
// eol). If not, as a consequences: visual result might be disappointing as
// leading or trailing existing spaces in p will not be considered when
// aligning p.
// BUG: alignment does not managed possible ANSI escape sequence so that added
// leading or trailing spaces might appear with the current ANSI formatting
// (like strike-out or in inverse color)
func (w *Writer) alignLine(p []byte) (int, error) {
	if freespace := w.width() - visual.Width(p); freespace > 0 {
		switch w.alignment {
		case alignRight:
			return w.out.Write(bytes.Repeat([]byte{' '}, freespace))

		case alignCenter:
			return w.out.Write(bytes.Repeat([]byte{' '}, freespace/2))
		}
	}

	return 0, nil
}

// padLine pads a line to fit current Writer's format.  It is assumed that p
// has been trimmed left and right from spaces (incl.  eol). If not, as a
// consequences visual result might be disappointing as padding will will
// happen after possible '\n'
// BUG: alignment does not managed possible ANSI escape sequence so that added
// leading or trailing spaces might appear with the current ANSI formatting
// (like strike-out or in inverse color)
func (w *Writer) padLine(p []byte) (int, error) {
	if freespace := w.width() - visual.Width(p); freespace > 0 {
		switch w.alignment {
		case alignLeft:
			return w.out.Write(bytes.Repeat([]byte{' '}, freespace))

		case alignCenter:
			return w.out.Write(bytes.Repeat([]byte{' '}, freespace-(freespace/2)))
		}
	}

	return 0, nil
}
