package main

import "fmt"

// vars with zero vals
func varTypes() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println()
}

// shorthand vars can only use inside func
func varVals() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	fmt.Printf("%v - %T \n", a, a)
	fmt.Printf("%v - %T \n", b, b)
	fmt.Printf("%v - %T \n", c, c)
	fmt.Printf("%v - %T \n", d, d)
	fmt.Printf("%v - %T \n", e, e)
	fmt.Printf("%v - %T \n", f, f)
	fmt.Printf("%v - %T \n", g, g)
}

func main() {
	varTypes()
	varVals()
}
