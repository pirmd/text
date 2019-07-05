package style

import (
	"fmt"
	"strings"
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

//Chainf creates a new FormatfFn styling function by combining provided
//FormatFn
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
//XXX: is it used?
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
//to apply a given text format.
//
//Styler is pretty flexible and can be customized through a FormatMap for
//simple text decoration cases. More advanced text formatting are provided for
//table
type Styler struct {
	fmtMap map[Format]FormatFn

	tabFn      func(int) FormatFn
	listFn     func(int) func(...string) string
	listItemFn func(int) FormatFn
	tableFn    func(...[]string) string
	defineFn   func(string, string) string
}

//New creates a new Styler.
//Provided tableFn  gives a recipe to build a table from the provided rows
func New(fmtMap FormatMap, tabFn func(int) FormatFn, listFn func(int) func(...string) string, listItemFn func(int) FormatFn, tableFn func(...[]string) string, defineFn func(string, string) string) *Styler {
	return &Styler{
		fmtMap:     fmtMap,
		tabFn:      tabFn,
		listFn:     listFn,
		listItemFn: listItemFn,
		tableFn:    tableFn,
		defineFn:   defineFn,
	}
}

//Extend dupplicates current Styler and extends it.
//Existing formats will be overriden by the ones provided. Extended styling
//functions (like table drawing) will be replaced by the provided ones when not
//nil.
func (st *Styler) Extend(stadd *Styler) *Styler {
	stext := New(FormatMap{}, st.tabFn, st.listFn, st.listItemFn, st.tableFn, st.defineFn)
	for f, fn := range st.fmtMap {
		stext.fmtMap[f] = fn
	}

	for f, fn := range stadd.fmtMap {
		stext.fmtMap[f] = fn
	}
	if stadd.tabFn != nil {
		stext.tabFn = stadd.tabFn
	}
	if stadd.listFn != nil {
		stext.listFn = stadd.listFn
	}
	if stadd.listItemFn != nil {
		stext.listItemFn = stadd.listItemFn
	}
	if stadd.tableFn != nil {
		stext.tableFn = stadd.tableFn
	}
	if stadd.defineFn != nil {
		stext.defineFn = stadd.defineFn
	}

	return stext
}

//Table draws a table according to style's table drawing function.
//If no table drawing function exists, it simply outputs "|"-separated text per
//line for each row.
//For styles where table depends on text width to adjust columns width, it is
//not advised to chained it with tabulation or indentation-based formats.
func (st *Styler) Table(rows ...[]string) string {
	var autoRows [][]string
	for _, row := range rows {
		var autoRow []string
		for _, cell := range row {
			autoRow = append(autoRow, st.auto(cell))
		}
		autoRows = append(autoRows, autoRow)
	}

	if st.tableFn == nil {
		//Default poor-man table ("|"-separated columns)
		var r []string
		for _, row := range autoRows {
			r = append(r, strings.Join(row, " | "))
		}

		var s string
		for _, row := range r {
			s += st.Line(row)
		}
		return st.Paragraph(s)
	}

	return st.tableFn(autoRows...)
}

//Table draws a table using the current Styler
func Table(rows ...[]string) string {
	return CurrentStyler.Table(rows...)
}

//Tab returns a FormatfFn that indents then wraps the provided string by the
//given level.
//If no table drawing function exists, it returns a "do nothing" function that
//returns the input string "as-is".
func (st *Styler) Tab(level int) FormatfFn {
	if st.tabFn == nil {
		return func(format string, a ...interface{}) string {
			s := fmt.Sprintf(format, a...)
			return st.auto(s)
		}
	}

	return func(format string, a ...interface{}) string {
		s := fmt.Sprintf(format, a...)
		return st.tabFn(level)(st.auto(s))
	}
}

//Tab indents then wraps the text using the current Styler
func Tab(level int) FormatfFn {
	return CurrentStyler.Tab(level)
}

//List creates a new with the given nested level.
//If no list drawing function exists, it returns one line per item.
func (st *Styler) List(level int) func(...string) string {
	if st.listFn == nil {
		return func(items ...string) string {
			return strings.Join(items, "\n")
		}
	}

	return func(items ...string) string {
		return st.listFn(level)(items...)
	}
}

//List creates a new list with the given nest-level using the current Styler
func List(level int) func(...string) string {
	return CurrentStyler.List(level)
}

//ListItem creates a new list item of teh given nest-level.
//If no list item function exists, it returns the input preceded by a "-"
func (st *Styler) ListItem(level int) FormatFn {
	if st.listItemFn == nil {
		return func(s string) string {
			return Tab(level)("- " + s)
		}
	}

	return func(s string) string {
		return st.listItemFn(level)(st.auto(s))
	}
}

//ListItem creates a new list item of the given nest-level using the current
//Styler.
func ListItem(level int) FormatFn {
	return CurrentStyler.ListItem(level)
}

//Define adds a term definition
//If no define function exists, it simply outputs term: description
//
//LIMITATION: Defines might already use Tab internally and threfore does not
//chain easily with another styling funvtion that use Wrap or Tab
func (st *Styler) Define(term, desc string) string {
	term, desc = st.auto(term), st.auto(desc)

	if st.defineFn == nil {
		return term + ": " + desc
	}

	return st.defineFn(term, desc)
}

//Define adds a term definition using current Styler
func Define(term, desc string) string {
	return CurrentStyler.Define(term, desc)
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
		return fn(st.auto(s))
	}
}

//stylef wraps style to offer an interface similar to fmt.Sprintf
func (st *Styler) stylef(f Format) FormatfFn {
	return func(format string, a ...interface{}) string {
		return st.style(f)(fmt.Sprintf(format, a...))
	}
}

//auto run automatic formatting
func (st *Styler) auto(s string) string {
	s = st.get(FmtEscape)(s)
	s = st.get(FmtAuto)(s)
	return s
}
