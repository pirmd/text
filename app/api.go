package app

//api module file gathers helpers functions and method to declare in a verbose
//manner commands with their flags, args and sub-commands.

//New creates a new command line application.
//New add automatically help and version commands.
func New(name, usage string) *Command {
	return &Command{
		Name:        name,
		Usage:       usage,
		ShowHelp:    ShowUsage,
		ShowVersion: ShowVersion,
	}
}

//NewCommand creates a new sub-command and adds it to the command's
//sub-commands pool.
//It automatically creates or complete the help sub-command.
func (c *Command) NewCommand(name, usage string) *Command {
	if c.ShowHelp == nil && name != "help" {
		c.ShowHelp = ShowUsage
	}

	newcmd := &Command{Name: name, Usage: usage}
	c.SubCommands.append(newcmd)
	return newcmd
}

//NewFlagToVar creates a new flag that is linked to the given var.
//Flag's type (int64, string, strings, bool) is guessed after the linked
//variable type.
func (c *Command) NewFlagToVar(p interface{}, name, usage string) {
	c.Flags.append(&Option{
		Name:  name,
		Usage: usage,
		Var:   p,
	})
}

//NewArgToVar creates a new arg that is linked to the given var.
//Arg's type (int64, string, strings, bool) is guessed after the linked
//variable type.
func (c *Command) NewArgToVar(p interface{}, name, usage string, optional bool) {
	c.Args.append(&Option{
		Name:     name,
		Usage:    usage,
		Var:      p,
		Optional: optional,
	})
}

//NewBoolFlag creates a new bolean flag
func (c *Command) NewBoolFlag(name, usage string) *bool {
	p := new(bool)
	c.NewFlagToVar(p, name, usage)
	return p
}

//NewBoolFlagToVar creates a new bolean flag that is linked to the given var
func (c *Command) NewBoolFlagToVar(p *bool, name, usage string) {
	c.NewFlagToVar(p, name, usage)
}

//NewStringFlag creates a new string flag
func (c *Command) NewStringFlag(name, usage string) *string {
	p := new(string)
	c.NewFlagToVar(p, name, usage)
	return p
}

//NewStringFlagToVar creates a new string flag that is linked to the given var
func (c *Command) NewStringFlagToVar(p *string, name, usage string) {
	c.NewFlagToVar(p, name, usage)
}

//NewStringArg creates a new string arg
func (c *Command) NewStringArg(name, usage string, optional bool) *string {
	p := new(string)
	c.NewArgToVar(p, name, usage, optional)
	return p
}

//NewStringArgToVar creates a new string arg that is linked to the given var
func (c *Command) NewStringArgToVar(p *string, name, usage string, optional bool) {
	c.NewArgToVar(p, name, usage, optional)
}

//NewStringsArg creates a new strings (slice of strings) arg.
//This arg is cumulative in that it will consume all the remaining command line
//arguments to feed a slice of strings. It shall be the last argument of the
//command otherwise command line parsing wil panic
func (c *Command) NewStringsArg(name, usage string, optional bool) *[]string {
	p := new([]string)
	c.NewArgToVar(p, name, usage, optional)
	return p
}

//NewStringsArgToVar creates a new strings (slice of strings) arg that is
//linked to the given var. This arg is cumulative in that it will consume all
//the remaining command line arguments to feed a slice of strings. It shall be
//the last argument of the command otherwise command line parsing wil panic
func (c *Command) NewStringsArgToVar(p *[]string, name, usage string, optional bool) {
	c.NewArgToVar(p, name, usage, optional)
}

//NewInt64Arg creates a new int64 arg
func (c *Command) NewInt64Arg(name, usage string, optional bool) *int64 {
	p := new(int64)
	c.NewArgToVar(p, name, usage, optional)
	return p
}

//NewInt64ArgToVar creates a new int64 arg that is linked to the given var
func (c *Command) NewInt64ArgToVar(p *int64, name, usage string, optional bool) {
	c.NewArgToVar(p, name, usage, optional)
}

//UseConfig creates a new configuration
func (c *Command) UseConfig(cfg interface{}, unmarshaller func([]byte, interface{}) error, files []ConfigFile) {
	c.Config = &Config{
		Var:          cfg,
		Unmarshaller: unmarshaller,
		Files:        files,
	}
}
