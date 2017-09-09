package main

import (
	"fmt"

	"github.com/brygrill/learning-golang/04-scope/closureex"
	"github.com/brygrill/learning-golang/04-scope/vis"
)

// x is available at the package level
var x = 42

func main() {
	a := "scoped to this function"
	fmt.Println(x)
	foo()
	// y is from same.go and is available at package level
	fmt.Println(y)
	// MyName is available bc it is exported
	fmt.Println(vis.MyName)
	// yourName is called within the exported func so it will work
	// even though its not exported
	vis.PrintVar()
	// access to a here
	fmt.Println(a)
	// closure example
	closureex.ClosureFunc()
}

func foo() {
	fmt.Println(x)
}

func noaccesstoa() {
	// can not access a from here
	// fmt.Println(a)
}
