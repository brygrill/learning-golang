package main

import (
	"fmt"
)

func main() {
	// true or false
	if true {
		fmt.Println("True")
	}

	// not
	if !false {
		fmt.Println("Not False")
	}

	if true || false {
		fmt.Println("One is true")
	}

	if false && !true {
		fmt.Println("Wont print")
	}
}
