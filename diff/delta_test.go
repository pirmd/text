package diff

import (
	"reflect"
	"testing"

	"github.com/pirmd/text/table"
)

func TestResultPrettyPrint(t *testing.T) {
	testCases := []struct {
		in                  Result
		wantL, wantR, wantT []string
	}{
		{
			Result{&diff{IsSame, "a"}, &diff{IsInserted, "b"}, &diff{IsSame, "c"}},
			[]string{"a", "", "c"},
			[]string{"a", "b", "c"},
			[]string{"=", "+", "="},
		},

		{
			Result{Result{&diff{IsSame, "a"}, &diff{IsInserted, "b"}, &diff{IsSame, "c"}}},
			[]string{"ac"},
			[]string{"abc"},
			[]string{"+"},
		},

		{
			Result{&diff{IsSame, "ab\n"}, &diff{IsInserted, "cd\n"}, &diff{IsSame, "ef"}},
			[]string{"ab\n", "", "ef"},
			[]string{"ab\n", "cd\n", "ef"},
			[]string{"=", "+", "="},
		},

		{
			Result{
				&diff{IsSame, "ab\n"},
				Result{&diff{IsDeleted, "dc"}, &diff{IsInserted, "cd"}, &diff{IsSame, "\n"}},
				&diff{IsSame, "ef"},
			},
			[]string{"ab\n", "dc\n", "ef"},
			[]string{"ab\n", "cd\n", "ef"},
			[]string{"=", "<>", "="},
		},

		{
			Result{
				&diff{IsSame, "a\nb\n"},
				Result{&diff{IsDeleted, "dc"}, &diff{IsInserted, "cd"}, &diff{IsSame, "\n"}},
				&diff{IsSame, "ef"},
			},
			[]string{"a\nb\n", "dc\n", "ef"},
			[]string{"a\nb\n", "cd\n", "ef"},
			[]string{"=", "<>", "="},
		},

		{
			Result{
				&diff{IsInserted, "#include \"diff.h\"\n"},
				Result{
					&diff{IsSame, "#include \""},
					&diff{IsDeleted, "old"},
					&diff{IsInserted, "new"},
					&diff{IsSame, ".h\"\n"},
				},
			},
			[]string{"", "#include \"old.h\"\n"},
			[]string{"#include \"diff.h\"\n", "#include \"new.h\"\n"},
			[]string{"+", "<>"},
		},

		{
			Result{
				Result{
					&diff{IsSame, "#include \""},
					&diff{IsDeleted, "old"},
					&diff{IsInserted, "new"},
					&diff{IsSame, ".h\"\n"},
				},
				&diff{IsInserted, "#include \"diff.h\"\n"},
				&diff{IsInserted, "\n"},
				&diff{IsDeleted, "char buf[64]\n"},
			},
			[]string{"#include \"old.h\"\n", "", "", "char buf[64]\n"},
			[]string{"#include \"new.h\"\n", "#include \"diff.h\"\n", "\n", ""},
			[]string{"<>", "+", "+", "-"},
		},
	}

	for _, tc := range testCases {
		gotL, gotR, gotT, _ := tc.in.PrettyPrint()
		if !reflect.DeepEqual(gotL, tc.wantL) || !reflect.DeepEqual(gotR, tc.wantR) || !reflect.DeepEqual(gotT, tc.wantT) {
			t.Logf("Want:\n" + table.New().SetMaxWidth(180).Col(tc.wantL, tc.wantT, tc.wantR).Draw())
			t.Logf("Got :\n" + table.New().SetMaxWidth(180).Col(gotL, gotT, gotR).Draw())
			t.Errorf("Pretty printing\n %#v\n failed.", tc.in)
		}
	}
}

func TestDifferentZones(t *testing.T) {
	testCases := []struct {
		in   *Result
		want [][2]int
	}{
		{&Result{&diff{IsSame, "a"}, &diff{IsInserted, "b"}, &diff{IsSame, "c"}}, nil},
		{&Result{&diff{IsSame, "a"}, &diff{IsInserted, "b"}, &diff{IsInserted, "c"}}, nil},
		{&Result{&diff{IsSame, "a"}, &diff{IsInserted, "b"}, &diff{IsSame, "c"}, &diff{IsDeleted, "d"}}, nil},
		{&Result{&diff{IsSame, "a"}, &diff{IsInserted, "b"}, &diff{IsDeleted, "c"}}, [][2]int{{1, 2}}},
		{&Result{&diff{IsDeleted, "a"}, &diff{IsInserted, "b"}, &diff{IsSame, "c"}, &diff{IsDeleted, "d"}}, [][2]int{{0, 1}}},
		{&Result{&diff{IsInserted, "a"}, &diff{IsDeleted, "b"}, &diff{IsSame, "c"}, &diff{IsInserted, "e"}, &diff{IsDeleted, "d"}}, [][2]int{{0, 1}, {3, 4}}},
		{&Result{&diff{IsInserted, "a"}, &diff{IsInserted, "b"}, &diff{IsInserted, "c"}, &diff{IsDeleted, "e"}, &diff{IsDeleted, "d"}}, [][2]int{{0, 4}}},
	}

	for _, tc := range testCases {
		got := tc.in.differentZones()
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Identification of changed zones for %#v failed.\nWant: %v\nGot : %v.", tc.in, tc.want, got)
		}
	}
}
