// constants
// https://golang.org/doc/effective_go.html#constants
package main

import (
	"fmt"
)

// single const
const always = "a constant"

// multiple
const (
	a = "a"
	b = 1234
)

// iota
const (
	d = iota // 0
	e = iota // 1
	f = iota // 2
)

const (
	g = iota // 0
	h        // 1
	i        // 2
)

// an iota in a const initializes at 0
// from there additional constants can be modified
const (
	_ = iota      // ignore first value by assigning to blank identifier
	j = iota * 10 // 1 x 10
	k = iota * 10 // 2 x 10
)

// Create a huge number by shifting a 1 bit left 100 places.
// In other words, the binary number that is 1 followed by 100 zeroes.
// https://tour.golang.org/basics/16
const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Println(always)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(f) // 2

	fmt.Println(i) // 2

	fmt.Println(k) // 20

	fmt.Printf("%b\t", KB) // base 2
	fmt.Printf("%d\t", KB) // base 10
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
