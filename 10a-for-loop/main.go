package main

import (
	"fmt"
)

// basic for loop
func forLoop() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

// nested for loop
func nestedForLoop() {
	for q := 0; q < 5; q++ {
		for p := 0; p < 10; p++ {
			fmt.Printf("%v - %v \n", q, p)
		}
	}
}

// kind of like a while
func kindofWhile() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

// no condition - infinite
func forever() {
	i := 0
	for {
		fmt.Println(i)
		i++
	}
}

func withBreak() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			// if i is greater than 10, exit the loop
			break
		}
		i++
	}
}

func withContinue() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			// if even jump to top of the loop
			continue
		}
		fmt.Println(i) // will print odd numbers
		if i >= 50 {
			break
		}
	}
}

func withRune() {
	for i := 0; i < 500; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}

func main() {
	forLoop()
	nestedForLoop()
	kindofWhile()
	// forever()
	withBreak()
	withContinue()
	withRune()
}
