package main

import (
	"fmt"
)

func changeMe(z *int) { // z is accepting the memory location
	fmt.Println(z)  // print memory location
	fmt.Println(*z) // print value

	*z = 24 // change value at memory location
}

func refType1(z map[string]int) {
	z["Hi"] = 44
}

func refType2(z []string) {
	z[0] = "Hi"
}

type customer struct {
	name string
	age  int
}

func changeMe2(z *customer) {
	z.name = "Not Carson"
}

func main() {
	age := 44
	fmt.Println(&age) // memory location

	changeMe(&age) // must pass in memory location

	fmt.Println(age) //24 and still at same memory location

	m := make(map[string]int)
	refType1(m)
	fmt.Println(m["Hi"])

	n := make([]string, 1, 25)
	fmt.Println(n)
	refType2(n)
	fmt.Println(n)

	c1 := customer{"Carson", 24}
	fmt.Println(c1)
	changeMe2(&c1)
	fmt.Println(c1)
}
