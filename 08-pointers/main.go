// Pointers
// https://tour.golang.org/moretypes/1
package main

import (
	"fmt"
)

// makes z a sep value of 0
func zero(z int) {
	fmt.Printf("%p\n", &z)
	fmt.Println(&z)
	z = 0
}

// set the value at the location of z to 0
func zeroPointer(z *int) {
	*z = 0
}

func main() {
	a := 43
	var b = &a

	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int
	fmt.Println(&a) // loc in mem
	fmt.Println(b)  // loc in mem

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
	// the * is an operator in this case
	fmt.Println(*b) // 43

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value
	*b = 42
	fmt.Println(a) // 42

	// using pointers
	x := 5
	fmt.Printf("%p\n", &x)
	zero(x)
	fmt.Println(x)

	// change address within function
	y := 9
	zeroPointer(&y) // pass in location
	fmt.Println(y)  // now y is 0
}
