package ansi

// ANSI escapes sequences for common  Graphic Rendition.
const (
	Reset = "\x1b[0m"

	BoldOn       = "\x1b[1m"
	FaintOn      = "\x1b[2m"
	ItalicOn     = "\x1b[3m"
	UnderlineOn  = "\x1b[4m"
	SlowBlinkOn  = "\x1b[5m"
	RapidBlinkOn = "\x1b[6m"
	InverseOn    = "\x1b[7m"
	ConcealOn    = "\x1b[8m"
	CrossedOutOn = "\x1b[9m"

	DoublyUnderlined = "\x1b[21m"

	Normal        = "\x1b[22m"
	BoldOff       = Normal
	FaintOff      = Normal
	ItalicOff     = "\x1b[23m"
	UnderlineOff  = "\x1b[24m"
	BlinkOff      = "\x1b[25m"
	InverseOff    = "\x1b[27m"
	ConcealOff    = "\x1b[28m"
	CrossedOutOff = "\x1b[29m"

	BlackOn   = "\x1b[30m"
	RedOn     = "\x1b[31m"
	GreenOn   = "\x1b[32m"
	YellowOn  = "\x1b[33m"
	BlueOn    = "\x1b[34m"
	MagentaOn = "\x1b[35m"
	CyanOn    = "\x1b[36m"
	WhiteOn   = "\x1b[37m"
	DefaultFG = "\x1b[39m"

	BlackBGOn   = "\x1b[40m"
	RedBGOn     = "\x1b[41m"
	GreenBGOn   = "\x1b[42m"
	YellowBGOn  = "\x1b[43m"
	BlueBGOn    = "\x1b[44m"
	MagentaBGOn = "\x1b[45m"
	CyanBGOn    = "\x1b[46m"
	WhiteBGOn   = "\x1b[47m"
	DefaultBG   = "\x1b[49m"

	FramedOn     = "\x1b[51m"
	EncircledOn  = "\x1b[52m"
	OverlinedOn  = "\x1b[53m"
	FramedOff    = "\x1b[54m"
	EncircledOff = FramedOff
	OverlinedOff = "\x1b[55m"

	BrightBlackOn   = "\x1b[90m"
	BrightRedOn     = "\x1b[91m"
	BrightGreenOn   = "\x1b[92m"
	BrightYellowOn  = "\x1b[93m"
	BrightBlueOn    = "\x1b[94m"
	BrightMagentaOn = "\x1b[95m"
	BrightCyanOn    = "\x1b[96m"
	BrightWhiteOn   = "\x1b[97m"

	BrightBlackBGOn   = "\x1b[100m"
	BrightRedBGOn     = "\x1b[101m"
	BrightGreenBGOn   = "\x1b[102m"
	BrightYellowBGOn  = "\x1b[103m"
	BrightBlueBGOn    = "\x1b[104m"
	BrightMagentaBGOn = "\x1b[105m"
	BrightCyanBGOn    = "\x1b[106m"
	BrightWhiteBGOn   = "\x1b[107m"
)

// FGColor8bit sets the foreground color in 8bit colors palette (256 colors).
func FGColor8bit(color string) string {
	return "\x1b[38;5;" + color + "m"
}

// BGColor8bit sets the background color in 8bit colors palette (256 colors).
func BGColor8bit(color string) string {
	return "\x1b[48;5;" + color + "m"
}

// FGColor24bit sets the foreground color in 24bit colors palette (R,G,B colors).
func FGColor24bit(r, g, b string) string {
	return "\x1b[38;2;" + r + ";" + g + ";" + b + "m"
}

// BGColor24bit sets the foreground color in 24bit colors palette (R,G,B colors).
func BGColor24bit(r, g, b string) string {
	return "\x1b[38;2;" + r + ";" + g + ";" + b + "m"
}

// Bold sets provided string to Bold.
func Bold(s string) string {
	if s == "" {
		return ""
	}

	return BoldOn + s + BoldOff
}

// Faint decorates provided string with Faint style.
func Faint(s string) string {
	if s == "" {
		return ""
	}

	return FaintOn + s + FaintOff
}

// Italic set provided string to Italic.
func Italic(s string) string {
	if s == "" {
		return ""
	}

	return ItalicOn + s + ItalicOff
}

// Underline underlines provided string.
func Underline(s string) string {
	if s == "" {
		return ""
	}

	return UnderlineOn + s + UnderlineOff
}

// SlowBlink makes provided string to blink slowly.
func SlowBlink(s string) string {
	if s == "" {
		return ""
	}

	return SlowBlinkOn + s + BlinkOff
}

// RapidBlink makes provided string to blink rapidly.
func RapidBlink(s string) string {
	if s == "" {
		return ""
	}

	return RapidBlinkOn + s + BlinkOff
}

// Inverse inverts provided string colors.
func Inverse(s string) string {
	if s == "" {
		return ""
	}

	return InverseOn + s + InverseOff
}

// Conceal conceals provided string.
func Conceal(s string) string {
	if s == "" {
		return ""
	}

	return ConcealOn + s + ConcealOff
}

// CrossedOut crosses out provided string.
func CrossedOut(s string) string {
	if s == "" {
		return ""
	}

	return CrossedOutOn + s + CrossedOutOff
}

// Black sets provided string foreground to black.
func Black(s string) string {
	if s == "" {
		return ""
	}

	return BlackOn + s + DefaultFG
}

// Red sets provided string foreground to red.
func Red(s string) string {
	if s == "" {
		return ""
	}

	return RedOn + s + DefaultFG
}

// Green sets provided string foreground to green.
func Green(s string) string {
	if s == "" {
		return ""
	}

	return GreenOn + s + DefaultFG
}

// Yellow sets provided string foreground to yellow.
func Yellow(s string) string {
	if s == "" {
		return ""
	}

	return YellowOn + s + DefaultFG
}

// Blue sets provided string foreground to blue.
func Blue(s string) string {
	if s == "" {
		return ""
	}

	return BlueOn + s + DefaultFG
}

// Magenta sets provided string foreground to magenta.
func Magenta(s string) string {
	if s == "" {
		return ""
	}

	return MagentaOn + s + DefaultFG
}

// Cyan sets provided string foreground to cyan.
func Cyan(s string) string {
	if s == "" {
		return ""
	}

	return CyanOn + s + DefaultFG
}

// White sets provided string foreground to white.
func White(s string) string {
	if s == "" {
		return ""
	}

	return WhiteOn + s + DefaultFG
}

// BlackBG sets provided string background to black.
func BlackBG(s string) string {
	if s == "" {
		return ""
	}

	return BlackBGOn + s + DefaultBG
}

// RedBG sets provided string background to red.
func RedBG(s string) string {
	if s == "" {
		return ""
	}

	return RedBGOn + s + DefaultBG
}

// GreenBG sets provided string background to green.
func GreenBG(s string) string {
	if s == "" {
		return ""
	}

	return GreenBGOn + s + DefaultBG
}

// YellowBG sets provided string background to yellow.
func YellowBG(s string) string {
	if s == "" {
		return ""
	}

	return YellowBGOn + s + DefaultBG
}

// BlueBG sets provided string background to blue.
func BlueBG(s string) string {
	if s == "" {
		return ""
	}

	return BlueBGOn + s + DefaultBG
}

// MagentaBG sets provided string background to magenta.
func MagentaBG(s string) string {
	if s == "" {
		return ""
	}

	return MagentaBGOn + s + DefaultBG
}

// CyanBG sets provided string background to cyan.
func CyanBG(s string) string {
	if s == "" {
		return ""
	}

	return CyanBGOn + s + DefaultBG
}

// WhiteBG sets provided string background to white.
func WhiteBG(s string) string {
	if s == "" {
		return ""
	}

	return WhiteBGOn + s + DefaultBG
}

// Framed draws a frame around the provided string.
func Framed(s string) string {
	if s == "" {
		return ""
	}

	return FramedOn + s + FramedOff
}

// Encircled encircles provided string.
func Encircled(s string) string {
	if s == "" {
		return ""
	}

	return EncircledOn + s + EncircledOff
}

// Overlined overlines provided string.
func Overlined(s string) string {
	if s == "" {
		return ""
	}

	return OverlinedOn + s + OverlinedOff
}

// BrightBlack sets provided string foreground in bright black.
func BrightBlack(s string) string {
	if s == "" {
		return ""
	}

	return BrightBlackOn + s + DefaultFG
}

// BrightRed sets provided string foreground in bright red.
func BrightRed(s string) string {
	if s == "" {
		return ""
	}

	return BrightRedOn + s + DefaultFG
}

// BrightGreen sets provided string foreground in bright green.
func BrightGreen(s string) string {
	if s == "" {
		return ""
	}

	return BrightGreenOn + s + DefaultFG
}

// BrightYellow sets provided string foreground in bright yellow.
func BrightYellow(s string) string {
	if s == "" {
		return ""
	}

	return BrightYellowOn + s + DefaultFG
}

// BrightBlue sets provided string foreground in bright blue.
func BrightBlue(s string) string {
	if s == "" {
		return ""
	}

	return BrightBlueOn + s + DefaultFG
}

// BrightMagenta sets provided string foreground in bright magenta.
func BrightMagenta(s string) string {
	if s == "" {
		return ""
	}

	return BrightMagentaOn + s + DefaultFG
}

// BrightCyan sets provided string foreground in bright cyan.
func BrightCyan(s string) string {
	if s == "" {
		return ""
	}

	return BrightCyanOn + s + DefaultFG
}

// BrightWhite sets provided string foreground in bright white.
func BrightWhite(s string) string {
	if s == "" {
		return ""
	}

	return BrightWhiteOn + s + DefaultFG
}

// BrightBlackBG sets provided string background in bright black.
func BrightBlackBG(s string) string {
	if s == "" {
		return ""
	}

	return BrightBlackBGOn + s + DefaultBG
}

// BrightRedBG sets provided string background in bright red.
func BrightRedBG(s string) string {
	if s == "" {
		return ""
	}

	return BrightRedBGOn + s + DefaultBG
}

// BrightGreenBG sets provided string background in bright green.
func BrightGreenBG(s string) string {
	if s == "" {
		return ""
	}

	return BrightGreenBGOn + s + DefaultBG
}

// BrightYellowBG sets provided string background in bright yellow.
func BrightYellowBG(s string) string {
	if s == "" {
		return ""
	}

	return BrightYellowBGOn + s + DefaultBG
}

// BrightBlueBG sets provided string background in bright blue.
func BrightBlueBG(s string) string {
	if s == "" {
		return ""
	}

	return BrightBlueBGOn + s + DefaultBG
}

// BrightMagentaBG sets provided string background in bright magenta.
func BrightMagentaBG(s string) string {
	if s == "" {
		return ""
	}

	return BrightMagentaBGOn + s + DefaultBG
}

// BrightCyanBG sets provided string background in bright cyan.
func BrightCyanBG(s string) string {
	if s == "" {
		return ""
	}

	return BrightCyanBGOn + s + DefaultBG
}

// BrightWhiteBG sets provided string background in bright white.
func BrightWhiteBG(s string) string {
	if s == "" {
		return ""
	}

	return BrightWhiteBGOn + s + DefaultBG
}

// FuncMap provides a text/template FuncMap compatible mapping
// to use 'ansi' functions within templates.
func FuncMap() map[string]interface{} {
	return map[string]interface{}{
		"Bold":            Bold,
		"Faint":           Faint,
		"Italic":          Italic,
		"Underline":       Underline,
		"SlowBlink":       SlowBlink,
		"RapidBlink":      RapidBlink,
		"Inverse":         Inverse,
		"Conceal":         Conceal,
		"CrossedOut":      CrossedOut,
		"Black":           Black,
		"Red":             Red,
		"Green":           Green,
		"Yellow":          Yellow,
		"Blue":            Blue,
		"Magenta":         Magenta,
		"Cyan":            Cyan,
		"White":           White,
		"BlackBG":         BlackBG,
		"RedBG":           RedBG,
		"GreenBG":         GreenBG,
		"YellowBG":        YellowBG,
		"BlueBG":          BlueBG,
		"MagentaBG":       MagentaBG,
		"CyanBG":          CyanBG,
		"WhiteBG":         WhiteBG,
		"Framed":          Framed,
		"Encircled":       Encircled,
		"Overlined":       Overlined,
		"BrightBlack":     BrightBlack,
		"BrightRed":       BrightRed,
		"BrightGreen":     BrightGreen,
		"BrightYellow":    BrightYellow,
		"BrightBlue":      BrightBlue,
		"BrightMagenta":   BrightMagenta,
		"BrightCyan":      BrightCyan,
		"BrightWhite":     BrightWhite,
		"BrightBlackBG":   BrightBlackBG,
		"BrightRedBG":     BrightRedBG,
		"BrightGreenBG":   BrightGreenBG,
		"BrightYellowBG":  BrightYellowBG,
		"BrightBlueBG":    BrightBlueBG,
		"BrightMagentaBG": BrightMagentaBG,
		"BrightCyanBG":    BrightCyanBG,
		"BrightWhiteBG":   BrightWhiteBG,
	}
}
