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
