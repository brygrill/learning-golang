package main

import (
	"fmt"
)

func hi(name string) { // hi is declared with a parameter
	fmt.Println("Hi", name)
}

func hiTwo(fname, lname string) { // if both params are same type, you can only define type once
	fmt.Println("Hi", fname, lname)
}

func basicReturn(name string) string { // take param of string and returns a string
	return fmt.Sprint("Your Name: ", name)
}

func namedReturn(name string) (s string) {
	s = fmt.Sprint("Your Name: ", name)
	return

	/*
		IMPORTANT
		Avoid using named returns.


		Occasionally named returns are useful. Read this article for more information:
		https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
	*/
}

func multiReturn(fname, lname string) (string, string) {
	return fmt.Sprint("Your First Name: ", fname), fmt.Sprint("Your First Name: ", lname)
}

func variadicParam(sf ...float64) float64 {
	var ttl float64
	for _, v := range sf {
		ttl += v
	}
	return ttl / float64(len(sf))
}

func sliceParam(sf []float64) float64 {
	var ttl float64
	for _, v := range sf {
		ttl += v
	}
	return ttl / float64(len(sf))
}

func funcExpression() {
	sayHi := func() {
		fmt.Println("Hi")
	}

	sayHi()
}

func funcExp2() func() string {
	return func() string {
		return "My Func Expression"
	}
}

func callbackEx(numbers []int, callback func(int)) {
	for _, v := range numbers {
		callback(v)
	}
}

func callbackEx2(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, v := range numbers {
		if callback(v) {
			xs = append(xs, v)
		}
	}
	return xs
}

func recursion(x int) int {
	// if x is 0 return 1
	if x == 0 {
		return 1
	}
	// otherwise call this self and run
	// on one less than x
	return x * recursion(x-1)
}

func deferEx() {
	hello := func() {
		fmt.Println("Hello")
	}

	world := func() {
		fmt.Println("World")
	}

	defer world()
	hello()
}

// main is the entry point to your program
func main() {
	f := "Carson"
	l := "Wentz"
	hi(f) // when calling hi, pass in an argument
	hiTwo(f, l)

	fmt.Println(basicReturn(f))
	fmt.Println(namedReturn(l))

	first, last := multiReturn(f, l)
	fmt.Println(first)
	fmt.Println(last)

	// variadic param
	avg := variadicParam(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(avg)

	// variadic args
	data := []float64{10, 11, 12, 13, 14, 15}
	avg2 := variadicParam(data...)
	fmt.Println(avg2)

	avg3 := sliceParam(data)
	fmt.Println(avg3)

	funcExpression()

	myExp := funcExp2()
	fmt.Println(myExp())

	callbackEx([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

	xs := callbackEx2([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)

	fmt.Println(recursion(4))

	deferEx()

	func() {
		fmt.Println("Self Executing!")
	}()
}
