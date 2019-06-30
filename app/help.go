package app

//help gathers functions that generate text documentation about a given app.command

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pirmd/cli/style"
)

//PrintSimpleVersion outputs to w a command's minimal usage message
func PrintSimpleVersion(w io.Writer, c *Command, st *style.Styler) {
	fmt.Fprintf(w, st.Line("%s %s - %s", fmtName(c), c.Version, c.Usage))
}

//PrintSimpleUsage outputs to w a command's minimal usage message
func PrintSimpleUsage(w io.Writer, c *Command, st *style.Styler) {
	PrintSimpleVersion(w, c, st)

	fmt.Fprintf(w, st.Header("Synopsis:"))
	for _, s := range fmtSynopsis(c, st) {
		fmt.Fprintf(w, st.Tab(st.Paragraph(s)))
	}
}

//PrintLongUsage outputs a complete help message similar to a manpage
func PrintLongUsage(w io.Writer, c *Command, st *style.Styler) {
	stHelp := st.Extend(style.New(
		style.FormatMap{
			style.FmtHeader:    style.Combine(st.Upper, st.Header),
			style.FmtParagraph: style.Combine(st.Tab, st.Paragraph),
			style.FmtLine:      style.Combine(st.Tab, st.Line),
		},
		nil,
	)).WithAutostyler(style.LightMarkup)

	printLongUsage(w, c, stHelp)
}

//GenerateHelpFile generates a help file in the markdown format for the given
//command. Help file is build after the LongUsage template.
func GenerateHelpFile(c *Command) error {
	fname := fmtName(c) + ".md"
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, cmd := range c.cmds {
		if len(cmd.cmds) > 0 {
			if err := GenerateHelpFile(cmd); err != nil {
				return err
			}
		}
	}

	fmt.Printf("Generating readme for command '%s' to file '%s'\n", c.name, fname)
	PrintLongUsage(f, c, style.Markdown)
	return nil
}

func printLongUsage(w io.Writer, c *Command, st *style.Styler) {
	fmt.Fprintf(w, st.Header("Name"))
	fmt.Fprintf(w, st.Paragraph("%s - %s", fmtName(c), c.Usage))

	fmt.Fprintf(w, st.Header("Synopsis"))
	for i, s := range fmtSynopsis(c, st) {
		if i == 0 {
			fmt.Fprintf(w, st.Paragraph(s))
		} else {
			fmt.Fprintf(w, st.Line(s))
		}
	}

	fmt.Fprintf(w, st.Header("Description"))
	fmt.Fprintf(w, st.Paragraph(description(c)))

	for i, flag := range c.flags {
		if i == 0 {
			fmt.Fprintf(w, st.Header("Options"))
		}
		fmt.Fprintf(w, st.Tab(st.DefTerm(fmtFlag(flag, st)))+st.Tab2(st.DefDesc(flag.Usage)))
	}

	for i, cmd := range c.cmds {
		if i == 0 {
			fmt.Fprintf(w, st.Header("Commands"))
		}
		fmt.Fprintf(w, st.Tab(st.DefTerm(fmtCmd(cmd, st)))+st.Tab2(st.DefDesc(description(cmd))))
	}

	for i, arg := range c.args {
		if i == 0 {
			fmt.Fprintf(w, st.Header("Arguments"))
		}
		fmt.Fprintf(w, st.Tab(st.DefTerm(st.Italic(arg.name)))+st.Tab2(st.DefDesc(arg.Usage)))
	}
}

func description(c *Command) string {
	if c.Description == "" {
		return c.Usage
	}
	return c.Description
}

func fmtName(c *Command) string {
	return strings.Replace(c.fullname, " ", "-", -1)
}

func fmtFlag(flag *option, st *style.Styler) string {
	switch {
	case flag.IsBool():
		return fmt.Sprintf("--%s", st.Bold(flag.name))
	case flag.IsEnum():
		return fmt.Sprintf("--%s=%s", st.Bold(flag.name), st.Italic(strings.Join(flag.value.(*enumValue).options, "|")))
	default:
		return fmt.Sprintf("--%s=%s", st.Bold(flag.name), st.Italic(st.Upper(flag.name)))
	}
}

func fmtArg(arg *option, st *style.Styler) string {
	switch {
	case arg.IsCumulative():
		return fmt.Sprintf("%s ...", st.Italic(arg.name))
	default:
		return st.Italic(arg.name)
	}
}

func fmtCmd(c *Command, st *style.Styler) (s string) {
	if len(c.flags) > 0 {
		s = fmt.Sprintf("%s [<flags>]", st.Bold(c.name))
	} else {
		s = st.Bold(c.name)
	}

	var sc string
	for i, cmd := range c.cmds {
		if i == 0 {
			sc = fmtCmd(cmd, st)
		} else {
			sc = fmt.Sprintf("%s|%s", sc, fmtCmd(cmd, st))
		}
	}
	if sc != "" {
		if c.Execute != nil {
			sc = fmt.Sprintf("[%s]", sc)
		}
		s = fmt.Sprintf("%s %s", s, sc)
	}

	for _, arg := range c.args {
		s = fmt.Sprintf("%s %s", s, fmtArg(arg, st))
	}
	return
}

func fmtSynopsis(c *Command, st *style.Styler) []string {
	prefix := st.Bold(c.name)
	for _, flag := range c.flags {
		prefix = fmt.Sprintf("%s [%s]", prefix, fmtFlag(flag, st))
	}

	var s []string

	if len(c.args) > 0 {
		a := prefix
		for _, arg := range c.args {
			a = fmt.Sprintf("%s %s", a, fmtArg(arg, st))
		}
		s = append(s, a)
	}

	if len(c.args) == 0 && c.Execute != nil {
		s = append(s, prefix)
	}

	if len(c.cmds) > 0 {
		for _, cmd := range c.cmds {
			for _, syn := range fmtSynopsis(cmd, st) {
				s = append(s, fmt.Sprintf("%s %s", prefix, syn))
			}
		}
	}

	if len(s) == 0 {
		return []string{prefix}
	}

	return s
}
