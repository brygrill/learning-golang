package main

import "fmt"

func main() {
	a := 15
	location := &a
	fmt.Println("memory address: ", &a)
	fmt.Println("memory address: ", location)
	fmt.Printf("%d \n", location)

	usingmem()
}
