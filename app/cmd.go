// Package app  provides simple support to build a command line application
// that supports nested commands, flags and arguments parsing.  It does not aim
// at be a feature-complete command line app builder but something simpler and
// hopefully lean for simple use.
package app

import (
	"fmt"
	"os"
	"strings"
)

var (
	version = "v?.?.?"  //should be set-up at compile-time through ldflags -X github.com/pirmd/cli/app.version
	build   = "unknown" //should be set-up at compile-time through ldflags -X github.com/pirmd/cli/app.build
)

// Command represents an application. An application is a tree of commands
// that can be nested. Each command is made of a set of flags, a set of args
// and a set of sub-commands. An app is the 'root' command of this tree.
type Command struct {
	//Name is the command's name.
	Name string

	//Usage is a short explanation of what the command does.
	Usage string

	//Description is a long description of the command.
	Description string

	//Version contains command's version. It defaults to $VERSION ($BUILD)
	//where $VERSION and $BUILD are set-up at compile time using ldflags
	//directive (e.g. taking values from git describe). Provided 'go' script
	//gives an example.
	Version string

	//Flags contains the set of command's flags
	Flags Options

	//Args contains the set of command's arguments
	Args Options

	//SubCommands contains the set of command's sub-commands
	SubCommands Commands

	//Execute is the function called to run the command
	Execute func() error

	//ShowHelp is a function that displays help about the command. It will be
	//associated with the "help" sub-command if not nil.  you can also directly
	//alter this behavious by declaring directly any further SubCommands or
	//Flags that manage this situation.
	ShowHelp func(c *Command)

	//ShowVersion is a function that displays version information about the
	//command. It will be associated with the "version" sub-command if not nil.
	//you can also directly alter this behavious by declaring directly any
	//further SubCommands or Flags that manage this situation.
	ShowVersion func(c *Command)

	parents []*Command //list of parent commands, build dynamiccaly during processing of the commands
	cmdline []string
}

// Run executes the command after having parsed the command line
func (c *Command) Run(args []string) error {
	c.cmdline = args

	if err := c.parseFlags(); err != nil {
		return err
	}

	if len(c.cmdline) > 0 {
		if subcmd := c.getSubCommand(c.cmdline[0]); subcmd != nil {
			subcmd.parents = append(subcmd.parents, c)
			return subcmd.Run(c.cmdline[1:])
		}
	}

	if c.Execute == nil {
		return fmt.Errorf("bad command or invalid number of arguments")
	}

	if err := c.parseArgs(); err != nil {
		return err
	}

	if err := c.Execute(); err != nil {
		return err
	}

	return nil
}

// MustRun executes the command after having parsed the command line
// In case of error, it prints the error to os.Stderr and exit with code 1.
func (c *Command) MustRun(args []string) {
	if err := c.Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func (c *Command) fullname() string {
	if len(c.parents) == 0 {
		return c.Name
	}

	return c.parents[len(c.parents)-1].fullname() + "-" + c.Name
}

//visitSubCommands iterates over all command's SubCommands including automatic
//'version' or 'help' sub-commands that where not manually added to command's
//sub-commands list.
func (c *Command) visitSubCommands(fn func(*Command)) {
	if h := c.autoHelpSubCommand(); h != nil {
		fn(h)
	}

	if v := c.autoVersionSubCommand(); v != nil {
		fn(v)
	}

	for _, cmd := range c.SubCommands {
		fn(cmd)
	}
}

func (c *Command) getSubCommand(name string) *Command {
	if h := c.autoHelpSubCommand(); h != nil && h.Name == name {
		return h
	}

	if v := c.autoVersionSubCommand(); v != nil && v.Name == name {
		return v
	}

	return c.SubCommands.get(name)
}

func (c *Command) autoHelpSubCommand() *Command {
	if c.SubCommands.get("help") != nil ||
		c.Flags.get("help") != nil ||
		c.ShowHelp == nil {
		return nil
	}

	return &Command{
		Name:  "help",
		Usage: "Show usage information.",
		Execute: func() error {
			c.ShowHelp(c)
			return nil
		},
	}
}

func (c *Command) autoVersionSubCommand() *Command {
	if c.SubCommands.get("version") != nil ||
		c.Flags.get("version") != nil ||
		c.ShowVersion == nil {
		return nil
	}

	return &Command{
		Name:  "version",
		Usage: "Show version information.",
		Execute: func() error {
			c.ShowVersion(c)
			return nil
		},
	}
}

func (c *Command) parseFlags() error {
	for len(c.cmdline) > 0 {
		if c.cmdline[0] == "--" {
			c.cmdline = c.cmdline[1:]
			return nil
		}
		if !strings.HasPrefix(c.cmdline[0], "--") {
			return nil
		}

		split := strings.Split(c.cmdline[0], "=")
		flag := c.Flags.get(strings.TrimPrefix(split[0], "--"))
		if flag == nil {
			return fmt.Errorf("invalid flag %q (in %q)", split[0], c.cmdline[0])
		}
		switch len(split) {
		case 1:
			if !flag.isBool() {
				return fmt.Errorf("invalid boolean flag %q", split[0])
			}
			if err := flag.value().Set("true"); err != nil {
				return fmt.Errorf("invalid boolean flag %q", split[0])
			}
		case 2:
			if err := flag.value().Set(split[1]); err != nil {
				return fmt.Errorf("invalid value %q for flag %q", split[1], split[0])
			}
		default:
			return fmt.Errorf("invalid flag assignment in %q: too many '='", c.cmdline[0])
		}
		c.cmdline = c.cmdline[1:]
	}
	return nil
}

func (c *Command) parseArgs() error {
	for i, a := range c.Args {
		if i >= len(c.cmdline) {
			if a.Optional {
				if i != len(c.Args)-1 {
					panic("arguments " + a.Name + " is optional but is not the last argument of the command " + c.Name)
				}

				break
			}

			return fmt.Errorf("bad command or invalid number of arguments")
		}

		if err := a.value().Set(c.cmdline[i]); err != nil {
			return fmt.Errorf("invalid value %q for argument %s: %v", c.cmdline[i], a.Name, err)
		}

		if a.isCumulative() {
			if i != len(c.Args)-1 {
				panic("arguments " + a.Name + " is cumulative but is not the last argument of the command " + c.Name)
			}

			for j := i + 1; j < len(c.cmdline); j++ {
				if err := a.value().Set(c.cmdline[j]); err != nil {
					return fmt.Errorf("invalid value %q for argument %s: %v", c.cmdline[j], a.Name, err)
				}
			}
		}
	}

	return nil
}

// Commands represents a set of commands and sub-commands
type Commands []*Command

func (c *Commands) get(name string) *Command {
	for _, cmd := range *c {
		if cmd.Name == name {
			return cmd
		}
	}
	return nil
}

func (c *Commands) append(newcmd *Command) {
	*c = append(*c, newcmd)
}
