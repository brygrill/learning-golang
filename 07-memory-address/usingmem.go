package main

import "fmt"

const metersToYards float64 = 1.09361

func usingmem() {
	var meters float64
	fmt.Print("Meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println("Yards swam: ", yards)
}
