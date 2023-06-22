package main

import (
	"bufio"
	"fmt"
	"math"
)

/*
 * Complete the 'solve' function below.
 *
 * The function accepts following parameters:
 *  1. DOUBLE meal_cost
 *  2. INTEGER tip_percent
 *  3. INTEGER tax_percent
 */

func solve(meal_cost float64, tip_percent, tax_percent int32) {
	// Write your code here
	tip := float64(tip_percent) / 100
	tax := float64(tax_percent) / 100
	total := meal_cost + (tip * meal_cost) + (tax * meal_cost)
	result := int(math.Round(total))
	fmt.Printf("%v", result)
}

func main() {
	var meal_cost float64
	fmt.Scan(&meal_cost)

	var tip_percent, tax_percent int32
	fmt.Scan(&tip_percent, &tax_percent)

	solve(meal_cost, tip_percent, tax_percent)
}

func readLine(reader *bufio.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	return scanner.Text()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
