package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Printf("\nA^2 + B^2 = %.2f\n", math.Pow(a, 2)+math.Pow(b, 2))
	fmt.Printf("A^2 - B^2 = %.2f\n", math.Pow(a, 2)-math.Pow(b, 2))
	fmt.Printf("A^2 * B^2 = %.2f\n", math.Pow(a, 2)*math.Pow(b, 2))
	fmt.Printf("A^2 / B^2 = %.2f\n", math.Pow(a, 2)/math.Pow(b, 2))
}
