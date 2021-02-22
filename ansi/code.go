package ansi

import (
	"strings"
)

// Code represents an ANSI code.
type Code = string

const (
	cCSI = "\x1b["

	cReset      Code = "0"
	cBold       Code = "1"
	cFaint      Code = "2"
	cItalic     Code = "3"
	cUnderline  Code = "4"
	cSlowBlink  Code = "5"
	cRapidBlink Code = "6"
	cInverse    Code = "7"
	cConceal    Code = "8"
	cCrossedOut Code = "9"

	cBoldOff       Code = "21"
	cNormal        Code = "22"
	cItalicOff     Code = "23"
	cUnderlineOff  Code = "24"
	cBlinkOff      Code = "25"
	cInverseOff    Code = "27"
	cReveal        Code = "28"
	cNotCrossedOut Code = "29"

	cBlack      Code = "30"
	cRed        Code = "31"
	cGreen      Code = "32"
	cYellow     Code = "33"
	cBlue       Code = "34"
	cMagenta    Code = "35"
	cCyan       Code = "36"
	cWhite      Code = "37"
	cSetFGColor Code = "38"
	cDefaultFG  Code = "39"

	cBlackBG    Code = "40"
	cRedBG      Code = "41"
	cGreenBG    Code = "42"
	cYellowBG   Code = "43"
	cBlueBG     Code = "44"
	cMagentaBG  Code = "45"
	cCyanBG     Code = "46"
	cWhiteBG    Code = "47"
	cSetBGColor Code = "48"
	cDefaultBG  Code = "49"

	cFramed       Code = "51"
	cEncircled    Code = "52"
	cOverlined    Code = "53"
	cNotFramed    Code = "54"
	cNotOverlined Code = "55"

	cBrightBlack   Code = "90"
	cBrightRed     Code = "91"
	cBrightGreen   Code = "92"
	cBrightYellow  Code = "93"
	cBrightBlue    Code = "94"
	cBrightMagenta Code = "95"
	cBrightCyan    Code = "96"
	cBrightWhite   Code = "97"

	cBrightBlackBG   Code = "100"
	cBrightRedBG     Code = "101"
	cBrightGreenBG   Code = "102"
	cBrightYellowBG  Code = "103"
	cBrightBlueBG    Code = "104"
	cBrightMagentaBG Code = "105"
	cBrightCyanBG    Code = "106"
	cBrightWhiteBG   Code = "107"
)

// isFGColor is true if the corresponding Code allows to modify fore-ground
// colors.
func isFGColor(c Code) bool {
	return c == cBlack || c == cRed || c == cGreen || c == cYellow || c == cBlue || c == cMagenta || c == cCyan || c == cWhite ||
		c == cBrightBlack || c == cBrightRed || c == cBrightGreen || c == cBrightYellow || c == cBrightBlue || c == cBrightMagenta || c == cBrightCyan || c == cBrightWhite ||
		c == cSetFGColor || c == cDefaultFG
}

// isBGColor is true if the corresponding Code allows to modify back-ground
// colors.
func isBGColor(c Code) bool {
	return c == cBlackBG || c == cRedBG || c == cGreenBG || c == cYellowBG || c == cBlueBG || c == cMagentaBG || c == cCyanBG || c == cWhiteBG ||
		c == cBrightBlackBG || c == cBrightRedBG || c == cBrightGreenBG || c == cBrightYellowBG || c == cBrightBlueBG || c == cBrightMagentaBG || c == cBrightCyanBG || c == cBrightWhiteBG ||
		c == cSetBGColor || c == cDefaultBG
}

// isStyleOff is true if the corresponding Code allows to set off graphic
// renditions style like Bold, Underline, Faint and so on.
func isStyleOff(c Code) bool {
	return c == cReset || c == cBoldOff || c == cNormal || c == cItalicOff ||
		c == cUnderlineOff || c == cBlinkOff || c == cInverseOff ||
		c == cReveal || c == cNotCrossedOut || c == cNotFramed || c == cNotOverlined ||
		c == cDefaultBG || c == cDefaultFG
}

// isSupersededBy is true if the corresponding Code visual effect is supeseded
// by the new Code (for example: Bold is superseded by BoldOff or Normal or
// Reset but not by Green).
func isSupersededBy(ca, cb Code) bool {
	return cb == cReset ||
		(isFGColor(cb) && isFGColor(ca)) ||
		(isBGColor(cb) && isBGColor(ca)) ||
		(cb == cBoldOff && ca == cBold) ||
		(cb == cNormal && (ca == cBold || ca == cFaint)) ||
		(cb == cItalicOff && ca == cItalic) ||
		(cb == cUnderlineOff && ca == cUnderline) ||
		(cb == cBlinkOff && (ca == cSlowBlink || ca == cRapidBlink)) ||
		(cb == cInverseOff && ca == cInverse) ||
		(cb == cReveal && ca == cConceal) ||
		(cb == cNotCrossedOut && ca == cCrossedOut) ||
		(cb == cNotOverlined && ca == cOverlined) ||
		(cb == cNotFramed && (ca == cFramed || ca == cEncircled))
}

// Sequence is a slice of ansi.Codes.
type Sequence []Code

// Combine adds a new ANSI SGR escape sequence. It deletes existing Codes that
// are superseded by new ones (like a Red by a Green) or as well as Codes that
// are not useful at the boundaries of the SGRSequence (like BoldOff without a
// BoldOn).
func (seq *Sequence) Combine(esc string) {
	for _, c := range ParseSGR(esc) {
		seq.add(c)
	}
}

// Esc returns the sequence's ANSI escape sequence.
func (seq Sequence) Esc() (s string) {
	if len(seq) != 0 {
		s = cCSI + strings.Join(seq, ";") + "m"
	}

	return
}

// Print decorates the provided string using the corresponding Code sequence.
func (seq Sequence) Print(s string) string {
	return seq.Esc() + s + cCSI + cReset + "m"
}

// Reset returns ANSI Reset code that nullifies any graphic rendition defined
// by sequence. If sequence is empty, returns an empty code.
func (seq Sequence) Reset() string {
	if len(seq) == 0 {
		return ""
	}
	return Reset
}

func (seq *Sequence) add(c Code) {
	if c == cReset {
		*seq = Sequence{}
	}

	for i := len(*seq) - 1; i >= 0; i-- {
		if isSupersededBy((*seq)[i], c) {
			*seq = append((*seq)[:i], (*seq)[i+1:]...)
		}
	}

	if !isStyleOff(c) {
		*seq = append(*seq, c)
	}
}

// ParseSGR parses an ANSI escape sequence  of SGR (Set Graphic Rendition,
// format ESC[code;..;codem) parameters into a slice of ansi.Code.
//
// ParseSGR returns nil if supplied Escape Sequence does not look slike an SGR
// sequence (does not start by '\x1b[' nor ends by 'm').
//
// ParseSGR does not check for correctness of the supplied escape sequence.
func ParseSGR(esc string) Sequence {
	if !isSGR(esc) {
		return nil
	}

	if esc == cCSI+"m" {
		return Sequence{cReset}
	}

	wesc := strings.TrimPrefix(esc, cCSI)
	wesc = strings.TrimSuffix(wesc, "m")

	codes := strings.Split(wesc, ";")

	var s Sequence
	for i := 0; i < len(codes); i++ {
		switch c := codes[i]; c {
		case "":
			s = append(s, cReset)

		case cSetFGColor, cSetBGColor:
			if i == len(codes)-1 {
				break
			}
			switch codes[i+1] {
			case "5":
				if i == len(codes)-2 {
					break
				}
				s, i = append(s, c+";5;"+codes[i+2]), i+2
			case "2":
				if i == len(codes)-4 {
					break
				}
				s, i = append(s, c+";2;"+codes[i+2]+";"+codes[i+3]+";"+codes[i+4]), i+4
			}

		default:
			s = append(s, c)
		}
	}
	return s
}

// Combine merges two ANSI sequences. It deletes existing Codes that are
// superseded by new ones (like a Red by a Green) or as well as Codes that are
// not useful at the boundaries of the SGRSequence (like BoldOff without a
// BoldOn).
func Combine(a, b Sequence) (s Sequence) {
	s = append(s, a...)

	for _, c := range b {
		s.add(c)
	}

	return
}

func isSGR(esc string) bool {
	return strings.HasPrefix(esc, cCSI) && strings.HasSuffix(esc, "m")
}
