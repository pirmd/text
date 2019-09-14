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
