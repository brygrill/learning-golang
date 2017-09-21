package main

import (
	"fmt"
)

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func halffloat(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}

func main() {
	h, e := half(21)
	fmt.Println(h, e)

	i, j := halffloat(21)
	fmt.Println(i, j)

	a := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	fmt.Println(a(21))
}
