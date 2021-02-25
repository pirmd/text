package ansi

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	testCases := []struct {
		in      string
		want    string
		wantSGR Sequence
	}{
		{
			in:      "\x1b[31mBonjour\x1b[m",
			want:    "Bonjour",
			wantSGR: Sequence{cReset},
		},
		{
			in:      "\x1b[31mBonjour\x1b[m, tout le \x1b[2mmonde !",
			want:    "Bonjour, tout le monde !",
			wantSGR: Sequence{cReset, cFaint},
		},
	}

	for _, tc := range testCases {
		var got string
		var gotSGR Sequence
		err := Walk(tc.in, func(curRune rune, curEsc string) error {
			if curRune > 0 {
				t.Logf("Walk gave a new rune: %s", string(curRune))
				got += string(curRune)
			} else {
				t.Logf("Walk gave a new SGR seq: %s", curEsc)
				gotSGR.Combine(curEsc)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("Walk failed for %#v: %v.", tc.in, err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Walk did not work as expected for %#v (non-ANSI string).\nWant: %v\nGot : %v.", tc.in, tc.want, got)
		}
		if !reflect.DeepEqual(gotSGR, tc.wantSGR) {
			t.Errorf("Walk did not work as expected for %#v (SGR).\nWant: %v\nGot : %v.", tc.in, tc.wantSGR, gotSGR)
		}
	}
}
