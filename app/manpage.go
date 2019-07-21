package app

//manpage poposes some basic template for documentation in manpage-like
//format.

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/pirmd/cli/style"
)

//ManSection identify the man section to which the command is belonging to It
//should be most of the time 1
const ManSection = "1"

//manDate contains the manpage creation date. It is manage through a global
//variable to be tweaked during test stage.
var manDate = time.Now().Format("2006-01-02")

//PrintManpage outputs to w the command's documentation in a manpage-like
//format.  It does not recurse over sub-commands if any, so that they are not
//fully documented in this page and need to be generated separatly.
func PrintManpage(w io.Writer, c *Command, st style.Styler) {
	fmt.Fprintf(w, st.Metadata(map[string]string{
		"title":      fmtName(c),
		"mansection": ManSection,
		"date":       manDate,
	}))

	printLongUsage(w, c, st)

	var seeAlso []string
	for _, cmd := range c.cmds {
		if len(cmd.cmds) > 0 {
			seeAlso = append(seeAlso, fmtName(cmd)+"("+ManSection+")")
		}
	}
	if len(seeAlso) > 0 {
		fmt.Fprintf(w, st.Header(1)("See Also"))
		fmt.Fprintf(w, st.Paragraph(strings.Join(seeAlso, ",")))
	}
}

//GenerateManpage generates manpage for the given command.  It also recures
//over sub-commands - if any - to generate their own manpages.
func GenerateManpage(c *Command) error {
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
	PrintManpage(f, c, style.NewMan())
	return nil
}
