package ansi_test

import (
	"fmt"

	"github.com/pirmd/text/ansi"
)

func Example() {
	fmt.Printf("%s %s le monde !\n", ansi.Red("Bonjour"), ansi.Bold("tout"))
	fmt.Print(ansi.GreenOn, "Have ", ansi.BoldOn, "fun ", ansi.BoldOff, "with ", ansi.BlueOn, "Colors", ansi.Reset)
	//Output:
	//[31mBonjour[39m [1mtout[22m le monde !
	//[32mHave [1mfun [22mwith [34mColors[0m
}

func ExampleWalkString() {
	hello := ansi.Red("Bonjour") + " " + ansi.Bold("tout") + " le monde !"

	var visualen int
	_ = ansi.WalkString(hello, func(n int, c rune, esc string) error {
		if c > -1 {
			visualen++
		}
		return nil
	})

	fmt.Printf("Length is %d, visual length is %d", len(hello), visualen)
	//Output:
	//Length is 42, visual length is 23
}

func ExampleWalkString_second() {
	hello := ansi.Red("Bonjour") + " " + ansi.Bold("tout") + " le monde !"

	var clean []rune
	_ = ansi.WalkString(hello, func(n int, c rune, esc string) error {
		if c > -1 {
			clean = append(clean, c)
		}
		return nil
	})

	fmt.Println(string(clean))
	//Output:
	//Bonjour tout le monde !
}
