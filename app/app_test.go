package app

import (
	"reflect"
	"testing"
)

func TestCmdlineFlags(t *testing.T) {
	testApp := New("testApp", "A test for my minimalist cli app building lib")
	testFlagBool := testApp.NewBoolFlag("bool", "Test Boolean flag")
	testFlagString := testApp.NewStringFlag("string", "Test String flag")

	t.Run("Simple flag parsing", func(t *testing.T) {
		*testFlagBool, *testFlagString = false, ""
		testApp.cmdline = []string{"--bool", "--string=testing is fun"}
		if err := testApp.parseFlags(); err != nil {
			t.Fatalf("Parsing failed: %v", err)
		}

		if *testFlagString != "testing is fun" {
			t.Errorf("String flag is not recognized (string is %s)", *testFlagString)
		}

		if !*testFlagBool {
			t.Errorf("Boolean flag is not recognized")
		}
	})

	t.Run("unknown flag parsing", func(t *testing.T) {
		*testFlagBool, *testFlagString = false, ""
		testApp.cmdline = []string{"--unknown"}
		if err := testApp.parseFlags(); err == nil {
			t.Errorf("Parsing succeed dispite malformed command line")
		}

		if *testFlagString != "" {
			t.Errorf("String flag was modified but should not (string is %s)", *testFlagString)
		}

		if *testFlagBool {
			t.Errorf("Boolean flag was modified but should not")
		}
	})

	t.Run("Test end of flags", func(t *testing.T) {
		*testFlagBool, *testFlagString = false, ""
		testApp.cmdline = []string{"--", "--unknown"}
		if err := testApp.parseFlags(); err != nil {
			t.Errorf("Parsing failed, %v", err)
		}
	})
}

func TestCmdlineArg(t *testing.T) {
	testApp := New("testApp", "A test for my minimalist cli app building lib")
	testArgInt64 := testApp.NewInt64Arg("int64", "Test int64 arg")
	testArgString := testApp.NewStringArg("string", "Test String arg")

	t.Run("Simple arg parsing", func(t *testing.T) {
		*testArgInt64, *testArgString = 0, ""
		testApp.cmdline = []string{"42", "I'm a string arg"}

		if err := testApp.parseArgs(); err != nil {
			t.Fatalf("Parsing of %#v failed: %v", testApp.cmdline, err)
		}

		if *testArgString != "I'm a string arg" {
			t.Errorf("String arg is not recognized (string is %s)", *testArgString)
		}

		if *testArgInt64 != 42 {
			t.Errorf("Int64 arg is not recognized (int is %d)", *testArgInt64)
		}
	})

	t.Run("Wrong arg number/order/type", func(t *testing.T) {
		*testArgInt64, *testArgString = 0, ""

		testApp.cmdline = []string{"7"}
		if err := testApp.parseArgs(); err == nil {
			t.Errorf("Parsing succeed with a malformed command line (%#v)", testApp.cmdline)
		}
		testApp.cmdline = []string{"3.14", "I'm a string arg"}
		if err := testApp.parseArgs(); err == nil {
			t.Errorf("Parsing succeed with a malformed command line (%#v)", testApp.cmdline)
		}
		testApp.cmdline = []string{"I'm a string", "7"}
		if err := testApp.parseArgs(); err == nil {
			t.Errorf("Parsing succeed with a malformed command line")
		}
	})

	t.Run("Test end of flags", func(t *testing.T) {
		*testArgInt64, *testArgString = 0, ""
		testApp.cmdline = []string{"--", "42", "--unknown"}

		if err := testApp.parseFlags(); err != nil {
			t.Errorf("Parsing failed, %v", err)
		}
		if err := testApp.parseArgs(); err != nil {
			t.Errorf("Parsing failed, %v", err)
		}

		if *testArgString != "--unknown" {
			t.Errorf("String arg is not recognized (string is %s)", *testArgString)
		}

		if *testArgInt64 != 42 {
			t.Errorf("Int64 arg is not recognized (int is %d)", *testArgInt64)
		}
	})
}

func TestCumulativeArg(t *testing.T) {
	testApp := New("testApp", "A test for my minimalist cli app building lib")
	testArgInt64 := testApp.NewInt64Arg("int64", "Test int64 arg")
	testArgStrings := testApp.NewStringsArg("string", "Test String arg")

	testApp.cmdline = []string{"42", "I'm a string arg", "I'm another string arg"}
	if err := testApp.parseArgs(); err != nil {
		t.Fatalf("Parsing of %#v failed: %v", testApp.cmdline, err)
	}

	if !reflect.DeepEqual(*testArgStrings, testApp.cmdline[1:]) {
		t.Errorf("Strings arg is not recognized (string is %s, instead of %v)", *testArgStrings, testApp.cmdline[1:])
	}

	if *testArgInt64 != 42 {
		t.Errorf("Int64 arg is not recognized (int is %d)", *testArgInt64)
	}
}

func TestEnumFlag(t *testing.T) {
	testApp := New("testApp", "A test for my minimalist cli app building lib")
	testFlagBool := testApp.NewBoolFlag("bool", "Test Boolean flag")
	testFlagEnum := testApp.NewEnumFlag("choice", "Test enum flag", []string{"first", "second"})

	testApp.cmdline = []string{"--bool", "--choice=first"}
	if err := testApp.parseFlags(); err != nil {
		t.Fatalf("Parsing failed: %v", err)
	}

	if *testFlagEnum != "first" {
		t.Errorf("Enum flag is not recognized (string is '%s' instead of 'first')", *testFlagEnum)
	}

	if !*testFlagBool {
		t.Errorf("Boolean flag is not recognized")
	}

	testApp.cmdline = []string{"--bool", "--choice=third"}
	if err := testApp.parseFlags(); err == nil {
		t.Fatalf("Parsing does not fail but should have (wrong option selection for enum flag)")
	}
}

func TestSubCommandsParsing(t *testing.T) {
	var testResult string
	testApp := New("testApp", "A test for my minimalist cli app building lib")
	testAppFlag := testApp.NewBoolFlag("bool", "Test bool flag")
	testArgString := testApp.NewStringArg("string", "Test String arg")
	testApp.Execute = func() error { testResult += *testArgString; return nil }

	testCmd := testApp.NewCommand("test", "Test a new sub-command")
	testCmdArg := testCmd.NewStringArg("test_arg", "Test String arg of sub-command")
	testCmd.Execute = func() error { testResult += *testCmdArg; return nil }

	t.Run("Without Subcommand", func(t *testing.T) {
		testResult, *testAppFlag, *testArgString, *testCmdArg = "", false, "", ""
		testApp.cmdline = []string{"--bool", "I'm the test string of test app"}

		if err := testApp.Run(); err != nil {
			t.Fatalf("Run of %#v failed: %v", testApp.cmdline, err)
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
		testApp.cmdline = []string{"--bool", "test", "I'm the test string of test sub-command"}

		if err := testApp.Run(); err != nil {
			t.Fatalf("Run of %#v failed: %v", testApp.cmdline, err)
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

		testApp.cmdline = []string{"--bool"}
		if err := testApp.Run(); err == nil {
			t.Errorf("Running succeed with a malformed command line (%#v)", testApp.cmdline)
		}

		testApp.cmdline = []string{"test"}
		if err := testApp.Run(); err == nil {
			t.Errorf("Running succeed with a malformed command line (%#v)", testApp.cmdline)
		}

		testApp.cmdline = []string{"test", "--bool", "I'm the test string of test sub-command"}
		if err := testApp.Run(); err == nil {
			t.Errorf("Running succeed with a malformed command line (%#v)", testApp.cmdline)
		}
	})
}
