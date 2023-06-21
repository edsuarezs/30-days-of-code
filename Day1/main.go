package main

import (
	"fmt"
)

func main() {
	var (
		i uint64  = 4
		d float64 = 4.0
		s string  = "HackerRank "

		a uint64
		b float64
		c string
	)
	// Read and save an integer, double, and String to your variables.
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	// Print the sum of the int variables on a new line.
	fmt.Println(i + a)
	// Print the sum of the double variables on a new line.
	db := d + b
	fmt.Printf("%.1f\n", db)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + c)

}
