package app

import (
	"reflect"
	"testing"
)

func TestStringsValue(t *testing.T) {
	var args []string

	v := newStringsValue(&args)

	test := []string{"1", "2", "3"}
	for _, a := range test {
		if err := v.Set(a); err != nil {
			t.Errorf("Assignement of arg to %v failed: %v", a, err)
		}
	}

	if !reflect.DeepEqual(args, test) {
		t.Errorf("stringsValue failed, got '%v', wanted '%v'", args, test)
	}
}

func TestEnumValue(t *testing.T) {
	var arg string

	v := newEnumValue(&arg, "blue", "red", "yellow")

	test := "green"
	if err := v.Set(test); err == nil {
		t.Errorf("Assignement of arg to '%s' did not fail (accepted options %v)", test, v.options)
	}

	test = "red"
	if err := v.Set(test); err != nil {
		t.Errorf("Assignement of arg to '%s' failed (accepted options %v): %v", test, v.options, err)
	}

	if arg != test {
		t.Errorf("enumValue failed, got '%v', wanted '%v'", arg, test)
	}
}
