package main

import (
	"fmt"
)

func basicSwitch(name string) {
	switch name {
	case "Brandon":
		fmt.Println("Graham")
	case "Fletcher":
		fmt.Println("Cox")
	case "Timmy":
		fmt.Println("Jernigan")
	case "Vinny":
		fmt.Println("Curry")
	default:
		fmt.Println("Not an Eagles defensive lineman")
	}
}

func fallThroughSwitch(name string) {
	switch name {
	case "Brandon":
		fmt.Println("Graham")
	case "Fletcher":
		fmt.Println("Cox")
	case "Timmy":
		fmt.Println("Jernigan")
		fallthrough
	case "Vinny":
		fmt.Println("Curry")
	default:
		fmt.Println("Not an Eagles defensive lineman")
	}
}

func multiEvalSwitch(name string) {
	switch name {
	case "Brandon", "Fletcher":
		fmt.Println("Defense")
	case "Carson", "Alshon":
		fmt.Println("Offense")
	default:
		fmt.Println("Not sure if its an Eagle")
	}
}

func noExpressionSwitch(name string) {
	switch {
	case name == "Brandon":
		fmt.Println("Graham")
	case name == "Fletcher":
		fmt.Println("Cox")
	case name == "Timmy":
		fmt.Println("Jernigan")
	case name == "Vinny":
		fmt.Println("Curry")
	default:
		fmt.Println("Not an Eagles defensive lineman")
	}
}

type contact struct {
	name    string
	address string
	age     int
}

func onTypeSwitch(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("idk")
	}
}

func main() {
	me := contact{"Bryan", "Main St", 34}
	basicSwitch("Fletcher")
	fallThroughSwitch("Timmy")
	multiEvalSwitch("Carson")
	noExpressionSwitch("Brandon")
	onTypeSwitch(42)
	onTypeSwitch(me)
	onTypeSwitch(me.age)
}
