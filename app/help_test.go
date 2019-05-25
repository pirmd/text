package app

import (
	"bytes"
	"testing"

	"github.com/pirmd/cli/style"
	"github.com/pirmd/verify"
)

func buildTestApp() *command {
	testApp := New("cli.test", "A test for my minimalist cli app building lib")

	testApp.Version = "3.14159"
	testApp.Description = `cli.test is mainly here to test and demonstrate the set-up of a commandline application using the small *cli* library.

Like many other test involving text issuance, using the famous _lorem ipsum_ pattern is a must.

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

	_ = testApp.NewBoolFlag("bool", "A boolean flag")
	_ = testApp.NewStringFlag("flag", "A string flag")
	_ = testApp.NewEnumFlag("enum", "An enum flag", []string{"good", "bad", "ugly"})
	_ = testApp.NewStringArg("string", "An argument that should be a string")

	testCmd := testApp.NewCommand("test", "Demonstrate a sub-command")
	testCmd.Description = `Many of you know about _lorem ipsum_ but few know what it means. So this sub-command's description is a perfect occasion to clarify that, while it was popularized in the XVI century, it originates from a text from Ciceron.`
	_ = testCmd.NewStringArg("test_arg", "Test String arg of sub-command")

	testCmd2 := testApp.NewCommand("test2", "Test a second sub-command")
	_ = testCmd2.NewStringArg("test2_arg1", "Test String arg of sub-command")
	_ = testCmd2.NewStringArg("test2_arg2", "Test String arg of sub-command")

	testCmd3 := testApp.NewCommand("test3", "Test another sub-command that has a sub-subcommand *Test31*")
	_ = testCmd3.NewStringFlag("stringflag", "Test31 String flag of sub-command")
	testCmd31 := testCmd3.NewCommand("test31", "Test a sub-sub-command")
	_ = testCmd31.NewInt64Arg("test31_arg", "Test Int64 arg of a sub-sub-command")

	testCmd4 := testApp.NewCommand("test4", "Test another sub-command with unlimited args")
	_ = testCmd4.NewStringsArg("test4_arg", "Test an arg with unlimited number of strings")

	return testApp
}

func TestHelp(t *testing.T) {
	testApp := buildTestApp()

	out := new(bytes.Buffer)
	PrintLongUsage(out, testApp, style.ColorTerm)
	verify.MatchGolden(t, out.String(), "Help message is incorrectly formatted")
}
