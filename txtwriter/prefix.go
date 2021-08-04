package txtwriter

import (
	"github.com/pirmd/text/visual"
)

// SetPrefix defines prefixes that are added at the start of each text line.
// Prefixes are used in the given order, if there are more lines than prefixes,
// last prefix is repeated.
// Prefix sequence order needs to be reset manually using ResetPrefix().
func (w *Writer) SetPrefix(prefixes ...string) *Writer {
	w.prefixes = prefixes
	w.ResetPrefix()
	return w
}

// ResetPrefix resets prefix's sequence.
func (w *Writer) ResetPrefix() {
	w.prefixIdx = -1

	// do not reset prefix if we have element in buffer, reset it just after.
	if w.curline.Len() == 0 {
		w.prefix = []byte{}
		w.prefixWidth = 0
		w.nextPrefix()
	}
}

func (w *Writer) nextPrefix() {
	if w.prefixIdx < len(w.prefixes)-1 {
		w.prefixIdx++
		w.prefix = []byte(w.prefixes[w.prefixIdx])
		w.prefixWidth = visual.Width(w.prefix)
	}
}

func (w *Writer) writePrefix() (int, error) {
	n, err := w.out.Write(w.prefix)
	if err != nil {
		return n, err
	}

	w.nextPrefix()
	return n, nil
}
