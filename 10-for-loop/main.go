// control flow
// - sequencing
// - loops
// - conditionals
// https://golang.org/doc/effective_go.html#for
package main

import (
	"fmt"
)

func basicLoop() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func main() {
	basicLoop()
}
