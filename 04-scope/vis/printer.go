package vis

import "fmt"

// PrintVar is exported because it starts with a capital letter
func PrintVar() {
	fmt.Println(MyName)
	// PrintVar has access to yourName therefore making it accessable
	// outside the package
	fmt.Println(yourName)
}