package style

import (
	"regexp"
)

var (
	//LightMarkup provides a limited syntax to format a text:
	//*Bold*, _Italic_, ~Strikethrough~, `Code`
	LightMarkup = Markup{
		regexp.MustCompile(`(?:[^\*]|^)(\*([^\*]+?)\*)(?:[^\*]|$)`): FmtBold,   // *Bold*
		regexp.MustCompile(`(?:[^_]|^)(_([^_]+?)_)(?:[^_]|$)`):      FmtItalic, // _Italic_
		regexp.MustCompile(`(?:[^~]|^)(~([^~]+?)~)(?:[^~]|$)`):      FmtStrike, // ~Strikethrough~
		regexp.MustCompile("(`([^`]+?)`)"):                          FmtCode,   // `Code`
	}
)

//Markup is a lightweight markup-like language to apply a style to a string.
//It works by defining a set of rules governing the use of a style format.
//Each rule is a regexp that identifies the pattern indicating part of texts to
//be styled.  Each pattern is linked through a map to the corresponding style
//to apply.
//
//Each regexp can feature up to two capturing groups: any non-captured char are
//kept as is in the final text, first submatch is deleted, the 2d submatch is
//passed to the style function whose result is incorporated in the final text.
//
//The implementation of Markup is really naive and simple, order of regexp
//processing is not guarantied so that trying to achieve complex markup
//rendering is hazardous.
type Markup map[*regexp.Regexp]Format

//Extend allows to complete or override a Markup
func (m Markup) Extend(madd Markup) Markup {
	mext := Markup{}
	for r, fn := range m {
		mext[r] = fn
	}
	for r, fn := range madd {
		mext[r] = fn
	}
	return mext
}

//Render returns a FormatFn that applies the specified markup with the given
//Styler
func (m Markup) Render(st *Styler) FormatFn {
	return func(s string) string {
		return m.render(st, s)
	}
}

func (m Markup) render(st *Styler, s string) string {
	for r, fmt := range m {
		if matches := r.FindAllStringSubmatchIndex(s, -1); matches != nil {
			var t string
			var lastpos int

			//each match is
			//match[0]-match[1]: index of complete regexp match, these are kept as-is in final text
			//match[2]-match[3]: index of 1st submatch to be delete
			//match[4]-match[5]: index of 2d submatch to be send to styling function
			for _, match := range matches {
				switch len(match) {
				case 2:
					match = append(match, match[0:2]...)
					fallthrough
				case 4:
					match = append(match, match[2:4]...)
				}

				t = t + s[lastpos:match[2]]
				t = t + st.get(fmt)(s[match[4]:match[5]])
				lastpos = match[3]
			}

			s = t + s[lastpos:]
		}
	}

	return s
}

//WithAutostyler creates a new Styler that automatically styles text using the
//given markup
func (st *Styler) WithAutostyler(m Markup) *Styler {
	return st.Extend(New(FormatMap{FmtAuto: m.Render(st)}, nil))
}
