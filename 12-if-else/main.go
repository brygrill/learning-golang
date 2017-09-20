package main

import (
	"fmt"
)

func basic() {
	if true {
		fmt.Println("This will print")
	}
	if false {
		fmt.Println("This will not print")
	}
}

func basicElse() {
	if false {
		fmt.Println("Print here")
	} else {
		fmt.Println("Print else")
	}
}

func basicElseIf() {
	if false {
		fmt.Println("first print statement")
	} else if true {
		fmt.Println("second print statement")
	} else {
		fmt.Println("third print statement")
	}
}

func exclaim() {
	if !true {
		fmt.Println("This will not print")
	}
	if !false {
		fmt.Println("This will print")
	}
}

func initStmt(myfood string) {
	b := true

	if food := myfood; b {
		fmt.Println(food)
	}
}

func main() {
	basic()
	basicElse()
	basicElseIf()
	exclaim()
	initStmt("Cake")
}
