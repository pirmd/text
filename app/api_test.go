package app

import (
	"testing"
)

func TestAPI(t *testing.T) {
	var testResult string

	testApp := New("testApp", "A test for my minimalist cli app building lib")
	testAppFlag := testApp.NewBoolFlag("bool", "Test bool flag")
	testArgString := testApp.NewStringArg("string", "Test String arg", false)
	testApp.Execute = func() error { testResult += *testArgString; return nil }

	testCmd := testApp.NewCommand("test", "Test a new sub-command")
	testCmdArg := testCmd.NewStringArg("test_arg", "Test String arg of sub-command", false)
	testCmd.Execute = func() error { testResult += *testCmdArg; return nil }

	t.Run("Without Subcommand", func(t *testing.T) {
		testResult, *testAppFlag, *testArgString, *testCmdArg = "", false, "", ""
		cmdline := []string{"--bool", "I'm the test string of test app"}

		if err := testApp.Run(cmdline); err != nil {
			t.Fatalf("Run of %#v failed: %v", cmdline, err)
		}

		if *testAppFlag != true {
			t.Errorf("Boolean flag is incorrectly set")
		}

		if testResult != "I'm the test string of test app" {
			t.Errorf("Test sub-command doe snot run as expected (get %s)", testResult)
		}
	})

	t.Run("With Subcommand", func(t *testing.T) {
		testResult, *testAppFlag, *testArgString, *testCmdArg = "", false, "", ""
		cmdline := []string{"--bool", "test", "I'm the test string of test sub-command"}

		if err := testApp.Run(cmdline); err != nil {
			t.Fatalf("Run of %#v failed: %v", cmdline, err)
		}

		if *testAppFlag != true {
			t.Errorf("Boolean flag is incorrectly set")
		}

		if testResult != "I'm the test string of test sub-command" {
			t.Errorf("Test sub-command doe snot run as expected (get %s)", testResult)
		}
	})

	t.Run("Wrong arg number/order/type", func(t *testing.T) {
		testResult, *testAppFlag, *testArgString, *testCmdArg = "", false, "", ""

		cmdline := []string{"--bool"}
		if err := testApp.Run(cmdline); err == nil {
			t.Errorf("Running succeed with a malformed command line (%#v)", cmdline)
		}

		cmdline = []string{"test"}
		if err := testApp.Run(cmdline); err == nil {
			t.Errorf("Running succeed with a malformed command line (%#v)", cmdline)
		}

		cmdline = []string{"test", "--bool", "I'm the test string of test sub-command"}
		if err := testApp.Run(cmdline); err == nil {
			t.Errorf("Running succeed with a malformed command line (%#v)", cmdline)
		}
	})
}
