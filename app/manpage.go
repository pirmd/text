package app

//manpage poposes some basic template for documentation in manpage-like
//format.

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/pirmd/cli/style"
)

//ManSection identify the man section to which the command is belonging to
//It should be most of the time 1
const ManSection = "1"

//PrintManpage outputs to w the command's documentation in a manpage-like format.
//It does not recurse over sub-commands if any, so that they are not fully documented
//in this page and need to be generated separatly.
func PrintManpage(w io.Writer, c *command, st style.Styler) {
	stMan := st.WithAutostyler(style.LightMarkup.Extend(style.Markup{
		regexp.MustCompile(`\s((?:[^\s]+?)\([1-7]\))\s`):        style.FmtBold,
		regexp.MustCompile(fmt.Sprintf(`\s(%s)\s`, fmtName(c))): style.FmtBold,
	}))

	fmt.Fprintf(w, st.DocHeader(st.Upper("%s %s \"%[1]s\\-%s\"", fmtName(c), ManSection, c.Version)))

	printLongUsage(w, c, stMan)

	var seeAlso []string
	for _, cmd := range c.cmds {
		if len(cmd.cmds) > 0 {
			seeAlso = append(seeAlso, fmtName(cmd)+"("+ManSection+")")
		}
	}
	if len(seeAlso) > 0 {
		fmt.Fprintf(w, stMan.Header("See Also"))
		fmt.Fprintf(w, stMan.Paragraph(strings.Join(seeAlso, ",")))
	}
}

//GenerateManpage generates manpage for the given command.
//It also recures over sub-commands - if any - to generate their
//own manpages.
func GenerateManpage(c *command) error {
	fname := fmtName(c) + "." + ManSection
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, cmd := range c.cmds {
		if len(cmd.cmds) > 0 {
			if err := GenerateManpage(cmd); err != nil {
				return err
			}
		}
	}

	fmt.Printf("Generating manpage for command '%s' to file '%s'\n", c.name, fname)
	PrintManpage(f, c, style.Mandoc)
	return nil
}
