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
func PrintSimpleVersion(w io.Writer, c *Command, st style.Styler) {
	fmt.Fprintf(w, st.Paragraph(fmtName(c)+" "+c.Version+" - "+c.Usage))
}

//PrintSimpleUsage outputs to w a command's minimal usage message
func PrintSimpleUsage(w io.Writer, c *Command, st style.Styler) {
	PrintSimpleVersion(w, c, st)

	fmt.Fprintf(w, st.Header(1)("Synopsis:"))
	for _, s := range fmtSynopsis(c, st) {
		fmt.Fprintf(w, st.Tab()(st.Paragraph(s)))
	}
}

//PrintLongUsage outputs a complete help message similar to a manpage
func PrintLongUsage(w io.Writer, c *Command, st style.Styler) {
	printLongUsage(w, c, st)
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
	PrintLongUsage(f, c, style.NewMarkdown())
	return nil
}

func printLongUsage(w io.Writer, c *Command, st style.Styler) {
	fmt.Fprintf(w, st.Header(1)("Name"))
	fmt.Fprintf(w, st.Paragraph(fmtName(c)+" - "+c.Usage))

	fmt.Fprintf(w, st.Header(1)("Synopsis"))
	fmt.Fprintf(w, st.Paragraph(strings.Join(fmtSynopsis(c, st), "\n")))

	fmt.Fprintf(w, st.Header(1)("Description"))
	fmt.Fprintf(w, st.Paragraph(description(c)))

	for i, flag := range c.flags {
		if i == 0 {
			fmt.Fprintf(w, st.Header(1)("Options"))
		}
		fmt.Fprintf(w, st.Define(fmtFlag(flag, st), flag.Usage))
	}

	for i, cmd := range c.cmds {
		if i == 0 {
			fmt.Fprintf(w, st.Header(1)("Commands"))
		}
		fmt.Fprintf(w, st.Define(fmtCmd(cmd, st), description(cmd)))
	}

	for i, arg := range c.args {
		if i == 0 {
			fmt.Fprintf(w, st.Header(1)("Arguments"))
		}
		fmt.Fprintf(w, st.Define(st.Italic(arg.name), arg.Usage))
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

func fmtFlag(flag *option, st style.Styler) string {
	switch {
	case flag.IsBool():
		return fmt.Sprintf("--%s", st.Bold(flag.name))
	case flag.IsEnum():
		return fmt.Sprintf("--%s=%s", st.Bold(flag.name), st.Italic(strings.Join(flag.value.(*enumValue).options, "|")))
	default:
		return fmt.Sprintf("--%s=%s", st.Bold(flag.name), st.Italic(st.Upper(flag.name)))
	}
}

func fmtArg(arg *option, st style.Styler) string {
	switch {
	case arg.IsCumulative():
		return fmt.Sprintf("%s ...", st.Italic(arg.name))
	default:
		return st.Italic(arg.name)
	}
}

func fmtCmd(c *Command, st style.Styler) (s string) {
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

	var sa string
	for i, arg := range c.args {
		if i == 0 {
			sa = fmtArg(arg, st)
		} else {
			sa = fmt.Sprintf("%s %s", sa, fmtArg(arg, st))
		}
	}
	if sa != "" {
		if c.CanRunWithoutArg {
			s = fmt.Sprintf("%s [%s]", s, sa)
		} else {
			s = fmt.Sprintf("%s %s", s, sa)
		}
	}

	return
}

func fmtSynopsis(c *Command, st style.Styler) []string {
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
