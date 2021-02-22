package ansi

// ANSI escapes sequences for common Set Graphic Rendition.
const (
	Reset = "\x1b[0m"

	Bold       = "\x1b[1m"
	Faint      = "\x1b[2m"
	Italic     = "\x1b[3m"
	Underline  = "\x1b[4m"
	SlowBlink  = "\x1b[5m"
	RapidBlink = "\x1b[6m"
	Inverse    = "\x1b[7m"
	Conceal    = "\x1b[8m"
	CrossedOut = "\x1b[9m"

	DoublyUnderlined = "\x1b[21m"

	Normal        = "\x1b[22m"
	ItalicOff     = "\x1b[23m"
	UnderlineOff  = "\x1b[24m"
	BlinkOff      = "\x1b[25m"
	InverseOff    = "\x1b[27m"
	Reveal        = "\x1b[28m"
	NotCrossedOut = "\x1b[29m"

	Black     = "\x1b[30m"
	Red       = "\x1b[31m"
	Green     = "\x1b[32m"
	Yellow    = "\x1b[33m"
	Blue      = "\x1b[34m"
	Magenta   = "\x1b[35m"
	Cyan      = "\x1b[36m"
	White     = "\x1b[37m"
	DefaultFG = "\x1b[39m"

	BlackBG   = "\x1b[40m"
	RedBG     = "\x1b[41m"
	GreenBG   = "\x1b[42m"
	YellowBG  = "\x1b[43m"
	BlueBG    = "\x1b[44m"
	MagentaBG = "\x1b[45m"
	CyanBG    = "\x1b[46m"
	WhiteBG   = "\x1b[47m"
	DefaultBG = "\x1b[49m"

	Framed       = "\x1b[51m"
	Encircled    = "\x1b[52m"
	Overlined    = "\x1b[53m"
	NotFramed    = "\x1b[54m"
	NotOverlined = "\x1b[55m"

	BrightBlack   = "\x1b[90m"
	BrightRed     = "\x1b[91m"
	BrightGreen   = "\x1b[92m"
	BrightYellow  = "\x1b[93m"
	BrightBlue    = "\x1b[94m"
	BrightMagenta = "\x1b[95m"
	BrightCyan    = "\x1b[96m"
	BrightWhite   = "\x1b[97m"

	BrightBlackBG   = "\x1b[100m"
	BrightRedBG     = "\x1b[101m"
	BrightGreenBG   = "\x1b[102m"
	BrightYellowBG  = "\x1b[103m"
	BrightBlueBG    = "\x1b[104m"
	BrightMagentaBG = "\x1b[105m"
	BrightCyanBG    = "\x1b[106m"
	BrightWhiteBG   = "\x1b[107m"
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

// SetBold sets provided string to Bold.
func SetBold(s string) string {
	return Bold + s + Normal
}

// SetFaint decorates provided string with Faint style.
func SetFaint(s string) string {
	return Faint + s + Normal
}

// SetItalic set provided string to Italic.
func SetItalic(s string) string {
	return Italic + s + ItalicOff
}

// SetUnderline underlines provided string.
func SetUnderline(s string) string {
	return Underline + s + UnderlineOff
}

// SetSlowBlink makes provided string to blink slowly.
func SetSlowBlink(s string) string {
	return SlowBlink + s + BlinkOff
}

// SetRapidBlink makes provided string to blink rapidly.
func SetRapidBlink(s string) string {
	return RapidBlink + s + BlinkOff
}

// SetInverse inverts provided string colors.
func SetInverse(s string) string {
	return Inverse + s + InverseOff
}

// SetConceal conceals provided string.
func SetConceal(s string) string {
	return Conceal + s + Reveal
}

// SetCrossedOut crosses out provided string.
func SetCrossedOut(s string) string {
	return CrossedOut + s + NotCrossedOut
}

// SetBlack sets provided string foreground to black.
func SetBlack(s string) string {
	return Black + s + DefaultFG
}

// SetRed sets provided string foreground to red.
func SetRed(s string) string {
	return Red + s + DefaultFG
}

// SetGreen sets provided string foreground to green.
func SetGreen(s string) string {
	return Green + s + DefaultFG
}

// SetYellow sets provided string foreground to yellow.
func SetYellow(s string) string {
	return Yellow + s + DefaultFG
}

// SetBlue sets provided string foreground to blue.
func SetBlue(s string) string {
	return Blue + s + DefaultFG
}

// SetMagenta sets provided string foreground to magenta.
func SetMagenta(s string) string {
	return Magenta + s + DefaultFG
}

// SetCyan sets provided string foreground to cyan.
func SetCyan(s string) string {
	return Cyan + s + DefaultFG
}

// SetWhite sets provided string foreground to white.
func SetWhite(s string) string {
	return White + s + DefaultFG
}

// SetBlackBG sets provided string background to black.
func SetBlackBG(s string) string {
	return BlackBG + s + DefaultBG
}

// SetRedBG sets provided string background to red.
func SetRedBG(s string) string {
	return RedBG + s + DefaultBG
}

// SetGreenBG sets provided string background to green.
func SetGreenBG(s string) string {
	return GreenBG + s + DefaultBG
}

// SetYellowBG sets provided string background to yellow.
func SetYellowBG(s string) string {
	return YellowBG + s + DefaultBG
}

// SetBlueBG sets provided string background to blue.
func SetBlueBG(s string) string {
	return BlueBG + s + DefaultBG
}

// SetMagentaBG sets provided string background to magenta.
func SetMagentaBG(s string) string {
	return MagentaBG + s + DefaultBG
}

// SetCyanBG sets provided string background to cyan.
func SetCyanBG(s string) string {
	return CyanBG + s + DefaultBG
}

// SetWhiteBG sets provided string background to white.
func SetWhiteBG(s string) string {
	return WhiteBG + s + DefaultBG
}

// SetFramed draws a frame around the provided string.
func SetFramed(s string) string {
	return Framed + s + NotFramed
}

// SetEncircled encircles provided string.
func SetEncircled(s string) string {
	return Encircled + s + NotFramed
}

// SetOverlined overlines provided string.
func SetOverlined(s string) string {
	return Overlined + s + NotOverlined
}

// SetBrightBlack sets provided string foreground in bright black.
func SetBrightBlack(s string) string {
	return BrightBlack + s + DefaultFG
}

// SetBrightRed sets provided string foreground in bright red.
func SetBrightRed(s string) string {
	return BrightRed + s + DefaultFG
}

// SetBrightGreen sets provided string foreground in bright green.
func SetBrightGreen(s string) string {
	return BrightGreen + s + DefaultFG
}

// SetBrightYellow sets provided string foreground in bright yellow.
func SetBrightYellow(s string) string {
	return BrightYellow + s + DefaultFG
}

// SetBrightBlue sets provided string foreground in bright blue.
func SetBrightBlue(s string) string {
	return BrightBlue + s + DefaultFG
}

// SetBrightMagenta sets provided string foreground in bright magenta.
func SetBrightMagenta(s string) string {
	return BrightMagenta + s + DefaultFG
}

// SetBrightCyan sets provided string foreground in bright cyan.
func SetBrightCyan(s string) string {
	return BrightCyan + s + DefaultFG
}

// SetBrightWhite sets provided string foreground in bright white.
func SetBrightWhite(s string) string {
	return BrightWhite + s + DefaultFG
}

// SetBrightBlackBG sets provided string background in bright black.
func SetBrightBlackBG(s string) string {
	return BrightBlackBG + s + DefaultBG
}

// SetBrightRedBG sets provided string background in bright red.
func SetBrightRedBG(s string) string {
	return BrightRedBG + s + DefaultBG
}

// SetBrightGreenBG sets provided string background in bright green.
func SetBrightGreenBG(s string) string {
	return BrightGreenBG + s + DefaultBG
}

// SetBrightYellowBG sets provided string background in bright yellow.
func SetBrightYellowBG(s string) string {
	return BrightYellowBG + s + DefaultBG
}

// SetBrightBlueBG sets provided string background in bright blue.
func SetBrightBlueBG(s string) string {
	return BrightBlueBG + s + DefaultBG
}

// SetBrightMagentaBG sets provided string background in bright magenta.
func SetBrightMagentaBG(s string) string {
	return BrightMagentaBG + s + DefaultBG
}

// SetBrightCyanBG sets provided string background in bright cyan.
func SetBrightCyanBG(s string) string {
	return BrightCyanBG + s + DefaultBG
}

// SetBrightWhiteBG sets provided string background in bright white.
func SetBrightWhiteBG(s string) string {
	return BrightWhiteBG + s + DefaultBG
}

// FuncMap provides a text/template FuncMap compatible mapping
// to use 'ansi' functions within templates.
func FuncMap() map[string]interface{} {
	return map[string]interface{}{
		"Bold":            SetBold,
		"Faint":           SetFaint,
		"Italic":          SetItalic,
		"Underline":       SetUnderline,
		"SlowBlink":       SetSlowBlink,
		"RapidBlink":      SetRapidBlink,
		"Inverse":         SetInverse,
		"Conceal":         SetConceal,
		"CrossedOut":      SetCrossedOut,
		"Black":           SetBlack,
		"Red":             SetRed,
		"Green":           SetGreen,
		"Yellow":          SetYellow,
		"Blue":            SetBlue,
		"Magenta":         SetMagenta,
		"Cyan":            SetCyan,
		"White":           SetWhite,
		"BlackBG":         SetBlackBG,
		"RedBG":           SetRedBG,
		"GreenBG":         SetGreenBG,
		"YellowBG":        SetYellowBG,
		"BlueBG":          SetBlueBG,
		"MagentaBG":       SetMagentaBG,
		"CyanBG":          SetCyanBG,
		"WhiteBG":         SetWhiteBG,
		"Framed":          SetFramed,
		"Encircled":       SetEncircled,
		"Overlined":       SetOverlined,
		"BrightBlack":     SetBrightBlack,
		"BrightRed":       SetBrightRed,
		"BrightGreen":     SetBrightGreen,
		"BrightYellow":    SetBrightYellow,
		"BrightBlue":      SetBrightBlue,
		"BrightMagenta":   SetBrightMagenta,
		"BrightCyan":      SetBrightCyan,
		"BrightWhite":     SetBrightWhite,
		"BrightBlackBG":   SetBrightBlackBG,
		"BrightRedBG":     SetBrightRedBG,
		"BrightGreenBG":   SetBrightGreenBG,
		"BrightYellowBG":  SetBrightYellowBG,
		"BrightBlueBG":    SetBrightBlueBG,
		"BrightMagentaBG": SetBrightMagentaBG,
		"BrightCyanBG":    SetBrightCyanBG,
		"BrightWhiteBG":   SetBrightWhiteBG,
	}
}
