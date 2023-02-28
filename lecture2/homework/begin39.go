package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Print("A = ")
	fmt.Scan(&a)
	fmt.Print("B = ")
	fmt.Scan(&b)
	fmt.Print("C = ")
	fmt.Scan(&c)
	d := math.Pow(b, 2) - 4*a*c
	x1 := (-b - math.Sqrt(d)) / (2 * a)
	x2 := (-b + math.Sqrt(d)) / (2 * a)
	fmt.Printf("\nx1 = %.2f\n", x1)
	fmt.Printf("x2 = %.2f\n", x2)
}
