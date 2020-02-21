package diff

import (
	"reflect"
	"testing"
)

func TestAdaptativeDiff(t *testing.T) {
	testCases := []struct {
		inL, inR string
		want     Result
	}{
		{
			"import (\n\t\"strings\"\n\t\"path\"\n)",
			"import (\n\t\"os\"\n\t\"strings\"\n\t\"path/filepath\"\n)",
			Result{
				&diff{IsSame, "import (\n"},
				&diff{IsInserted, "\t\"os\"\n"},
				&diff{IsSame, "\t\"strings\"\n"},
				Result{
					&diff{IsSame, "\t"},
					&diff{IsSame, "\""},
					&diff{IsSame, "path"},
					&diff{IsInserted, "/"},
					&diff{IsInserted, "filepath"},
					&diff{IsSame, "\""},
					&diff{IsSame, "\n"},
				},
				&diff{IsSame, ")\n"},
			},
		},
	}

	for _, tc := range testCases {
		got := adaptative(diffLCS, tc.inL, tc.inR, ByLines, ByWords, ByRunes)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Adaptative diff between\n%s\nand\n%s\nfailed.\nWant: %#v\nGot : %#v.", tc.inL, tc.inR, tc.want, got)
		}
	}
}
