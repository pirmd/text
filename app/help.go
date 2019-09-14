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
	fmt.Fprint(w, st.Paragraph(fmtName(c)+" "+c.Version+" - "+c.Usage))
}

//PrintSimpleUsage outputs to w a command's minimal usage message
func PrintSimpleUsage(w io.Writer, c *Command, st style.Styler) {
	PrintSimpleVersion(w, c, st)

	fmt.Fprint(w, st.Header(1)("Synopsis:"))
	for _, s := range fmtSynopsis(c, st) {
		fmt.Fprint(w, st.Tab()(st.Paragraph(s)))
	}
}

//PrintLongUsage outputs a complete help message similar to a manpage
func PrintLongUsage(w io.Writer, c *Command, st style.Styler) {
	fmt.Fprint(w, st.Header(1)("Name"))
	fmt.Fprint(w, st.Paragraph(fmtName(c)+" - "+c.Usage))

	fmt.Fprint(w, st.Header(1)("Synopsis"))
	fmt.Fprint(w, st.Paragraph(strings.Join(fmtSynopsis(c, st), "\n")))

	fmt.Fprint(w, st.Header(1)("Description"))
	fmt.Fprint(w, st.Paragraph(description(c)))

	if len(c.Flags) > 0 {
		fmt.Fprint(w, st.Header(1)("Options"))

		for _, flag := range c.Flags {
			fmt.Fprint(w, st.Define(fmtFlag(flag, st), flag.Usage))
		}
	}

	if len(c.SubCommands) > 0 {
		fmt.Fprint(w, st.Header(1)("Commands"))

		c.visitSubCommands(func(cmd *Command) {
			fmt.Fprint(w, st.Define(fmtCmd(cmd, st), description(cmd)))
		})
	}

	if len(c.Args) > 0 {
		fmt.Fprint(w, st.Header(1)("Arguments"))

		for _, arg := range c.Args {
			fmt.Fprint(w, st.Define(st.Italic(arg.Name), arg.Usage))
		}
	}
}

//ShowVersion prints to os.Stderr a short information about command's version
func ShowVersion(c *Command) {
	if c.Version == "" {
		c.Version = fmt.Sprintf("%s (build %s)", version, build)
	}

	//TODO: allow customization of style.CurrentStyler
	PrintSimpleVersion(os.Stderr, c, style.CurrentStyler)
}

//ShowUsage prints to os.Stderr a command's minimal usage message
func ShowUsage(c *Command) {
	//TODO: allow customization of style.CurrentStyler
	PrintSimpleUsage(os.Stderr, c, style.CurrentStyler)
}

//ShowHelp prints to os.Stderr a command's detailed help message
func ShowHelp(c *Command) {
	//TODO: allow customization of style.CurrentStyler
	PrintLongUsage(os.Stderr, c, style.CurrentStyler)
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

	PrintLongUsage(f, c, style.NewMarkdown())

	for _, cmd := range c.SubCommands { //no need to use c.visitSubCommands as no help files are expected for 'help' or 'version' anyway
		if len(cmd.SubCommands) > 0 {
			cmd.parents = append(cmd.parents, c)
			if err := GenerateHelpFile(cmd); err != nil {
				return err
			}
		}
	}

	return nil
}

func description(c *Command) string {
	if c.Description == "" {
		return c.Usage
	}
	return c.Description
}

func fmtName(c *Command) string {
	return c.fullname()
}

func fmtFlag(flag *Option, st style.Styler) string {
	switch {
	case flag.isBool():
		return fmt.Sprintf("--%s", st.Bold(flag.Name))
	default:
		return fmt.Sprintf("--%s=%s", st.Bold(flag.Name), st.Italic(st.Upper(flag.Name)))
	}
}

func fmtArg(arg *Option, st style.Styler) string {
	a := st.Italic(arg.Name)

	if arg.isCumulative() {
		a = fmt.Sprintf("%s ...", a)
	}

	if arg.Optional {
		a = fmt.Sprintf("[%s]", a)
	}

	return a
}

func fmtCmd(c *Command, st style.Styler) (s string) {
	if len(c.Flags) > 0 {
		s = fmt.Sprintf("%s [<flags>]", st.Bold(c.Name))
	} else {
		s = st.Bold(c.Name)
	}

	var subcmds []string
	c.visitSubCommands(func(cmd *Command) {
		subcmds = append(subcmds, fmtCmd(cmd, st))
	})
	sc := strings.Join(subcmds, "|")

	if sc != "" {
		if c.Execute != nil {
			sc = fmt.Sprintf("[%s]", sc)
		}
		s = fmt.Sprintf("%s %s", s, sc)
	}

	for _, arg := range c.Args {
		s = fmt.Sprintf("%s %s", s, fmtArg(arg, st))
	}

	return
}

func fmtSynopsis(c *Command, st style.Styler) []string {
	prefix := st.Bold(c.Name)
	for _, flag := range c.Flags {
		prefix = fmt.Sprintf("%s [%s]", prefix, fmtFlag(flag, st))
	}

	var s []string

	if len(c.Args) > 0 {
		a := prefix
		for _, arg := range c.Args {
			a = fmt.Sprintf("%s %s", a, fmtArg(arg, st))
		}
		s = append(s, a)
	}

	if len(c.Args) == 0 && c.Execute != nil {
		s = append(s, prefix)
	}

	if len(c.SubCommands) > 0 {
		c.visitSubCommands(func(cmd *Command) {
			for _, syn := range fmtSynopsis(cmd, st) {
				s = append(s, fmt.Sprintf("%s %s", prefix, syn))
			}
		})
	}

	if len(s) == 0 {
		return []string{prefix}
	}

	return s
}
