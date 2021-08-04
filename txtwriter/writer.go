package txtwriter

import (
	"bytes"
	"io"
)

// Writer represents a text's Writer that knows how to format input text that
// possibly contains ANSI escape sequences.
//
// Writer handles text by lines to implement the proper alignment directives.
// You might need to use Flush() to ensure any none-terminated line (line that
// is shorter than the maximum Writer's width or not yet ended by a '\n') is
// properly writing to the output.
type Writer struct {
	out io.Writer

	// Writer's options

	maxwidth      int
	indentScheme  []byte
	blockindent   bool
	lazywrap      bool
	padWithSpaces bool
	alignment     alignment
	prefixes      []string

	//TODO: add support to interrupt ANSI at each line (get inspiration from
	//table module).
	//TODO: add support to strip ANSI sequence or replace ANSI sequence by
	//other idiom's marker.

	// Writer's status

	curline     *bytes.Buffer
	indentLvl   int
	indent      []byte
	prefixIdx   int
	prefix      []byte
	prefixWidth int
}

// New creates a new Writer.
func New(out io.Writer) *Writer {
	return &Writer{
		out:          out,
		indentScheme: defaultIndentScheme,
		curline:      new(bytes.Buffer),
	}
}

// Write appends the content of p to Writer's current line. Current line is
// written to Writer's output when the current line reaches Writer's maximum
// width (if any) or when an end-of-line is encountered.
func (w *Writer) Write(p []byte) (n int, err error) {
	n, err = w.writeWithWrap(p)
	return
}

// Flush writes the remaining bytes that can be living in the Writer's buffers,
// notably incomplete lines or words (line that a have not reached maximum
// width yet or end-of-line).
func (w *Writer) Flush() error {
	return w.writeLine(w.curline.Bytes())
}

// writeLine writes a line to Writer's output and take care of Writer's
// formatting options.
//
// writeLine assumes p is trimmed from leading and trailing spaces notably p is
// not ending by '\n' or newline/padding will be voided.  It is safe here as
// line is trimmed from spaces by Writer.Write before being write to Writer's
// output.
//
// BUG: when aligning and padding, current ANSI sequence is not interrupted so
// that added spaces can be made visible.
func (w *Writer) writeLine(p []byte) (err error) {
	if _, err = w.writeIndentation(); err != nil {
		return
	}

	if _, err = w.writePrefix(); err != nil {
		return
	}

	if _, err = w.alignLine(p); err != nil {
		return
	}

	if _, err = w.out.Write(p); err != nil {
		return err
	}

	if w.padWithSpaces {
		if _, err = w.padLine(p); err != nil {
			return
		}
	}

	return
}

func (w Writer) width() int {
	if w.blockindent {
		return w.maxwidth - 2*len(w.indent) - w.prefixWidth
	}
	return w.maxwidth - len(w.indent) - w.prefixWidth
}
