// Package numutil contains utility functions for working with and formatting numbers.
// https://github.com/GoesToEleven/GolangTraining/tree/master/01_getting-started/02_numeral-systems
// https://golang.org/pkg/fmt/?m=all#hdr-Printing
package numutil

import (
	"fmt"
)

// PrintDecimal will print the supplied decimal to the console
func PrintDecimal(dec int) int {
	fmt.Printf("Decimal: %d \n", dec)
	return dec
}

// PrintBinary will print the supplied decimal to the console
// as a decimal and in binary value
func PrintBinary(dec int) int {
	fmt.Printf("Decimal: %d - Binary: %b \n", dec, dec)
	return dec
}

// PrintHex will print the supplied decimal to the console
// as a decimal and as a hexadecimal value
func PrintHex(dec int) int {
	fmt.Printf("Decimal: %d - Hexadecimal: %x \n", dec, dec)
	return dec
}

// PrintRange will print the supplied range of values to the console
// as a decimal, a binary value and a hexadecimal value
func PrintRange(start int, end int) int {
	print("\n")
	for i := start; i <= end; i++ {
		fmt.Printf("Decimal: %d - Binary: %b - Hexadecimal: %x \n", i, i, i)
	}
	return end
}