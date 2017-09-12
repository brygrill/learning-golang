package closureex

import "fmt"

// ClosureFunc is an example of closure
func ClosureFunc() {
	b := "hi"
	{
		c := "bye"
		fmt.Println(c)
	}
	// cant print c here
	fmt.Println(b)
	// fmt.Println(c)
}

// AnonFunc is an example of an anonymous function
func AnonFunc() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

// function that returns function
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// CallWrapper is a function that calls another func
func CallWrapper() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

anonymous function
a function without a name

func expression
assigning a func to a variable
*/
