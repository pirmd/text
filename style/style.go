package style

import (
	"fmt"
)

//FormatFn represents a function that can transform a string to another one
//The transformation is usually used to apply a given style to a string
type FormatFn func(string) string

//FormatfFn is an extension of FormatFn to provide an interface similar to
//fmt.Sprintf
type FormatfFn func(format string, a ...interface{}) string

//FormatMap registers a set of FormatFn and maps them to a style's Format
type FormatMap map[Format]FormatFn

//must is a wrapper to transform a function that outputs (string, error) to a
//FormatFn signature.  It is done by capturing the error and feedbacking it to
//the output string (similar to fprintf principle)
func must(fn func(string) (string, error)) FormatFn {
	return func(s string) string {
		ts, err := fn(s)
		if err != nil {
			return fmt.Sprintf("!Error(%v):%s", err, s)
		}
		return ts
	}
}

//Sprintf is a wrapper around fmt.Sprintf that builds a FormatFn for a given
//format directive
func Sprintf(format string) FormatFn {
	return func(s string) string {
		return fmt.Sprintf(format, s)
	}
}

//Chain combines several FormatFn into one
func Chain(fn ...FormatFn) FormatFn {
	return func(src string) string {
		s := src
		for _, f := range fn {
			s = f(s)
		}
		return s
	}
}

//Chainf creates a new FormatfFn styling function by combining existing FormatFn
func Chainf(fn ...FormatfFn) FormatfFn {
	//TODO(pirmd): figure out how to prevent escape function to be applied each time 'f' is call
	return func(format string, a ...interface{}) string {
		s := fmt.Sprintf(format, a...)
		for _, f := range fn {
			s = f(s)
		}
		return s
	}
}

//Combine combines several FormatfFn into one FormatFn
func Combine(fn ...FormatfFn) FormatFn {
	return func(src string) string {
		s := src
		for _, f := range fn {
			s = f(s)
		}
		return s
	}
}

//CurrentStyler is the current selected Styler
//Default to the Term styler
var CurrentStyler = Term

//Styler repesents a collection of styling functions that can decorate a string
//apply a given text format.  Styler usually implements a given markup idiom
//but is pretty flexible to support things like decorating a text with ANSI
//colors escape sequences.
type Styler struct {
	fmtMap map[Format]FormatFn
}

//New creates a new Styler
func New(fmtMap FormatMap) *Styler {
	return &Styler{
		fmtMap: fmtMap,
	}
}

//Extend dupplicates current Styler and extends it.
//Existing styles will be overriden by the ones provided
func (st *Styler) Extend(fmtMap FormatMap) *Styler {
	stext := New(FormatMap{})
	for f, fn := range st.fmtMap {
		stext.fmtMap[f] = fn
	}

	for f, fn := range fmtMap {
		stext.fmtMap[f] = fn
	}
	return stext
}

//get retrieves a FormatFn from the given format.  If no FormatFn exists for this
//format, it returns a "do nothing" function that returns the input string
//"as-is".
func (st *Styler) get(f Format) FormatFn {
	if fn, ok := st.fmtMap[f]; ok {
		return fn
	}
	return func(s string) string { return s }
}

//style retrieves a FormatFn from the markup definition and ensures that
//escaping and auto-styling functions are applied to input text, if any.
func (st *Styler) style(f Format) FormatFn {
	fn := st.get(f)
	return func(s string) string {
		s = st.get(FmtEscape)(s)
		s = st.get(FmtAuto)(s)
		return fn(s)
	}
}

//stylef wraps style to offer an interface similar to fmt.Sprintf
func (st *Styler) stylef(f Format) FormatfFn {
	return func(format string, a ...interface{}) string {
		return st.style(f)(fmt.Sprintf(format, a...))
	}
}
