package main

import (
	"fmt"
	"math"
)

func main() {
	var a, h float64
	for i := 0; i < 3; i++ {
		fmt.Print("a = ")
		fmt.Scan(&a)
		fmt.Print("h = ")
		fmt.Scan(&h)
		fmt.Printf("\n%.2f\n", triangleP(a, h))
	}
}

func triangleP(a, h float64) float64 {
	b := math.Sqrt(math.Pow(a / 2, 2) + math.Pow(h, 2))
	return a + 2 * b
}