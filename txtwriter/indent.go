package txtwriter

var (
	//defaultIndentScheme is the default indent.
	defaultIndentScheme = []byte{' ', ' ', ' ', ' '}
)

// Indent increases the indent level (up to Writer's maximum width if any).
func (w *Writer) Indent() {
	if w.width() < len(w.indentScheme) {
		return
	}

	w.indentLvl++
	w.indent = append(w.indent, w.indentScheme...)
}

// Unindent decreases the indent level.
func (w *Writer) Unindent() {
	if w.indentLvl == 0 {
		return
	}

	w.indentLvl--
	w.indent = w.indent[:len(w.indent)-len(w.indentScheme)]
}

// BlockIndent applies indent level to both left- and right-margin.
func (w *Writer) BlockIndent() {
	w.blockindent = true
}

// Unblockindent applies indent level to left-margin only 'default behavior).
func (w *Writer) Unblockindent() {
	w.blockindent = false

}

func (w *Writer) writeIndentation() (int, error) {
	return w.out.Write(w.indent)
}
