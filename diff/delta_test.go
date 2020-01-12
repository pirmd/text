package diff

import (
	"testing"

	"github.com/pirmd/verify"

	"github.com/pirmd/text"
)

func TestResultPrettyPrint(t *testing.T) {
	testCases := []struct {
		in                  Result
		wanta, wantb, wantt string
	}{
		{
			Result{&diff{IsSame, "a"}, &diff{IsInserted, "b"}, &diff{IsSame, "c"}},
			"ac",
			"abc",
			"=+=",
		},

		{
			Result{Result{&diff{IsSame, "a"}, &diff{IsInserted, "b"}, &diff{IsSame, "c"}}},
			"ac",
			"abc",
			"+",
		},

		{
			Result{&diff{IsSame, "ab\n"}, &diff{IsInserted, "cd\n"}, &diff{IsSame, "ef"}},
			"ab\n\nef",
			"ab\ncd\nef",
			"=\n+\n=",
		},

		{
			Result{
				&diff{IsSame, "ab\n"},
				Result{&diff{IsDeleted, "dc"}, &diff{IsInserted, "cd"}, &diff{IsSame, "\n"}},
				&diff{IsSame, "ef"},
			},
			"ab\ndc\nef",
			"ab\ncd\nef",
			"=\n<>\n=",
		},

		{
			Result{
				&diff{IsSame, "a\nb\n"},
				Result{&diff{IsDeleted, "dc"}, &diff{IsInserted, "cd"}, &diff{IsSame, "\n"}},
				&diff{IsSame, "ef"},
			},
			"a\nb\ndc\nef",
			"a\nb\ncd\nef",
			"=\n\n<>\n=",
		},

		{
			Result{
				&diff{IsInserted, "#include \"diff.h\"\n"},
				Result{
					&diff{IsSame, "#include \""},
					&diff{IsDeleted, "old"}, &diff{IsInserted, "new"},
					&diff{IsSame, ".h\"\n"},
				},
			},
			"\n#include \"old.h\"\n",
			"#include \"diff.h\"\n#include \"new.h\"\n",
			"+\n<>\n",
		},

		{
			Result{
				Result{
					&diff{IsSame, "#include \""},
					&diff{IsDeleted, "old"},
					&diff{IsInserted, "new"},
					&diff{IsSame, ".h\"\n"},
				},
				&diff{IsInserted, "#include \"diff.h\"\n"}, &diff{IsInserted, "\n"},
				&diff{IsDeleted, "char buf[64]\n"},
			},
			"#include \"old.h\"\n\n\nchar buf[64]\n",
			"#include \"new.h\"\n#include \"diff.h\"\n\n\n",
			"<>\n+\n+\n-\n",
		},
	}

	for _, tc := range testCases {
		gota, gotb, gott := tc.in.PrettyPrint()
		t.Logf("\n" + text.NewTable().SetMaxWidth(180).Rows([]string{tc.wanta, tc.wantt, tc.wantb}).Draw())
		t.Logf("\n" + text.NewTable().SetMaxWidth(180).Rows([]string{gota, gott, gotb}).Draw())
		verify.Equal(t, gota, tc.wanta, "Result of\n%#v failed (for 'a' side)", tc.in.GoString())
		verify.Equal(t, gotb, tc.wantb, "Result of\n%#v failed (for 'b' side)", tc.in.GoString())
		verify.Equal(t, gott, tc.wantt, "Result of  %#v failed (for 't' side)", tc.in.GoString())
		//XXX: sort out verify Equal args issue -> accept interface and not string
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
		verify.Equal(t, got, tc.want, "Identification of changed zones for %s failed", tc.in.GoString())
	}
}
