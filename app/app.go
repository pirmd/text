// This module provides simple support to build a command line application that supports
// nested commands, flags and arguments parsing.
// It aims not ot be a feature-complete command line app builder (spf13/cobra is way better)
// but something simpler and hopefully lean for simple use.
package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/pirmd/cli/style"
)

var (
	version = "v?.?.?"    //should be set-up at compile-time through ldflags -X github.com/pirmd/cli/app.version
	build   = "unknown"   //should be set-up at compile-time through ldflags -X github.com/pirmd/cli/app.build
)

type command struct {
	//Usage is a short explanation of what the command does
	Usage string

	//Description is a long description of the command
	Description string

    //Version contains command's version. It defaults to $VERSION ($BUILD)
    //where $VERSION and $BUILD are set-up at compile time using ldflags
    //directive (e.g. taking values from git describe). Provided 'go' script
    //gives an example
	Version string

	//Execute is the function called to run the command
	Execute func() error

	name     string
	fullname string //fullname is the complete 'path' to the command (space-separated parent's commands names)
	flags    options
	args     options
	cmds     commands
	cmdline  []string
}

func newCommand(name, usage, parent string) *command {
	if parent == "" {
        return &command{
            name:     name,
            fullname: name,
            Version:  fmt.Sprintf("%s (build %s)", version, build),
            Usage:    usage,
        }
	}

	return &command{
        name:     name,
        fullname: parent + " " + name,
        Version:  fmt.Sprintf("%s (build %s)", version, build),
        Usage:    usage,
    }
}

//New creates a new command line application
//An application is a tree of commands that can be nested. Each command is made of
//a set of flags, a set of args and a set of sub-commands. Some convenient helpers are
//provided to parse the command line or format help information.
//An app is the 'root' command of this tree. An help and version commands are automatically created
//Other commands - with their own sub-commands, flags and args -, flags and options can then be added.
func New(name, usage string) *command {
	a := newCommand(name, usage, "")

	help := a.NewCommand("help", "Show usage information.")
	help.Execute = a.showHelp

	version := a.NewCommand("version", "Show version information.")
	version.Execute = a.showVersion

	a.cmdline = os.Args[1:]
	return a
}

//NewCommand creates a new sub-command and add it to the command's
//sub-commands pool
//
//It automatically creates or complete the help sub-command.
func (c *command) NewCommand(name, usage string) *command {
	if h := c.cmds.Get("help"); h == nil && name != "help" {
		help := c.NewCommand("help", "Show usage information.")
		help.Execute = c.showHelp
	}

	return c.cmds.Add(name, usage, c.fullname)
}

//NewBoolFlag creates a new bolean flag
func (c *command) NewBoolFlag(name, usage string) *bool {
	return c.flags.Bool(name, usage)
}

//NewBoolFlagToVar creates a new bolean flag that is linked to the given var
func (c *command) NewBoolFlagToVar(p *bool, name, usage string) {
	c.flags.BoolVar(p, name, usage)
}

//NewStringFlag creates a new string flag
func (c *command) NewStringFlag(name, usage string) *string {
	return c.flags.String(name, usage)
}

//NewStringFagToVar creates a new string flag that is linked to the given var
func (c *command) NewStringFlagToVar(p *string, name, usage string) {
	c.flags.StringVar(p, name, usage)
}

//NewEnumFlag creates a new enum flag that only accept a specifed list of values
func (c *command) NewEnumFlag(name, usage string, values []string) *string {
	return c.flags.Enum(name, usage, values)
}

//NewEnumFlagToVar creates a new enum flag that only accept a specifed list of values
//and linked it to the given var
func (c *command) NewEnumFlagToVar(p *string, name, usage string, values []string) {
	c.flags.EnumVar(p, name, usage, values)
}

//NewStringArg creates a new string arg
func (c *command) NewStringArg(name, usage string) *string {
	return c.args.String(name, usage)
}

//NewStringArgToVar creates a new string arg that is linked to the given var
func (c *command) NewStringArgToVar(p *string, name, usage string) {
	c.args.StringVar(p, name, usage)
}

//NewStringArg creates a new strings (slice of strings) arg
//This arg is cumulative in that it will consume all the remaining command line arguments to feed
//a slice of strings. It shall be the last argument of teh command otherwise command line parsing
//wil panic
func (c *command) NewStringsArg(name, usage string) *[]string {
	return c.args.Strings(name, usage)
}

//NewStringsArgToVar creates a new strings (slice of strings) arg that is linked to the given var
//This arg is cumulative in that it will consume all the remaining command line arguments to feed
//a slice of strings. It shall be the last argument of teh command otherwise command line parsing
//wil panic
func (c *command) NewStringsArgToVar(p *[]string, name, usage string) {
	c.args.StringsVar(p, name, usage)
}

//NewInt64Arg creates a new int64 arg
func (c *command) NewInt64Arg(name, usage string) *int64 {
	return c.args.Int64(name, usage)
}

//NewInt64ArgToVar creates a new int64 arg that is linked to the given var
func (c *command) NewInt64ArgToVar(p *int64, name, usage string) {
	c.args.Int64Var(p, name, usage)
}

//MustRun executes the command after having parsed the command line
//In case of error, it print the error to stderr and exit with code 1.
func (c *command) MustRun() {
	if err := c.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

//Run executes the command after having parsed the command line
func (c *command) Run() error {
	if err := c.parseFlags(); err != nil {
		return err
	}

	if len(c.cmdline) > 0 {
		if subcmd := c.cmds.Get(c.cmdline[0]); subcmd != nil {
			subcmd.cmdline = c.cmdline[1:]
			return subcmd.Run()
		}
	}

	if c.Execute == nil {
		return fmt.Errorf("Bad command or invalid number of arguments")
	}

	if err := c.parseArgs(); err != nil {
		return err
	}

	if err := c.Execute(); err != nil {
		return err
	}

	return nil
}

func (c *command) parseFlags() error {
	for len(c.cmdline) > 0 {
		if c.cmdline[0] == "--" {
			c.cmdline = c.cmdline[1:]
			return nil
		}
		if !strings.HasPrefix(c.cmdline[0], "--") {
			return nil
		}

		split := strings.Split(c.cmdline[0], "=")
		flag := c.flags.Get(strings.TrimPrefix(split[0], "--"))
		if flag == nil {
			return fmt.Errorf("Invalid flag %q (in %q)", split[0], c.cmdline[0])
		}
		switch len(split) {
		case 1:
			if !flag.IsBool() {
				return fmt.Errorf("Invalid boolean flag %q", split[0])
			}
			if err := flag.Set("true"); err != nil {
				return fmt.Errorf("Invalid boolean flag %q", split[0])
			}
		case 2:
			if err := flag.Set(split[1]); err != nil {
				return fmt.Errorf("Invalid value %q for flag %q", split[1], split[0])
			}
		default:
			return fmt.Errorf("Invalid flag assignment in %q: too many '='", c.cmdline[0])
		}
		c.cmdline = c.cmdline[1:]
	}
	return nil
}

func (c *command) parseArgs() error {
	if len(c.args) > len(c.cmdline) {
		return fmt.Errorf("Bad command or invalid number of arguments")
	}

	for i, a := range c.args {
		if err := a.Set(c.cmdline[i]); err != nil {
			return fmt.Errorf("Invalid value %q for argument %s: %v", c.cmdline[i], a.name, err)
		}

		if a.IsCumulative() {
			if i != len(c.args)-1 {
				panic("arguments " + a.name + " is cumulative but is not the last argument of the command " + c.name)
			}

			for j := i + 1; j < len(c.cmdline); j++ {
				if err := a.Set(c.cmdline[j]); err != nil {
					return fmt.Errorf("Invalid value %q for argument %s: %v", c.cmdline[j], a.name, err)
				}
			}
		}
	}

	return nil
}

//showHelp outputs help message
func (c *command) showHelp() error {
	PrintSimpleUsage(os.Stderr, c, style.CurrentStyler)
	return nil
}

//showVersion outputs version information
func (c *command) showVersion() error {
	PrintSimpleVersion(os.Stderr, c, style.CurrentStyler)
	return nil
}

type commands []*command

func (c *commands) Get(name string) *command {
	for _, cmd := range *c {
		if cmd.name == name {
			return cmd
		}
	}
	return nil
}

func (c *commands) Add(name, usage, parent string) *command {
	if cmd := c.Get(name); cmd != nil {
		panic(fmt.Sprintf("command '%s' cannot be added twice", name))
	}

	newcmd := newCommand(name, usage, parent)
	*c = append(*c, newcmd)

	return newcmd
}

type option struct {
	value

	//Usage is a short description of what the option does
	Usage string

	name string
}

func newOption(v value, name string, usage string) *option {
	return &option{value: v, name: name, Usage: usage}
}

func (o *option) IsBool() bool {
	_, ok := o.value.(*boolValue)
	return ok
}

func (o *option) IsCumulative() bool {
	_, ok := o.value.(*stringsValue)
	return ok
}

func (o *option) IsEnum() bool {
	_, ok := o.value.(*enumValue)
	return ok
}

type options []*option

func (o *options) Get(name string) *option {
	for _, opt := range *o {
		if opt.name == name {
			return opt
		}
	}
	return nil
}

func (o *options) add(v value, name string, usage string) {
	if opt := o.Get(name); opt != nil {
		panic(fmt.Sprintf("option '%s' cannot be added twice", name))
	}

	*o = append(*o, newOption(v, name, usage))
}

func (o *options) BoolVar(p *bool, name string, usage string) {
	o.add(newBoolValue(p), name, usage)
}

func (o *options) Bool(name string, usage string) *bool {
	p := new(bool)
	o.BoolVar(p, name, usage)
	return p
}

func (o *options) Int64Var(p *int64, name string, usage string) {
	o.add(newInt64Value(p), name, usage)
}

func (o *options) Int64(name string, usage string) *int64 {
	p := new(int64)
	o.Int64Var(p, name, usage)
	return p
}

func (o *options) StringVar(p *string, name string, usage string) {
	o.add(newStringValue(p), name, usage)
}

func (o *options) String(name string, usage string) *string {
	p := new(string)
	o.StringVar(p, name, usage)
	return p
}

func (o *options) StringsVar(p *[]string, name string, usage string) {
	o.add(newStringsValue(p), name, usage)
}

func (o *options) Strings(name string, usage string) *[]string {
	p := new([]string)
	o.StringsVar(p, name, usage)
	return p
}

func (o *options) EnumVar(p *string, name string, usage string, values []string) {
	o.add(newEnumValue(p, values...), name, usage)
}

func (o *options) Enum(name string, usage string, values []string) *string {
	p := new(string)
	o.EnumVar(p, name, usage, values)
	return p
}
