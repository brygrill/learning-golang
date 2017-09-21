package main

import (
	"fmt"
)

func max(sf ...int) int {
	var g int
	for _, v := range sf {
		if v > g {
			g = v
		}
	}
	return g
}

func main() {
	fmt.Println(max(-1, 1, 2, 3, 4, 5, 6, 7))
}
