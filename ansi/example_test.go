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

func ExampleLen() {
	hello := ansi.Red("Bonjour") + " " + ansi.Bold("tout") + " le monde !"
	fmt.Printf("Length is %d, visual length is %d", len(hello), ansi.Len(hello))
	//Output:
	//Length is 42, visual length is 23
}

func ExampleRemoveANSI() {
	hello := ansi.Red("Bonjour") + " " + ansi.Bold("tout") + " le monde !"
	fmt.Println(ansi.RemoveANSI(hello))
	//Output:
	//Bonjour tout le monde !
}
