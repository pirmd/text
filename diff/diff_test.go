package diff_test

import (
	"testing"

	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pirmd/text/diff"
)

var (
	testCases = []struct {
		name     string
		inL, inR string
	}{

		{
			name: "Code",
			inL: `#include "win31.h"
#include "notfullyworkingyet.h"

char make_prog_look_big[800000];

void main()
{
	display_copyright_message();
	basically_run_windows_3.1();

	/* printf("Welcome to Windows 3.11"); */
	printf("Welcome to Windows 95");
	if (system_ok_for_too_long())
	{
		bsod(random_err());
		crash(to_dos_prompt);
	}
	else
	system_memory = open("swp0001.swp", O_CREATE);

	while(!system_up_for_too_long())
	{
		sleep(15);
		get_user_input();
		sleep(15);
		act_on_user_input();
		sleep(15);
	}
	create_general_protection_fault();
}`,

			inR: `#include "win31.h"
#include "win95.h"
#include "stillnotfullyworking.h"

char make_prog_look_big[1600000];

void main()
{
	if (fast_cpu())
	{
		set_wait_states(lots);
		set_mouse(speed, very_slow);
	}

	while(LESS_THAN_FOREVER)
	{
		display_copyright_message();
		simulate_important_action();
		basically_run_windows_3.1();
		make_think_we_are_busy();
	}

	/* printf("Welcome to Windows 3.11"); */
	/* printf("Welcome to Windows 95"); */
	printf("Welcome to Windows 98");
	if (system_ok_for_too_long())
	{
		bsod(random_err());
		crash(to_dos_prompt);
	}
	else
	system_memory = open("swp0001.swp", O_CREATE);

	while(!system_up_for_too_long())
	{
		sleep(5);
		get_user_input();
		sleep(5);
		act_on_user_input();
		sleep(5);
	}
	create_general_protection_fault();
}`,
		},

		{
			name: "JSON-like",
			inL: `{ 
	"title": "Alice's Adventures in Wonderland",
	"authors": [
	"Lewis Carroll"
	],
	"description": "This edition contains Alice's Adventures in Wonderland. Tweedledum and Tweedledee, the Mad Hatter, the Cheshire Cat, the Red Queen and the White Rabbit all make their appearances, and are now familiar figures in writing, conversation and idiom.",
	"pageCount": 132,
	"language": "gb",
}`,
			inR: `{ 
	"title": "Alice's Adventures in Wonderland & Through the Looking-glass",
	"authors": [
	"Lewis Carroll"
	],
	"description": "This edition contains Alice's Adventures in Wonderland and its sequel Through the Looking Glass. It is illustrated throughout by Sir John Tenniel, whose drawings for the books add so much to the enjoyment of them. Tweedledum and Tweedledee, the Mad Hatter, the Cheshire Cat, the Red Queen and the White Rabbit all make their appearances, and are now familiar figures in writing, conversation and idiom. So too, are Carroll's delightful verses such as 'The Walrus and the Carpenter' and the inspired jargon of that masterly Wordsworthian parody, 'The Jabberwocky'.",
	"pageCount": 272,
	"categories": [
	"Fiction"
	],
	"averageRating": 4.0,
	"language": "en",
}`,
		},

		{
			name: "Prose",
			inL: `Ceci est un test pour trouver une bonne façon
de représenter les diférences entre deux textes ou deux chaînes
de caractères.

Happy end.`,
			inR: `Ceci est un test pour trouver une (très) bonne façon
de représenter les différences entre deux textes ou deux chaînes
de caractères.

Il s'agirait ensuite de l'inclure dans le package verify pour obtenir
un outil de test.

Happy end.`,
		},

		{
			name: "Poetry",
			inL: `Three Rings for the Elven-kings under the sky,
Seven for the Dwarf-lords in their halls of stone,
Nine for Mortal Men doomed to die,
One for the Light Lord on his light throne
In the Land of Mordor where the Lights shine.
One Ring to serve them all, One Ring to help them,
In the Land of Mordor where the Lights shine.`,

			inR: `Three Rings for the Elven-kings under the sky,
Seven for the Dwarf-lords in their halls of stone,
Nine for Mortal Men doomed to die,
One for the Dark Lord on his dark throne
In the Land of Mordor where the Shadows lie.
One Ring to rule them all, One Ring to find them,
One Ring to bring them all and in the darkness bind them
In the Land of Mordor where the Shadows lie.`,
		},

		{
			name: "List",
			inL: `Whatever goes upon two legs is an enemy.
Whatever goes upon four legs, or has wings, is a friend.
No animal shall wear clothes.
No animal shall sleep in a bed.
No animal shall drink alcohol.
No animal shall kill any other animal.
All animals are equal.`,
			inR: `Four legs good, two legs better.
No animal shall sleep in a bed without sheets.
No animal shall drink alcohol to excess.
No animal shall kill any other animal without cause.
All animals are equal but some are more equal than others.`,
		},
	}
)

func TestLCSDiffByLines(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := diff.LCS(tc.inL, tc.inR, diff.ByLines)
			diffTable := d.PrintSideBySide(diff.WithColor, diff.WithoutMissingContent)
			if err := matchGolden(t.Name(), diffTable); err != nil {
				t.Errorf("LCS diff is not working as expected. %v", err)
			}
		})
	}
}

func TestLCSDiffByLinesByWordsByRunes(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := diff.LCS(tc.inL, tc.inR, diff.ByLines, diff.ByWords, diff.ByRunes)
			diffTable := d.PrintSideBySide(diff.WithColor, diff.WithoutMissingContent)
			if err := matchGolden(t.Name(), diffTable); err != nil {
				t.Errorf("LCS diff is not working as expected. %v", err)
			}
		})
	}
}

func TestPatienceDiffByLines(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := diff.Patience(tc.inL, tc.inR, diff.ByLines)
			diffTable := d.PrintSideBySide(diff.WithColor, diff.WithoutMissingContent)
			if err := matchGolden(t.Name(), diffTable); err != nil {
				t.Errorf("Patience diff is not working as expected. %v", err)
			}
		})
	}
}

func TestPatienceDiffByLinesByWordsByRunes(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := diff.Patience(tc.inL, tc.inR, diff.ByLines, diff.ByWords)
			diffTable := d.PrintSideBySide(diff.WithColor)
			if err := matchGolden(t.Name(), diffTable); err != nil {
				t.Errorf("LCS diff is not working as expected. %v", err)
			}
		})
	}
}

var (
	updateGolden = flag.Bool("test.golden-update", false, "update golden file with test result")
	goldenDir    = flag.String("test.goldendir", "./testdata", "path to folder hosting golden files")
)

func matchGolden(name string, got string) error {
	goldenPath := filepath.Join(*goldenDir, name+".golden")

	if *updateGolden {
		if err := updateGoldenFiles(goldenPath, []byte(got)); err != nil {
			return err
		}
	}

	want, err := readGolden(goldenPath)
	if err != nil {
		return err
	}

	if len(want) == 0 {
		return fmt.Errorf("no existing or empty golden file.\nTest output is:\n%s", got)
	}

	if got != string(want) {
		return fmt.Errorf("Got:\n%v\n\nWant :\n%v", got, string(want))
	}
	return nil
}

func readGolden(path string) ([]byte, error) {
	want, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read golden file %s: %v", path, err)
	}
	return want, nil
}

func updateGoldenFiles(path string, actual []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return fmt.Errorf("cannot update golden file %s: %v", path, err)
	}

	if err := ioutil.WriteFile(path, actual, 0666); err != nil {
		return fmt.Errorf("cannot update golden file %s: %v", path, err)
	}

	return nil
}
