package ansi_test

import (
	"fmt"

	"github.com/pirmd/text/ansi"
)

func Example() {
	hello := ansi.SetRed("Bonjour") + " " + ansi.SetBold("tout") + " le monde !"
	fmt.Println(ansi.RemoveANSI(hello))
	fmt.Printf("Length is %d, visual length is %d", len(hello), ansi.Len(hello))
	//Output:
	//Bonjour tout le monde !
	//Length is 42, visual length is 23
}
