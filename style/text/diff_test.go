package text

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	testA = `Ceci est un test pour trouver une bonne façon
de représenter les diférences entre deux textes ou deux chaînes
de caractères.

Happy end.
`

	testB = `Ceci est un test pour trouver une (très) bonne façon
de représenter les diférences entre deux textes ou deux chaînes
de caractères.

Il s'agirait ensuite de l'inclure dans le package verify pour obtenir
un chouette outil de test.

Happy end.
`
)

func TestDiff(t *testing.T) {
	testCases := []struct {
		inL, inR   string
		outL, outR string
	}{
		{inL: "List things to do.", inR: "List things to do, schedule things to do.",
			outL: "List things to do.", outR: "List things to do~~, schedule things to do~~."},
		{inL: "List thing to do.", inR: "List things to do.",
			outL: "List thing to do.", outR: "List thing~~s~~ to do."},
		{inL: "List things to do.", inR: "Miss things to do.",
			outL: "**L**is**t** things to do.", outR: "**M**is**s** things to do."},
		{inL: "2019-04-10T13:57:13.4696379+02:00", inR: "<no value>",
			outL: "**2019-04-10T13:57:13.4696379+02:00**", outR: "**<no value>**"},
	}

	for _, tc := range testCases {
		_, dL, dR := LightMkdDiff.Bylines(tc.inL, tc.inR)
		gotL, gotR := strings.Join(dL, ""), strings.Join(dR, "")
		if gotL != tc.outL || gotR != tc.outR {
			t.Errorf("Diff between '%s' and '%s' failed:\nWant: %s -> %s\nGot : %s -> %s", tc.inL, tc.inR, tc.outL, tc.outR, gotL, gotR)
		}
	}
}

func TestDiffFullExample(t *testing.T) {
	dT1, dL1, dR1 := ColorDiff.Bylines(testA, testB)
	dT2, dL2, dR2 := ColorDiff.Bylines(testB, testA)
	matchGolden(t, Table().SetMaxWidth(180).Col(dL1, dT1, dR1).String()+"\n\n"+Table().SetMaxWidth(180).Col(dL2, dT2, dR2).String(), "Incorrect diff output")
}

var (
	update = flag.Bool("test.update", false, "update golden file with test result")
)

func matchGolden(tb testing.TB, got string, message string) {
	if *update {
		updateGolden(tb, []byte(got))
	}

	expected := readGolden(tb)

	tb.Logf("Got:\n%s\n\nExpected:\n%s\n", got, string(expected))
	if string(expected) != got {
		tb.Errorf("%s\n", message)
	}
}

func goldenPath(tb testing.TB) string {
	return filepath.Join("./testdata", tb.Name()+".golden")
}

func readGolden(tb testing.TB) []byte {
	f := goldenPath(tb)

	expected, err := ioutil.ReadFile(f)
	if err != nil {
		tb.Logf("cannot read golden file %s: %v", f, err)
		return []byte{}
	}
	return expected
}

func updateGolden(tb testing.TB, actual []byte) {
	f := goldenPath(tb)

	tb.Logf("update golden file %s", f)
	if err := ioutil.WriteFile(f, actual, os.ModePerm); err != nil {
		tb.Fatalf("cannot update golden file %s: %v", f, err)
	}
}
