package text

//diff proposes simple and naive functions to visualize differences between
//strings. It probably is only working to have some eyecandies when loking at
//test results or to support some command line interactions.

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

type diffType int

const (
	isSame diffType = iota
	isDiff
	onlyLeft
	onlyRight

	minSimilarity int = 20
)

var (
	//Diff does not highlight differences
	Diff = DiffHighlighter{
		SameL:   func(s string) string { return s },
		DiffL:   func(s string) string { return s },
		LackL:   func(s string) string { return "" },
		ExcessL: func(s string) string { return s },

		SameR:   func(s string) string { return s },
		DiffR:   func(s string) string { return s },
		LackR:   func(s string) string { return "" },
		ExcessR: func(s string) string { return s },

		Symbols: [...]string{"=", "<>", "-", "+"},
	}

	//LightMkdDiff highlights differences in pseudo Markdonw
	LightMkdDiff = DiffHighlighter{
		SameL:   func(s string) string { return s },
		DiffL:   func(s string) string { return fmt.Sprintf("**%s**", s) },
		LackL:   func(s string) string { return "" },
		ExcessL: func(s string) string { return fmt.Sprintf("~~%s~~", s) },

		SameR:   func(s string) string { return s },
		DiffR:   func(s string) string { return fmt.Sprintf("**%s**", s) },
		LackR:   func(s string) string { return "" },
		ExcessR: func(s string) string { return fmt.Sprintf("~~%s~~", s) },

		Symbols: [...]string{"=", "<>", "-", "+"},
	}

	//ColorDiff highlights differences in colors
	ColorDiff = DiffHighlighter{
		SameL:   func(s string) string { return s },
		DiffL:   func(s string) string { return fmt.Sprintf("\x1b[31m%s\x1b[0m", s) }, //Red
		LackL:   func(s string) string { return "" },
		ExcessL: func(s string) string { return fmt.Sprintf("\x1b[31m%s\x1b[0m", s) }, //Red

		SameR:   func(s string) string { return s },
		DiffR:   func(s string) string { return fmt.Sprintf("\x1b[31m%s\x1b[0m", s) }, //Red
		LackR:   func(s string) string { return "" },
		ExcessR: func(s string) string { return fmt.Sprintf("\x1b[31m%s\x1b[0m", s) }, //Red

		Symbols: [...]string{"=", "<>", "-", "+"},
	}

	//DeviationDiff highlights differences in colors showing differences as deviations
	DeviationDiff = DiffHighlighter{
		SameL:   func(s string) string { return s },
		DiffL:   func(s string) string { return fmt.Sprintf("\x1b[32m%s\x1b[0m", s) }, //Green
		LackL:   func(s string) string { return "" },
		ExcessL: func(s string) string { return fmt.Sprintf("\x1b[32;1m%s\x1b[0m", s) }, //Green + Bold

		SameR: func(s string) string { return s },
		DiffR: func(s string) string { return fmt.Sprintf("\x1b[31m%s\x1b[0m", s) }, //Red
		LackR: func(s string) string { //Underlined space
			if len(s) > 0 {
				return fmt.Sprintf("\x1b[31;4m%s\x1b[0m", strings.Repeat(" ", len(s)-1))
			}
			return ""
		},
		ExcessR: func(s string) string { return fmt.Sprintf("\x1b[31;9m%s\x1b[0m", s) }, //Red + Strikeout

		Symbols: [...]string{"=", "<>", "-", "+"},
	}
)

//DiffHighlighter gathers styling functions to highlight differences between
//strings All functions have to be defined otherwise diff will panic
type DiffHighlighter struct {
	SameL, DiffL, LackL, ExcessL func(s string) string
	SameR, DiffR, LackR, ExcessR func(s string) string
	Symbols                      [4]string //Same order than diffType constants
}

//Slices returns the differences between two slices of strings Differences are
//looked at line by line, then word by word and finally rune by rune
func (h *DiffHighlighter) Slices(left, right []string) (dT []string, dL []string, dR []string) {
	d := findDiffAndInsertion(left, right)

	for i := range d.T {
		switch d.T[i] {
		case isDiff:
			_, lwords, rwords := h.Bywords(d.L[i], d.R[i])
			dT, dL, dR = append(dT, h.Symbols[isDiff]), append(dL, strings.Join(lwords, "")), append(dR, strings.Join(rwords, ""))

		default:
			ht, hl, hr := h.highlights(d.T[i], d.L[i], d.R[i])
			dT, dL, dR = append(dT, ht), append(dL, hl), append(dR, hr)
		}
	}

	return
}

//Anything returns a visualisation of the differences between two objects of
//unknown types.  It 'stringifies' these interface{} using a human friendly
//Json representation (or if not possible golang internal string (Gostring)
//then proced to a string diff
func (h *DiffHighlighter) Anything(l, r interface{}) (diffType []string, diffLeft []string, diffRight []string) {
	return h.Bylines(stringify(l), stringify(r))
}

//Bylines returns the differences between 'left' and 'right' strings. Differences are looked at
//line by line, then word by word and finally rune by rune
func (h *DiffHighlighter) Bylines(left, right string) (dT []string, dL []string, dR []string) {
	Llines, Rlines := strings.Split(left, "\n"), strings.Split(right, "\n")
	return h.Slices(Llines, Rlines)
}

//Bywords returns the differences between 'left' and 'right' strings.
//Differences are looked at word by word, then rune by rune.
func (h *DiffHighlighter) Bywords(left, right string) (dT []string, dL []string, dR []string) {
	Lwords, Rwords := splitInWords(left), splitInWords(right)
	d := findDiffAndInsertion(Lwords, Rwords).Compact()

	if d.SimilarityLevel() < minSimilarity {
		ht, hl, hr := h.highlights(isDiff, left, right)
		return append(dT, ht), append(dL, hl), append(dR, hr)
	}

	for i := range d.T {
		switch d.T[i] {
		case isDiff:
			_, lrunes, rrunes := h.Byrunes(d.L[i], d.R[i])
			dT, dL, dR = append(dT, h.Symbols[isDiff]), append(dL, strings.Join(lrunes, "")), append(dR, strings.Join(rrunes, ""))

		default:
			ht, hl, hr := h.highlights(d.T[i], d.L[i], d.R[i])
			dT, dL, dR = append(dT, ht), append(dL, hl), append(dR, hr)
		}
	}

	return
}

//Byrunes returns the differences between 'left' and 'right'. Differences are
//looked at rune by rune
func (h *DiffHighlighter) Byrunes(left, right string) (dT []string, dL []string, dR []string) {
	Lrunes, Rrunes := strings.Split(left, ""), strings.Split(right, "")
	d := findDiff(Lrunes, Rrunes).Compact()

	if d.SimilarityLevel() < minSimilarity {
		ht, hl, hr := h.highlights(isDiff, left, right)
		return append(dT, ht), append(dL, hl), append(dR, hr)
	}

	for i := range d.T {
		ht, hl, hr := h.highlights(d.T[i], d.L[i], d.R[i])
		dT, dL, dR = append(dT, ht), append(dL, hl), append(dR, hr)
	}

	return
}

func (h *DiffHighlighter) highlights(dT diffType, dL, dR string) (t string, l string, r string) {
	switch dT {
	case isSame:
		t, l, r = h.Symbols[dT], h.SameL(dL), h.SameR(dR)
	case isDiff:
		t, l, r = h.Symbols[dT], h.DiffL(dL), h.DiffR(dR)
	case onlyLeft:
		t, l, r = h.Symbols[dT], h.ExcessL(dL), h.LackR(dR)
	case onlyRight:
		t, l, r = h.Symbols[dT], h.LackL(dL), h.ExcessR(dR)
	}

	return
}

func findDiff(left, right []string) *diff {
	d := &diff{}
	i := 0
	lenL, lenR := len(left), len(right)

	for {
		if i == lenL {
			for n := i; n < lenR; n++ {
				d.OnlyRight(right[n])
			}
			break
		}

		if i == lenR {
			for n := i; n < lenL; n++ {
				d.OnlyLeft(left[n])
			}
			break
		}

		if left[i] == right[i] {
			d.IsSame(left[i])
			i++
			continue
		}

		d.IsDifferent(left[i], right[i])
		i++
	}

	return d
}

func findDiffAndInsertion(left, right []string) *diff {
	d := &diff{}
	iL, iR := 0, 0
	lenL, lenR := len(left), len(right)

Loop:
	for {
		if iL == lenL {
			for n := iR; n < lenR; n++ {
				d.OnlyRight(right[n])
			}
			break
		}

		if iR == lenR {
			for n := iL; n < lenL; n++ {
				d.OnlyLeft(left[n])
			}
			break
		}

		for n := iR; n < lenR; n++ {
			if left[iL] == right[n] {
				for p := iR; p < n; p++ {
					d.OnlyRight(right[p])
				}
				d.IsSame(left[iL])
				iL++
				iR = n + 1
				goto Loop
			}
		}

		for n := iL; n < lenL; n++ {
			if left[n] == right[iR] {
				for p := iL; p < n; p++ {
					d.OnlyLeft(left[p])
				}
				d.IsSame(right[iR])
				iR++
				iL = n + 1
				goto Loop
			}
		}

		d.IsDifferent(left[iL], right[iR])
		iL++
		iR++
	}

	return d
}

//diff captures differences between two strings
type diff struct {
	T    []diffType
	L, R []string
}

func (d *diff) IsSame(left string) {
	d.L = append(d.L, left)
	d.R = append(d.R, left)
	d.T = append(d.T, isSame)
}

func (d *diff) OnlyLeft(left string) {
	d.L = append(d.L, left)
	d.R = append(d.R, left)
	d.T = append(d.T, onlyLeft)
}

func (d *diff) OnlyRight(right string) {
	d.L = append(d.L, right)
	d.R = append(d.R, right)
	d.T = append(d.T, onlyRight)
}

func (d *diff) IsDifferent(left, right string) {
	d.L = append(d.L, left)
	d.R = append(d.R, right)
	d.T = append(d.T, isDiff)
}

func (d *diff) SimilarityLevel() int {
	same := 0
	for _, t := range d.T {
		if t == isSame {
			same++
		}
	}
	return int(100.0 * same / len(d.T))
}

func (d *diff) Compact() *diff {
	var t []diffType
	var r, l []string
	var curT diffType
	var curL, curR string

	for i := range d.T {
		if i > 0 && d.T[i] != curT {
			t, l, r = append(t, curT), append(l, curL), append(r, curR)
			curT, curL, curR = d.T[i], "", ""
		}
		curT, curL, curR = d.T[i], curL+d.L[i], curR+d.R[i]
	}

	t, l, r = append(t, curT), append(l, curL), append(r, curR)
	return &diff{t, l, r}
}

func splitInWords(s string) (split []string) {
	var lasti int

	for i, r := range s {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			split = append(split, s[lasti:i])
			split = append(split, string(r))
			lasti = i + 1
		}
	}

	if lasti < len(s) {
		split = append(split, s[lasti:])
	}

	return
}

func stringify(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)

	default:
		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return fmt.Sprintf("%#v", v)
		}
		return string(b)
	}
}
