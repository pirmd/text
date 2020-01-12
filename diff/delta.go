package diff

import (
	"fmt"
	"strings"
)

// Type represents the differences's types that can be encountered
type Type int

const (
	// IsUnknown when status is not known (usuallt initialisation state)
	IsUnknown Type = iota
	// IsSame when strings are the same
	IsSame
	// IsDeleted when a string has been deleted
	IsDeleted
	// IsInserted when a string has been inserted
	IsInserted
	// IsDifferent when two sets of strings are differents
	IsDifferent
)

// String represents a difference's Type in an easy to understand format.
func (typ Type) String() string {
	return [...]string{"?", "=", "-", "+", "<>"}[typ]
}

// Delta represents an atomic piece of diff
type Delta interface {
	// Type is the kind of difference for Delta
	Type() Type

	left() (string, bool)
	right() (string, bool)
	prettyPrint(highlighters) (string, string)
}

// Result gathers any diff results
type Result []Delta

// Type returns the difference's Type of Result
func (r Result) Type() Type {
	var dT []Type
	for _, delta := range r {
		dT = append(dT, delta.Type())
	}

	return cumulTypes(dT...)
}

// PrettyPrint translates a difference's result into a human (or machine)
// readable text. It outputs a representation of the differences for the first
// reference string, for the second one as well as a representation of the type
// of difference.
//
// Output format depends on the selected Highlighter(s) if any.
//
// LIMITATION: whereas representation of left and right strings are working in
// all cases, the representation of type of difference will provide tedious
// output if difference was not computed by lines.
func (r Result) PrettyPrint(h ...Highlighter) (dL string, dR string, dT string) {
	for _, delta := range r {
		diffL, diffR := delta.prettyPrint(h)
		diffT := highlighters(h).formatT(delta.Type())
		diffL, diffR, diffT = alignNumberOfLines(diffL, diffR, diffT)
		dL, dR, dT = dL+diffL, dR+diffR, dT+diffT
	}

	return
}

// GoString represents a diff's Result in an easy to read format.
func (r Result) GoString() (s string) {
	for i, delta := range r {
		if i > 0 {
			s += "\n"
		}
		s += fmt.Sprintf("%#v", delta)
	}
	return
}

func (r *Result) append(deltas ...Delta) {
	*r = append(*r, deltas...)
}

func (r *Result) insert(deltas ...Delta) {
	*r = append(deltas, *r...)
}

// differentZones returns zones that are a mix of deletions and insertions.
// Zones are identified by their start and end indexes.
func (r Result) differentZones() (zones [][2]int) {
	var curType Type
	var curZone int

	for i, delta := range r {
		switch diffT := delta.Type(); curType {
		case IsDifferent:
			if diffT == IsSame || diffT == IsUnknown {
				zones = append(zones, [2]int{curZone, i - 1})
				curType = diffT
			}

		case IsInserted, IsDeleted:
			if diffT == IsSame || diffT == IsUnknown {
				curType, curZone = diffT, i
			} else if diffT != curType {
				curType = IsDifferent
			}

		default:
			curType, curZone = diffT, i
		}
	}

	if curType == IsDifferent {
		zones = append(zones, [2]int{curZone, len(r) - 1})
	}

	return
}

func (r Result) content() (dL []string, dR []string) {
	for _, delta := range r {
		if diffL, exists := delta.left(); exists {
			dL = append(dL, diffL)
		}
		if diffR, exists := delta.right(); exists {
			dR = append(dR, diffR)
		}
	}
	return
}

// implements interface Delta so that we can stack different levels of Results
func (r Result) left() (dL string, ok bool) {
	for _, delta := range r {
		if diffL, exists := delta.left(); exists {
			dL += diffL
			ok = true
		}
	}
	return
}

// implements interface Delta so that we can stack different levels of Results
func (r Result) right() (dR string, ok bool) {
	for _, delta := range r {
		if diffR, exists := delta.right(); exists {
			dR += diffR
			ok = true
		}
	}
	return
}

// implements interface Delta so that we can stack different levels of Results
func (r Result) prettyPrint(h highlighters) (dL string, dR string) {
	for _, delta := range r {
		diffL, diffR := delta.prettyPrint(h)
		dL, dR = dL+diffL, dR+diffR
	}
	return
}

// diff represents an atomic piece of difference between strings
type diff struct {
	operation Type
	content   string
}

func newSameDiff(text string) *diff {
	return &diff{
		operation: IsSame,
		content:   text,
	}
}

func newInsertedDiff(text string) *diff {
	return &diff{
		operation: IsInserted,
		content:   text,
	}
}

func newDeletedDiff(text string) *diff {
	return &diff{
		operation: IsDeleted,
		content:   text,
	}
}

func (d diff) Type() Type {
	return d.operation
}

func (d diff) GoString() string {
	return fmt.Sprintf("(%2s) %#v", d.operation, d.content)
}

func (d diff) left() (string, bool) {
	if d.operation == IsInserted {
		return "", false
	}

	return d.content, true
}

func (d diff) right() (string, bool) {
	if d.operation == IsDeleted {
		return "", false
	}

	return d.content, true
}

func (d diff) prettyPrint(h highlighters) (string, string) {
	l, _ := d.left()
	r, _ := d.right()
	return h.formatLR(l, r, d.operation)
}

func cumulTypes(types ...Type) (t Type) {
	for _, dT := range types {
		switch {
		case dT == t:

		case t == IsUnknown || t == IsSame:
			t = dT

		case dT == IsSame:

		default:
			t = IsDifferent
		}
	}
	return
}

func alignNumberOfLines(dL, dR, dT string) (string, string, string) {
	diffL, diffR, diffT := dL, dR, dT
	nL, nR, nT := strings.Count(dL, "\n"), strings.Count(dR, "\n"), strings.Count(dT, "\n")

	maxN := nL
	if nR > maxN {
		maxN = nR
	}
	if nT > maxN {
		maxN = nT
	}

	if nL < maxN {
		diffL += strings.Repeat("\n", maxN-nL)
	}
	if nR < maxN {
		diffR += strings.Repeat("\n", maxN-nR)
	}
	if nT < maxN {
		diffT += strings.Repeat("\n", maxN-nT)
	}

	return diffL, diffR, diffT
}
